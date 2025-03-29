package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"pdftools/pdf" // 添加本地包导入
)

func main() {
	// 定义子命令
	mergeCmd := flag.NewFlagSet("merge", flag.ExitOnError)
	mergeDirPath := mergeCmd.String("dir", "", "要合并的PDF文件所在目录路径")

	extractCmd := flag.NewFlagSet("extract", flag.ExitOnError)
	extractInputDir := extractCmd.String("input", "", "输入目录路径")
	extractOutputDir := extractCmd.String("output", "", "输出目录路径")
	extractPageRange := extractCmd.String("pages", "", "页码范围，如：1,3-5")

	// 检查参数
	if len(os.Args) < 2 {
		fmt.Println("使用方法：")
		fmt.Println("  合并PDF：pdftools merge [-dir 目录路径]")
		fmt.Println("  提取页面：pdftools extract -input 输入目录 [-output 输出目录] -pages 页码范围")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "merge":
		mergeCmd.Parse(os.Args[2:])
		if err := pdf.MergePDFs(*mergeDirPath); err != nil { // 修改函数调用
			fmt.Fprintf(os.Stderr, "合并失败: %v\n", err)
			os.Exit(1)
		}

	case "extract":
		extractCmd.Parse(os.Args[2:])
		if *extractInputDir == "" || *extractPageRange == "" {
			extractCmd.PrintDefaults()
			os.Exit(1)
		}

		outputDir := *extractOutputDir
		if outputDir == "" {
			outputDir = filepath.Join(*extractInputDir, "output")
		}

		if err := pdf.ExtractPDFPages(*extractInputDir, outputDir, *extractPageRange); err != nil { // 修改函数调用
			fmt.Fprintf(os.Stderr, "提取失败: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Printf("未知的命令: %s\n", os.Args[1])
		os.Exit(1)
	}

	//fmt.Print("\n按回车键退出...")
	//fmt.Scanln()
}
