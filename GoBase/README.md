# GoBase AI 文本处理工具

这是一个基于Go语言开发的AI文本处理工具，支持文本翻译、润色、总结和问答等功能。项目使用模块化设计，支持多种AI模型，可以灵活配置和使用。

## 功能特点

- 文本翻译：支持多语言之间的文本翻译
- 文本润色：改进文本的表达和语法
- 文本总结：生成文本的简洁摘要
- 智能问答：与AI助手进行对话交互
- 支持多种AI模型：可配置使用不同的AI服务提供商
- 命令行界面：简单直观的使用方式

## 环境要求

- Go 1.16 或更高版本
- 网络连接（用于访问AI服务）

## 安装步骤

1. 克隆项目代码：
   ```bash
   git clone https://github.com/yourusername/GoBase.git
   cd GoBase
   ```

2. 安装依赖：
   ```bash
   go mod download
   ```

3. 编译项目：
   ```bash
   go build
   ```

## 配置说明

在使用之前，需要正确配置`config.json`文件，设置AI模型的相关参数：

```json
{
  "models": {
    "模型名称": {
      "url": "API端点URL",
      "token": "你的API密钥",
      "model": "具体的模型名称",
      "temperature": 0.7
    }
  }
}
```

配置参数说明：
- `url`：AI服务的API端点
- `token`：访问API所需的密钥
- `model`：具体使用的模型名称
- `temperature`：生成文本的随机性参数（0.0-1.0）

## 使用方法

### 基本命令格式

```bash
go run main.go [函数名] [文本]
```

可选参数：
- `-model`：指定要使用的AI模型名称（默认：qwen-plus）

### 支持的功能

1. 文本翻译
   ```bash
   go run main.go translate "Hello, world!"
   ```

2. 文本润色
   ```bash
   go run main.go polish "This is a draft text."
   ```

3. 文本总结
   ```bash
   go run main.go summarize "这是一段需要总结的长文本..."
   ```

4. 智能问答
   ```bash
   go run main.go ask "什么是人工智能？"
   ```

### 指定模型

```bash
go run main.go -model deepseek translate "Hello, world!"
```

## 错误处理

- 如果命令格式错误，程序会显示正确的使用方法
- API调用失败时会返回详细的错误信息
- 可以通过错误信息快速定位问题

## 开发说明

项目使用模块化设计，主要包含以下部分：

- `main.go`：主程序入口，处理命令行参数
- `ai/`：AI功能模块
  - `ai.go`：核心AI调用功能
  - `config.go`：配置管理
  - `text_utils.go`：文本处理工具

## 贡献指南

欢迎提交Issue和Pull Request来改进项目。在提交代码前，请确保：

1. 代码符合Go的编码规范
2. 添加了必要的测试用例
3. 更新了相关文档

## 许可证

[MIT License](LICENSE)