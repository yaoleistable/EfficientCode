package ai

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// AIConfig 存储单个AI模型的配置信息
type AIConfig struct {
	URL         string  `json:"url"`
	Token       string  `json:"token"`
	Model       string  `json:"model"`
	Temperature float64 `json:"temperature"`
}

// Config 存储所有AI模型的配置信息
type Config struct {
	Models map[string]AIConfig `json:"models"`
}

// LoadConfig 从config.json文件中读取AI配置
func LoadConfig() (*Config, error) {
	// 获取可执行文件的路径
	execPath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("获取可执行文件路径失败: %v", err)
	}

	// 获取可执行文件所在目录
	execDir := filepath.Dir(execPath)
	// 构建配置文件路径（与可执行文件在同一目录）
	configPath := filepath.Join(execDir, "config.json")

	fmt.Printf("正在读取配置文件: %s\n", configPath)

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析JSON数据
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	// 验证配置
	if len(config.Models) == 0 {
		return nil, fmt.Errorf("配置文件中没有找到任何AI模型配置")
	}

	return &config, nil
}

// GetModelConfig 根据模型名称获取对应的配置
func (c *Config) GetModelConfig(modelName string) (*AIConfig, error) {
	config, exists := c.Models[modelName]
	if !exists {
		return nil, fmt.Errorf("未找到模型 '%s' 的配置信息", modelName)
	}
	return &config, nil
}
