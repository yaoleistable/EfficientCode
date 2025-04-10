package ai

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/openai"
)

// AskAI 统一的AI调用函数
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

	// 创建上下文
	ctx := context.Background()

	// 创建 LLM 实例
	var llm llms.Model
	if strings.HasPrefix(modelName, "gemini-") {
		// 使用 Google AI
		customClient := &http.Client{
			Transport: &customTransport{
				base:    http.DefaultTransport,
				baseURL: "generativelanguage.googleapis.com:443",
			},
		}
		llm, err = createGoogleLLM(ctx, modelConfig, customClient)
	} else {
		// 使用 OpenAI 或兼容接口
		llm, err = createOpenAILLM(modelConfig)
	}

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
	_, err = llm.GenerateContent(ctx, content,
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

// createOpenAILLM 创建 OpenAI 及其兼容的 LLM 实例
func createOpenAILLM(config *AIConfig) (llms.Model, error) {
	return openai.New(
		openai.WithToken(config.Token),
		openai.WithModel(config.Model),
		openai.WithBaseURL(config.URL),
	)
}

// createGoogleLLM 创建 Google AI LLM 实例
func createGoogleLLM(ctx context.Context, config *AIConfig, client *http.Client) (llms.LLM, error) {
	options := []googleai.Option{
		googleai.WithAPIKey(config.Token),
		googleai.WithDefaultModel(config.Model),
		googleai.WithHTTPClient(client),
	}
	return googleai.New(ctx, options...)
}

// customTransport 自定义传输层结构体
type customTransport struct {
	base    http.RoundTripper
	baseURL string
}

// RoundTrip 实现 http.RoundTripper 接口
func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Host = t.baseURL
	return t.base.RoundTrip(req)
}
