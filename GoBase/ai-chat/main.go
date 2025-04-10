package main

import (
	"ai/ai"
	"flag"
	"fmt"
	"os"
)

func main() {
	// 定义命令行参数
	modelName := flag.String("model", "qwen-plus", "要使用的AI模型名称")

	// 解析命令行参数
	flag.Parse()

	// 获取非flag参数
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("使用方法: go run main.go [函数名] [文本]")
		fmt.Println("可用函数: translate, polish, summarize, ask")
		os.Exit(1)
	}

	// 获取函数名和输入文本
	funcName := args[0]
	inputText := args[1]

	// 根据函数名调用对应的函数
	var result string
	var err error

	switch funcName {
	case "translate":
		result, err = ai.TranslateText(*modelName, inputText)
	case "polish":
		result, err = ai.PolishText(*modelName, inputText)
	case "summarize":
		result, err = ai.SummarizeText(*modelName, inputText)
	case "ask":
		result, err = ai.AskAssistant(*modelName, inputText)
	default:
		fmt.Printf("未知的函数名: %s\n", funcName)
		fmt.Println("可用函数: translate, polish, summarize, ask")
		os.Exit(1)
	}

	// 处理结果
	if err != nil {
		fmt.Printf("执行失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("输入文本: %s\n", inputText)
	fmt.Printf("处理结果: %s\n", result)
}
