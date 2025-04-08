# deskAI 使用说明

deskAI 是一个多功能的 AI 辅助工具，集成了文本处理、翻译、内容发送等功能。

## 1. 功能特点

### 1.1 DinoAI 笔记发送
- 支持将文本内容快速发送到 DinoAI 平台
- 自动处理认证和格式化

### 1.2 文本处理功能
- **翻译功能**：支持中英文互译
- **文本润色**：优化文本表达，使其更专业流畅
- **文本总结**：将长文本提炼为简洁的摘要
- **AI 助手**：智能问答功能

### 1.3 PDF 文档处理
- **PDF合并**：合并指定目录下的所有PDF文件
- **PDF页面提取**：从PDF文件中提取指定页面范围

// ... 其他内容保持不变 ...

### 3.2 AI 文本处理
```bash
# 文本翻译（支持中英互译）
deskAI.exe translate [-model 模型名] "要翻译的文本"

# 文本润色
deskAI.exe polish [-model 模型名] "要润色的文本"

# 文本总结
deskAI.exe summarize [-model 模型名] "要总结的文本"

# AI 助手问答
deskAI.exe ask [-model 模型名] "你的问题"
```

## 2. 安装配置

### 2.1 环境要求
- Go 1.24.1 或更高版本
- Windows 操作系统

### 2.2 配置文件
在程序运行目录下创建 `config.json` 文件：

```json
{
    "dinox": {
        "token": "your_dinox_token"
    },
    "models": {
        "qwen-plus": {
            "url": "模型接口地址",
            "token": "访问令牌",
            "model": "模型名称",
            "temperature": 0.7
        }
    }
}
```
## 3. 使用方法

### 3.1 基本命令
```bash
# 查看帮助信息
deskAI.exe help

# 发送内容到 DinoAI
deskAI.exe dinoxPost "要发送的内容"
 ```

### 3.2 AI 文本处理
```bash
# 文本翻译（支持中英互译）
deskAI.exe translate [-model 模型名] "要翻译的文本"

# 文本润色
deskAI.exe polish [-model 模型名] "要润色的文本"

# 文本总结
deskAI.exe summarize [-model 模型名] "要总结的文本"

# AI 助手问答
deskAI.exe ask [-model 模型名] "你的问题"
 ```

### 3.3 参数说明
- -model ：指定使用的 AI 模型，默认为 "qwen-plus"
- 所有文本内容需要用引号包围
```
# 合并PDF文件
deskAI.exe pdfMerge -dir "PDF文件目录"

# 提取PDF页面
deskAI.exe pdfExtract -input "输入目录" [-output "输出目录"] -pages "1,3-5"
```
### 3.4 参数说明
- -model：指定使用的 AI 模型，默认为 "qwen-plus"
- -dir：指定要处理的PDF文件目录
- -input：PDF提取功能的输入目录
- -output：PDF提取功能的输出目录（可选，默认为输入目录下的output文件夹）
- -pages：要提取的页码范围，支持如"1,3-5"的格式
- 所有文本内容需要用引号包围

## 4. 示例
```bash
# 翻译示例
deskAI.exe translate "Hello, how are you?"

# 使用指定模型进行翻译
deskAI.exe translate -model gpt-4 "Hello, how are you?"

# 文本润色示例
deskAI.exe polish "这是一段需要优化的文本"

# 总结长文本
deskAI.exe summarize "这是一段很长的文本..."

# 向 AI 助手提问
deskAI.exe ask "什么是人工智能？"

# 合并PDF文件
deskAI.exe pdfMerge -dir "D:\Documents\PDFs"

# 提取PDF页面
deskAI.exe pdfExtract -input "D:\Documents\PDFs" -output "D:\Documents\Output" -pages "1,3-5"
 ```
## 5. 注意事项
1. 确保 config.json 配置文件正确设置
2. 文本内容包含空格时必须使用引号
3. 模型名称区分大小写
4. 程序返回错误时，请检查：
   - 配置文件是否正确
   - 网络连接是否正常
   - API 令牌是否有效
## 6. 常见问题
### 6.1 配置文件找不到
- 确保 config.json 与可执行文件在同一目录
- 检查文件权限
### 6.2 API 调用失败
- 验证 token 是否正确
- 检查网络连接
- 确认 API 服务是否可用
### 6.3 命令执行错误
- 检查命令格式是否正确
- 确保参数完整且格式正确
- 文本内容是否正确使用引号包围
## 7. 更新日志
### 1.1.0
- 新增 PDF 文件处理功能
  - 支持合并多个PDF文件
  - 支持提取PDF指定页面
### v1.0.0
- 支持 DinoAI 笔记发送
- 集成 AI 文本处理功能
- 支持多种 AI 模型配置

```plaintext
这些更新添加了：
1. PDF处理功能的介绍
2. PDF相关命令的使用说明
3. 新的参数说明
4. PDF处理的使用示例
5. 更新日志中添加了新版本信息
```