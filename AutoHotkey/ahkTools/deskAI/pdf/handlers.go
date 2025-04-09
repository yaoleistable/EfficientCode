package pdf

import (
	"flag"
	"fmt"
	"path/filepath"
)

// HandleMerge 处理 PDF 合并命令
func HandleMerge(args []string) (string, error) {
	pdfMergeCmd := flag.NewFlagSet("pdfMerge", flag.ExitOnError)
	mergeDirPath := pdfMergeCmd.String("dir", "", "要合并的PDF文件所在目录路径")
	pdfMergeCmd.Parse(args)

	if *mergeDirPath == "" {
		return "", fmt.Errorf("使用方法: deskAI.exe pdfMerge -dir 目录路径")
	}

	if err := MergePDFs(*mergeDirPath); err != nil {
		return "", err
	}
	return "PDF合并完成", nil
}

// HandleExtract 处理 PDF 提取命令
func HandleExtract(args []string) (string, error) {
	pdfExtractCmd := flag.NewFlagSet("pdfExtract", flag.ExitOnError)
	extractInputDir := pdfExtractCmd.String("input", "", "输入目录路径")
	extractOutputDir := pdfExtractCmd.String("output", "", "输出目录路径")
	extractPageRange := pdfExtractCmd.String("pages", "", "页码范围，如：1,3-5")
	pdfExtractCmd.Parse(args)

	if *extractInputDir == "" || *extractPageRange == "" {
		return "", fmt.Errorf("使用方法: deskAI.exe pdfExtract -input 输入目录 [-output 输出目录] -pages 页码范围")
	}

	outputDir := *extractOutputDir
	if outputDir == "" {
		outputDir = filepath.Join(*extractInputDir, "output")
	}

	if err := ExtractPDFPages(*extractInputDir, outputDir, *extractPageRange); err != nil {
		return "", err
	}
	return "PDF页面提取完成", nil
}