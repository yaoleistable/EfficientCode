package main

import (
	"deskAI/ai"
	"deskAI/dinox"
	"deskAI/pdf"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func showUsage() {
	fmt.Println("使用方法:")
	fmt.Println("  deskAI.exe [命令] [参数]")
	fmt.Println("命令列表:")
	fmt.Println("  dinoxPost \"内容\"       - 发送内容到 DinoAI")
	fmt.Println("  translate [-model 模型名] \"文本\"  - 执行翻译（中英互译）")
	fmt.Println("  polish [-model 模型名] \"文本\"     - 执行文本润色")
	fmt.Println("  summarize [-model 模型名] \"文本\"  - 执行文本总结")
	fmt.Println("  ask [-model 模型名] \"问题\"        - 向AI助手提问")
	fmt.Println("  pdfMerge -dir 目录路径           - 合并指定目录下的PDF文件")
	fmt.Println("  pdfExtract -input 输入目录 [-output 输出目录] -pages 页码范围  - 提取PDF页面")
	fmt.Println("  help                   - 显示帮助信息")
}

// 写入结果到文件
func writeResult(result string) error {
	// 获取当前执行目录
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("获取当前目录失败: %v", err)
	}

	// 构建结果文件的完整路径
	resultPath := filepath.Join(currentDir, "result.txt")

	// 写入结果
	err = os.WriteFile(resultPath, []byte(result), 0644)
	if err != nil {
		return fmt.Errorf("写入结果文件失败: %v", err)
	}

	fmt.Printf("结果已写入: %s\n", resultPath)
	return nil
}

// 处理命令并返回结果
func handleCommand(command string, args []string, modelName *string) (string, error) {
	switch command {
	case "dinoxPost":
		if len(args) != 1 {
			return "", fmt.Errorf("使用方法: deskAI.exe dinoxPost \"你要发送的内容\"")
		}
		if err := dinox.DinoxPost(args[0]); err != nil {
			return "", err
		}
		return "", nil

	case "translate", "polish", "summarize", "ask":
		if len(args) < 1 {
			return "", fmt.Errorf("使用方法: deskAI.exe %s [-model 模型名] \"文本\"", command)
		}
		return ai.HandleCommand(command, *modelName, args[len(args)-1])

	case "pdfMerge":
		return pdf.HandleMerge(args)

	case "pdfExtract":
		return pdf.HandleExtract(args)

	case "help":
		showUsage()
		return "", nil

	default:
		return "", fmt.Errorf("未知命令: %s", command)
	}
}

func main() {
	modelName := flag.String("model", "qwen-plus", "AI模型名称")
	flag.Parse()

	if len(flag.Args()) < 1 {
		showUsage()
		return
	}

	command := flag.Arg(0)
	args := flag.Args()[1:]

	result, err := handleCommand(command, args, modelName)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		if result != "" {
			if writeErr := writeResult(fmt.Sprintf("错误: %v", err)); writeErr != nil {
				fmt.Printf("写入错误信息失败: %v\n", writeErr)
			}
		}
		os.Exit(1)
	}

	// 写入结果
	if result != "" && command != "dinoxPost" && command != "help" {
		if err := writeResult(result); err != nil {
			fmt.Printf("写入结果失败: %v\n", err)
			os.Exit(1)
		}
	}
}
