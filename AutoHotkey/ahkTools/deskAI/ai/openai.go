package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

// AskAI 调用AI模型并返回响应
func AskAI(modelName, prompt, question string) (string, error) {
	// 加载配置
	config, err := LoadConfig()
	if err != nil {
		return "", fmt.Errorf("加载配置失败: %v", err)
	}

	// 获取指定模型的配置
	modelConfig, err := config.GetModelConfig(modelName)
	if err != nil {
		return "", err
	}

	// 创建 LLM 实例
	llm, err := createLLM(modelConfig)
	if err != nil {
		return "", fmt.Errorf("创建LLM实例失败: %v", err)
	}

	// 准备对话内容
	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, prompt),
		llms.TextParts(llms.ChatMessageTypeHuman, question),
	}

	// 创建一个用于存储响应的 builder
	var responseBuilder strings.Builder

	// 生成 AI 响应内容
	_, err = llm.GenerateContent(context.Background(), content,
		llms.WithMaxTokens(2048),
		llms.WithTemperature(modelConfig.Temperature),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			// 实时打印响应
			fmt.Print(string(chunk))
			// 同时保存到 builder 中
			responseBuilder.Write(chunk)
			return nil
		}))

	if err != nil {
		return "", fmt.Errorf("生成内容失败: %v", err)
	}

	return responseBuilder.String(), nil
}

// createLLM 根据配置创建 LLM 实例
func createLLM(config *AIConfig) (llms.Model, error) {
	return openai.New(
		openai.WithToken(config.Token),
		openai.WithModel(config.Model),
		openai.WithBaseURL(config.URL),
	)
}
