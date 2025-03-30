package main

import (
	"deskAI/dinox"
	"fmt"
	"os"
)

func showUsage() {
	fmt.Println("使用方法:")
	fmt.Println("  deskAI.exe dinoxPost \"内容\"  - 发送内容到 DinoAI")
	fmt.Println("  deskAI.exe pdf ...           - PDF 相关功能（待实现）")
	fmt.Println("  deskAI.exe help              - 显示帮助信息")
}

func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "dinoxPost":
		if len(os.Args) != 3 {
			fmt.Println("使用方法: deskAI.exe dinoxPost \"你要发送的内容\"")
			return
		}
		content := os.Args[2]
		if err := dinox.DinoxPost(content); err != nil {
			fmt.Printf("Error: %v\n", err)
		}

	case "pdf":
		fmt.Println("PDF 功能待实现")
		// TODO: 实现 PDF 相关功能

	case "help":
		showUsage()

	default:
		fmt.Printf("未知命令: %s\n", command)
		showUsage()
	}
}
