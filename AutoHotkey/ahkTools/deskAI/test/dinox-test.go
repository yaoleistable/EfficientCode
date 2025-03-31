package main

import (
	"deskAI/dinox"
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("无法获取当前文件路径\n")
		return
	}
	// 获取 dinox 目录的路径
	dinoxDir := filepath.Dir(filename)
	fmt.Printf("dinoxDir: %s\n", dinoxDir)
	// 获取 DeskAI 目录的路径
	deskAIDir := filepath.Dir(dinoxDir)
	fmt.Printf("deskAIDir: %s\n", deskAIDir)
	content := " DinoAI! 发送内容测试！"
	if err := dinox.DinoxPost(content); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
