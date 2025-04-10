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
	execPath, err := os.Executable()
	if err != nil {
		LogError("获取可执行文件路径失败: %v", err)
		return nil, fmt.Errorf("获取可执行文件路径失败: %v", err)
	}

	execDir := filepath.Dir(execPath)
	configPath := filepath.Join(execDir, "config.json")

	// 将输出改为日志记录
	LogInfo("正在读取配置文件: %s", configPath)

	data, err := os.ReadFile(configPath)
	if err != nil {
		LogError("读取配置文件失败: %v", err)
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		LogError("解析配置文件失败: %v", err)
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	if len(config.Models) == 0 {
		LogError("配置文件中没有找到任何AI模型配置")
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
