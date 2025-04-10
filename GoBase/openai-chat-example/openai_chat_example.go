// 主程序包
package main

import (
	"context"
	"fmt"
	"log"

	// 导入 langchaingo 的 LLM 接口包
	"github.com/tmc/langchaingo/llms"
	// 导入 OpenAI 实现包
	"github.com/tmc/langchaingo/llms/openai"
)

// 通过修改，实现open ai 兼容的llm，方便使用这个来调用各种openai兼容模型
// 下一步进一步修改，改造为函数式调用，方便使用
func main() {
	// 创建一个新的 OpenAI LLM 实例
	llm, err := openai.New(
		// 直接设置 API 密钥
		openai.WithToken("your-api-key"),
		// 直接设置自定义 API 地址
		openai.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1"),
		// 设置模型名称为阿里云支持的模型
		openai.WithModel("qwen-plus"),
	)
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个背景上下文
	ctx := context.Background()

	// 准备对话内容
	content := []llms.MessageContent{
		// 设置系统角色消息，定义 AI 的行为
		llms.TextParts(llms.ChatMessageTypeSystem, "你是一个综合智能助手"),
		// 设置用户问题
		llms.TextParts(llms.ChatMessageTypeHuman, "如何提升办公office使用技巧?"),
	}

	// 生成 AI 响应内容
	if _, err := llm.GenerateContent(ctx, content,
		// 设置最大令牌数为 1024
		llms.WithMaxTokens(1024),
		// 设置流式处理函数，实时打印 AI 的响应
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		})); err != nil {
		log.Fatal(err)
	}
}
