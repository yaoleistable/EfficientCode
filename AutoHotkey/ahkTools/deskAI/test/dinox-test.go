package main

import (
	"deskAI/dinox"
	"fmt"
)

func main() {
	content := " DinoAI! 发送内容测试！"
	if err := dinox.DinoxPost(content); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
