package ai

import (
	"fmt"
)

// HandleCommand 处理 AI 相关命令
func HandleCommand(command string, modelName string, text string) (string, error) {
	switch command {
	case "translate":
		return TranslateText(modelName, text)
	case "polish":
		return PolishText(modelName, text)
	case "summarize":
		return SummarizeText(modelName, text)
	case "ask":
		return AskAssistant(modelName, text)
	}
	return "", fmt.Errorf("未知的 AI 命令: %s", command)
}
