package pdf

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// 解析页码范围字符串（如"1,3-5"）
func parsePageRangeStr(pageStr string, totalPages int) ([]int, error) {
	seen := make(map[int]struct{})
	parts := regexp.MustCompile(`\s*,\s*`).Split(pageStr, -1)
	validRange := regexp.MustCompile(`^(\d+)(?:-(\d+))?$`)

	for _, part := range parts {
		match := validRange.FindStringSubmatch(part)
		if match == nil {
			return nil, fmt.Errorf("invalid page format: %s", part)
		}

		start, _ := strconv.Atoi(match[1])
		end := start
		if len(match) > 2 && match[2] != "" {
			end, _ = strconv.Atoi(match[2])
		}

		if start < 1 || end > totalPages || start > end {
			return nil, fmt.Errorf("page out of range (1-%d): %s", totalPages, part)
		}

		for i := start; i <= end; i++ {
			seen[i] = struct{}{}
		}
	}

	pages := make([]int, 0, len(seen))
	for k := range seen {
		pages = append(pages, k)
	}
	sort.Ints(pages)
	return pages, nil
}

// 生成输出文件名（如input_p1-3.pdf）
func generateOutputFilePath(inputPath, outputDir, pageRange string) string {
	base := filepath.Base(inputPath)
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)

	// 处理中文文件名，替换特殊字符
	sanitizedName := regexp.MustCompile(`[\\/:*?"<>|]`).ReplaceAllString(name, "_")

	return filepath.Join(outputDir, fmt.Sprintf("%s_p%s%s", sanitizedName, pageRange, ext))
}

// 核心提取函数
func extractPagesFromPDF(inputPath, outputDir, pageRange string) error {
	// 生成输出路径
	outputPath := generateOutputFilePath(inputPath, outputDir, pageRange)

	// 将页码范围转换为 pdfcpu 支持的格式
	// 例如: "1,4-6" -> ["1", "4-6"]
	pages := strings.Split(pageRange, ",")

	// 使用 pdfcpu 的 TrimFile 功能来提取页面
	err := api.TrimFile(inputPath, outputPath, pages, nil)
	if err != nil {
		return fmt.Errorf("failed to extract pages: %w", err)
	}

	log.Printf("Successfully created: %s", outputPath)
	return nil
}

// 批量处理（支持并发）
func processPDFBatch(inputDir, outputDir, pageRange string, maxWorkers int) error {
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxWorkers)
	var processError error
	var errorMutex sync.Mutex

	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.EqualFold(filepath.Ext(path), ".pdf") {
			wg.Add(1)
			sem <- struct{}{}

			go func(p string) {
				defer func() {
					<-sem
					wg.Done()
				}()

				if err := extractPagesFromPDF(p, outputDir, pageRange); err != nil {
					errorMutex.Lock()
					processError = err
					errorMutex.Unlock()
					log.Printf("Error processing %s: %v", p, err)
				}
			}(path)
		}
		return nil
	})

	wg.Wait()

	if err != nil {
		return fmt.Errorf("error walking directory: %w", err)
	}
	if processError != nil {
		return fmt.Errorf("error during processing: %w", processError)
	}
	return nil
}

// ExtractPDFPages 提取PDF页面
func ExtractPDFPages(inputDir string, outputDir string, pageRange string) error {
	maxWorkers := 5

	// 检查输入目录是否存在
	inputInfo, err := os.Stat(inputDir)
	if err != nil {
		return fmt.Errorf("error accessing input directory: %v", err)
	}
	if !inputInfo.IsDir() {
		return fmt.Errorf("input path is not a directory: %s", inputDir)
	}

	// 创建输出目录
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// 验证页码格式
	if matched, _ := regexp.MatchString(`^[\d,\s-]+$`, pageRange); !matched {
		return fmt.Errorf("invalid page range format. Example: \"1,3-5\"")
	}

	if err := processPDFBatch(inputDir, outputDir, pageRange, maxWorkers); err != nil {
		return fmt.Errorf("batch processing failed: %v", err)
	}

	log.Printf("Processing completed successfully")
	return nil
}
