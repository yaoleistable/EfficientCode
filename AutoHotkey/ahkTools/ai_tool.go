package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"gopkg.in/ini.v1"
)

// 配置结构
type Config struct {
	API       APIConfig                 `json:"api"`
	Functions map[string]FunctionConfig `json:"functions"`
}

// API配置
type APIConfig struct {
	BaseURL string `json:"base_url"`
	APIKey  string `json:"api_key"`
	Model   string `json:"model"`
	Timeout string `json:"timeout"`
}

// 功能配置
type FunctionConfig struct {
	SystemPrompt string  `json:"system_prompt"`
	UserPrompt   string  `json:"user_prompt"`
	Temperature  float64 `json:"temperature"`
	Stream       bool    `json:"stream"`
}

// API请求结构
type APIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
	Stream      bool      `json:"stream"`
}

// 消息结构
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// API响应结构
type APIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// AITool 结构体
type AITool struct {
	Config Config
	Logger *log.Logger
}

// 初始化AI工具
func NewAITool(configPath string) (*AITool, error) {
	// 设置日志
	logFile, err := os.OpenFile("ai_tool.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("无法创建日志文件: %v", err)
	}

	logger := log.New(logFile, "AITool: ", log.LstdFlags)

	// 加载配置
	config, err := loadConfig(configPath)
	if err != nil {
		return nil, err
	}

	return &AITool{
		Config: config,
		Logger: logger,
	}, nil
}

// 加载配置文件
func loadConfig(configPath string) (Config, error) {
	var config Config

	// 检查文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return config, fmt.Errorf("配置文件不存在: %s", configPath)
	}

	// 解析INI文件
	cfg, err := ini.Load(configPath)
	if err != nil {
		return config, fmt.Errorf("无法加载配置文件: %v", err)
	}

	// 初始化结构
	config.API = APIConfig{}
	config.Functions = make(map[string]FunctionConfig)

	// 解析API部分
	apiSection := cfg.Section("API")
	config.API.BaseURL = apiSection.Key("base_url").String()
	config.API.APIKey = apiSection.Key("api_key").String()
	config.API.Model = apiSection.Key("model").String()
	config.API.Timeout = apiSection.Key("timeout").String()

	// 验证热键格式
	hotkeySection := cfg.Section("Hotkey")
	for _, key := range hotkeySection.Keys() {
		hotkey := strings.Split(key.String(), ";")[0]
		hotkey = strings.TrimSpace(hotkey)
		if !validateHotkey(hotkey) {
			return config, fmt.Errorf("无效的热键格式 '%s': %s", key.Name(), hotkey)
		}
	}

	// 解析Functions部分
	functionsSection := cfg.Section("Functions")
	for _, key := range functionsSection.Keys() {
		var functionConfig FunctionConfig
		jsonStr := key.String()
		err := json.Unmarshal([]byte(jsonStr), &functionConfig)
		if err != nil {
			return config, fmt.Errorf("Functions配置解析错误 '%s': %v", key.Name(), err)
		}
		config.Functions[key.Name()] = functionConfig
	}

	return config, nil
}

// 验证热键格式
func validateHotkey(hotkey string) bool {
	validModifiers := []string{"!", "^", "+", "#"}

	if len(hotkey) < 2 {
		return false
	}

	modifier := hotkey[0:1]
	key := hotkey[1:]

	for _, validMod := range validModifiers {
		if modifier == validMod && len(key) == 1 {
			return true
		}
	}

	return false
}

// 发送API请求
func (a *AITool) makeAPIRequest(messages []Message, functionConfig FunctionConfig) (APIResponse, error) {
	var response APIResponse

	// 准备请求数据
	data := APIRequest{
		Model:       a.Config.API.Model,
		Messages:    messages,
		Temperature: functionConfig.Temperature,
		Stream:      functionConfig.Stream,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return response, fmt.Errorf("JSON编码错误: %v", err)
	}

	// 设置超时
	timeout, err := strconv.Atoi(a.Config.API.Timeout)
	if err != nil {
		timeout = 30 // 默认30秒
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	// 创建请求
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/chat/completions", a.Config.API.BaseURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return response, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.Config.API.APIKey))

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return response, fmt.Errorf("API请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return response, fmt.Errorf("API返回错误状态码: %d, 响应: %s", resp.StatusCode, string(bodyBytes))
	}

	// 解析响应
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response, fmt.Errorf("解析响应失败: %v", err)
	}

	return response, nil
}

// 处理文本
func (a *AITool) ProcessText(functionName string, text string, kwargs map[string]string) (string, error) {
	functionConfig, ok := a.Config.Functions[functionName]
	if !ok {
		return "", fmt.Errorf("未定义的功能: %s", functionName)
	}

	// 准备消息
	messages := []Message{
		{
			Role:    "system",
			Content: functionConfig.SystemPrompt,
		},
		{
			Role:    "user",
			Content: formatPrompt(functionConfig.UserPrompt, text, kwargs),
		},
	}

	// 发送请求
	result, err := a.makeAPIRequest(messages, functionConfig)
	if err != nil {
		a.Logger.Printf("处理失败: %v", err)
		return "", err
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("API返回空结果")
	}

	return strings.TrimSpace(result.Choices[0].Message.Content), nil
}

// 格式化提示词
func formatPrompt(prompt string, text string, kwargs map[string]string) string {
	result := strings.Replace(prompt, "{text}", text, -1)
	for key, value := range kwargs {
		result = strings.Replace(result, fmt.Sprintf("{%s}", key), value, -1)
	}
	return result
}

// 检测语言
func (a *AITool) DetectLanguage(text string) string {
	for _, r := range text {
		if unicode.Is(unicode.Han, r) {
			return "zh"
		}
	}
	return "en"
}

// 翻译文本
func (a *AITool) Translate(text string) (string, error) {
	sourceLang := a.DetectLanguage(text)
	targetLang := "English"
	if sourceLang == "en" {
		targetLang = "Chinese"
	}

	kwargs := map[string]string{
		"source_lang": sourceLang,
		"target_lang": targetLang,
	}

	return a.ProcessText("translate", text, kwargs)
}

// 润色文本
func (a *AITool) Polish(text string) (string, error) {
	return a.ProcessText("polish", text, nil)
}

// 总结文本
func (a *AITool) Summarize(text string) (string, error) {
	return a.ProcessText("summarize", text, nil)
}

func main() {
	// 检查参数
	if len(os.Args) < 3 {
		fmt.Println("用法: ai_tool <function> <text>")
		os.Exit(1)
	}

	functionName := os.Args[1]
	text := os.Args[2]

	// 初始化AI工具
	aiTool, err := NewAITool("config.ini")
	if err != nil {
		writeResult(fmt.Sprintf("错误: %v", err))
		os.Exit(1)
	}

	var result string

	// 根据功能名称调用相应的方法
	switch functionName {
	case "translate":
		result, err = aiTool.Translate(text)
	case "polish":
		result, err = aiTool.Polish(text)
	case "summarize":
		result, err = aiTool.Summarize(text)
	default:
		result, err = aiTool.ProcessText(functionName, text, nil)
	}

	if err != nil {
		writeResult(fmt.Sprintf("错误: %v", err))
		os.Exit(1)
	}

	// 写入结果
	writeResult(result)
}

// 写入结果到文件
func writeResult(result string) {
	err := os.WriteFile("result.txt", []byte(result), 0644)
	if err != nil {
		log.Printf("写入结果失败: %v", err)
	}
}
