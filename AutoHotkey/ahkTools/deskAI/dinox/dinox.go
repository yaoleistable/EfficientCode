package dinox

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Dinox struct {
		Token string `json:"token"`
	} `json:"dinox"`
}

func getToken() string {
	// 获取可执行文件的路径
	execPath, err := os.Executable()
	if err != nil {
		fmt.Printf("无法获取可执行文件路径: %v\n", err)
		return ""
	}

	// 获取可执行文件所在目录
	execDir := filepath.Dir(execPath)
	// 构建配置文件路径（与可执行文件在同一目录）
	configPath := filepath.Join(execDir, "config.json")

	fmt.Printf("正在读取配置文件: %s\n", configPath)

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("无法读取配置文件: %v\n", err)
		return ""
	}

	// 解析 JSON
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Printf("解析配置文件失败: %v\n", err)
		return ""
	}

	return config.Dinox.Token
}

func DinoxPost(content string) error {
	url := "https://dinoai.chatgo.pro/openapi/text/input"
	method := "POST"
	token := getToken()
	if token == "" {
		return fmt.Errorf("token not found in config.json")
	}

	payload := strings.NewReader(`{
		"content": "` + content + `"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println("Response Code:", res.Status)
	fmt.Println("Response Body:", string(body))
	return nil
}
