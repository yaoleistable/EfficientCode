package pdf

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

// MergePDFs 合并PDF文件
func MergePDFs(dirPath string) error {
	// 如果没有提供路径，使用当前目录
	if dirPath == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("获取当前目录失败: %v", err)
		}
		dirPath = currentDir
	}

	// 获取目录中所有PDF文件
	var pdfFiles []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.ToLower(filepath.Ext(path)) == ".pdf" {
			pdfFiles = append(pdfFiles, path)
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("遍历目录失败: %v", err)
	}

	if len(pdfFiles) == 0 {
		return fmt.Errorf("目录中没有找到PDF文件")
	}

	// 对文件名进行排序
	sort.Strings(pdfFiles)

	// 合并后的文件名
	outputPath := filepath.Join(dirPath, "merged.pdf")

	// 创建默认配置
	conf := model.NewDefaultConfiguration()

	// 使用更新后的 API 调用方式
	err = api.MergeCreateFile(pdfFiles, outputPath, false, conf)
	if err != nil {
		return fmt.Errorf("合并PDF文件失败: %v", err)
	}

	fmt.Printf("成功合并 %d 个PDF文件到: %s\n", len(pdfFiles), outputPath)
	return nil
}
