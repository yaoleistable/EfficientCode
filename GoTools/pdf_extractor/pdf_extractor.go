package main

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

				if err := extractPagesFromPDF(p, outputDir, pageRange); err != nil { // 改用新的函数名
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

func main() {
	// 使用命令行参数
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <input_dir> [output_dir] <page_range>\n", filepath.Base(os.Args[0]))
		fmt.Println("Example: .\\pdf_extractor.exe .\\input \"1,3-5\"")
		fmt.Println("   或者: .\\pdf_extractor.exe .\\input .\\output \"1,3-5\"")
		os.Exit(1)
	}

	inputDir := os.Args[1]
	var outputDir string
	var pageRange string

	// 根据参数数量判断输出目录
	if len(os.Args) == 3 {
		// 如果只有两个参数，使用默认输出目录
		outputDir = filepath.Join(inputDir, "output")
		pageRange = os.Args[2]
	} else {
		// 如果有三个参数，使用指定的输出目录
		outputDir = os.Args[2]
		pageRange = os.Args[3]
	}

	maxWorkers := 5

	// 检查输入目录是否存在
	inputInfo, err := os.Stat(inputDir)
	if err != nil {
		log.Fatalf("Error accessing input directory: %v", err)
	}
	if !inputInfo.IsDir() {
		log.Fatalf("Input path is not a directory: %s", inputDir)
	}

	// 创建输出目录
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// 验证页码格式
	if matched, _ := regexp.MatchString(`^[\d,\s-]+$`, pageRange); !matched {
		log.Fatalf("Invalid page range format. Example: \"1,3-5\"")
	}

	if err := processPDFBatch(inputDir, outputDir, pageRange, maxWorkers); err != nil {
		log.Fatalf("Batch processing failed: %v", err)
	}

	log.Printf("Processing completed successfully")
	fmt.Print("\n按回车键退出...")
	fmt.Scanln()
}
