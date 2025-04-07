package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// OpenAIRequest OpenAI API请求结构
type OpenAIRequest struct {
	Model       string      `json:"model"`
	Messages    []OpenAIMsg `json:"messages"`
	Temperature float64     `json:"temperature"`
}

type OpenAIMsg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIResponse OpenAI API响应结构
type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// AnthropicRequest Anthropic API请求结构
type AnthropicRequest struct {
	Model       string         `json:"model"`
	Messages    []AnthropicMsg `json:"messages"`
	Temperature float64        `json:"temperature"`
}

type AnthropicMsg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// AnthropicResponse Anthropic API响应结构
type AnthropicResponse struct {
	Content string `json:"content"`
	Error   *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// QwenRequest 阿里云Qwen API请求结构
type QwenRequest struct {
	Model      string         `json:"model"`
	Input      QwenInput      `json:"input"`
	Parameters QwenParameters `json:"parameters"`
}

type QwenInput struct {
	Prompt string `json:"prompt"`
}

type QwenParameters struct {
	Temperature float64 `json:"temperature"`
}

// QwenResponse 阿里云Qwen API响应结构
type QwenResponse struct {
	Output struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Text string `json:"text,omitempty"`
	} `json:"output"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

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

	// 根据URL判断使用哪个API
	var response string
	if strings.Contains(modelConfig.URL, "openrouter.ai") {
		response, err = callOpenRouter(modelConfig, prompt, question)
	} else if strings.Contains(modelConfig.URL, "openai.com") {
		response, err = callOpenAI(modelConfig, prompt, question)
	} else if strings.Contains(modelConfig.URL, "anthropic.com") {
		response, err = callAnthropic(modelConfig, prompt, question)
	} else if strings.Contains(modelConfig.URL, "dashscope.aliyuncs.com") {
		response, err = callQwen(modelConfig, prompt, question)
	} else {
		return "", fmt.Errorf("不支持的API URL: %s", modelConfig.URL)
	}

	if err != nil {
		return "", fmt.Errorf("调用AI API失败: %v", err)
	}

	return response, nil
}

// callOpenAI 调用OpenAI API
func callOpenAI(config *AIConfig, prompt, question string) (string, error) {
	// 构建请求体
	request := OpenAIRequest{
		Model:       config.Model,
		Temperature: config.Temperature,
		Messages: []OpenAIMsg{
			{Role: "system", Content: prompt},
			{Role: "user", Content: question},
		},
	}

	// 发送请求
	response, err := sendRequest(config, request)
	if err != nil {
		return "", err
	}

	// 解析响应
	var openAIResp OpenAIResponse
	if err := json.Unmarshal(response, &openAIResp); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}

	// 检查错误
	if openAIResp.Error != nil {
		return "", fmt.Errorf("API错误: %s", openAIResp.Error.Message)
	}

	// 返回结果
	if len(openAIResp.Choices) > 0 {
		return openAIResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("API没有返回任何结果")
}

// callAnthropic 调用Anthropic API
func callAnthropic(config *AIConfig, prompt, question string) (string, error) {
	// 构建请求体
	request := AnthropicRequest{
		Model:       config.Model,
		Temperature: config.Temperature,
		Messages: []AnthropicMsg{
			{Role: "user", Content: prompt + "\n" + question},
		},
	}

	// 发送请求
	response, err := sendRequest(config, request)
	if err != nil {
		return "", err
	}

	// 解析响应
	var anthropicResp AnthropicResponse
	if err := json.Unmarshal(response, &anthropicResp); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}

	// 检查错误
	if anthropicResp.Error != nil {
		return "", fmt.Errorf("API错误: %s", anthropicResp.Error.Message)
	}

	return anthropicResp.Content, nil
}

// callOpenRouter 调用OpenRouter API
func callOpenRouter(config *AIConfig, prompt, question string) (string, error) {
	// 构建请求体
	request := OpenAIRequest{
		Model:       config.Model,
		Temperature: config.Temperature,
		Messages: []OpenAIMsg{
			{Role: "system", Content: prompt},
			{Role: "user", Content: question},
		},
	}

	// 发送请求
	response, err := sendRequest(config, request)
	if err != nil {
		return "", err
	}

	// 解析响应
	var openAIResp OpenAIResponse
	if err := json.Unmarshal(response, &openAIResp); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}

	// 检查错误
	if openAIResp.Error != nil {
		return "", fmt.Errorf("API错误: %s", openAIResp.Error.Message)
	}

	// 返回结果
	if len(openAIResp.Choices) > 0 {
		return openAIResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("API没有返回任何结果")
}

// callQwen 调用阿里云Qwen API
func callQwen(config *AIConfig, prompt, question string) (string, error) {
	// 构建请求体
	// 使用更规范的方式组织prompt和question
	formattedPrompt := fmt.Sprintf("系统指令：\n%s\n\n用户输入：\n%s", prompt, question)
	request := QwenRequest{
		Model: config.Model,
		Input: QwenInput{
			Prompt: formattedPrompt,
		},
		Parameters: QwenParameters{
			Temperature: config.Temperature,
		},
	}

	// 发送请求
	response, err := sendRequest(config, request)
	if err != nil {
		return "", err
	}

	// 解析响应
	var qwenResp QwenResponse
	if err := json.Unmarshal(response, &qwenResp); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}

	// 检查错误
	if qwenResp.Code != "" {
		return "", fmt.Errorf("API错误: [%s] %s", qwenResp.Code, qwenResp.Message)
	}

	// 优先使用choices中的内容，如果没有则使用text字段
	if len(qwenResp.Output.Choices) > 0 {
		return qwenResp.Output.Choices[0].Message.Content, nil
	}
	if qwenResp.Output.Text != "" {
		return qwenResp.Output.Text, nil
	}

	return "", fmt.Errorf("API没有返回任何结果")
}

// sendRequest 发送HTTP请求
func sendRequest(config *AIConfig, request interface{}) ([]byte, error) {
	// 序列化请求体
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("序列化请求体失败: %v", err)
	}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", config.URL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("创建HTTP请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Token)

	// 为OpenRouter API设置特殊请求头
	if strings.Contains(config.URL, "openrouter.ai") {
		req.Header.Set("HTTP-Referer", "https://github.com/trae-cn/gobase")
		req.Header.Set("X-Title", "GoBase")
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送HTTP请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %v", err)
	}

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP请求失败，状态码: %d，响应: %s", resp.StatusCode, string(body))
	}

	return body, nil
}
