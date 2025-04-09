### Go语言编译过程
```bash
cd f:\Code\EfficientCode\AutoHotkey\ahkTools
go mod tidy
# go mod init ahktools
go get gopkg.in/ini.v1
go build -o ai_tool.exe ai_tool.go
# 优化编译文件大小，减小go程序的体积
go build -ldflags="-s -w" -o ai_tool.exe ai_tool.go 

go build -o deskAI.exe deskAI.go
go build -ldflags="-s -w" -o deskAI.exe deskAI.go
```


### TODO:
- 调用github.com/tmc/langchaingo的LLM，实现了调用gemini兼容的API请求
### V0.2.15
- 调用github.com/tmc/langchaingo的LLM，实现了调用openai兼容的API请求。

### V0.2.14
- 优化了代码结果，便于后续维护。
### V0.2.13
- deskAI增加了PDF提取和合并功能。
### V0.2.12
- 邮箱快捷输入，通过读取config.ini来实现。

### V0.2.11
- 删除ai_tool.go，已在deskAI中实现。
- 优化程序总体体积。

### V0.2.10
- 优化ai_tool.go文件中io处理相关文件。

### V0.2.9
- 优化sublime快捷方式的获取方式，在不同电脑都很容易实现打开应用。
- 调整sublime为单独代码块。

### V0.2.8
- 修改发送到Dinox的提示。
- 修改发送到Dinox默认快捷键为ALT+A。
- 完善了程序说明文件。

### V0.2.7
- 修复deskAI读取配置文件失败的问题。

### V0.2.6
- 快捷键alt+d实现发送内容到Dinox。
- 点击按钮时清空上次的内容。

### V0.2.5
- 优化了代码结构，提高了代码的可读性和可维护性。
- 增加deskAI，实现go语言发送文本到dinox。


### V0.2.4
- 优化显示GUI界面

### V0.2.3
- 优化了代码结构，提高了代码的可读性和可维护性。
- 增加了托盘自定义图标显示。

### V0.2.2
- 优化ai_tool.exe文件大小，减小go程序的体积
### V0.2.1
- 修复调用Sublime Text文件路径有误的问题
```bash
# 使用环境变量获取开始菜单程序路径
startMenuPath := EnvGet("APPDATA") "\Microsoft\Windows\Start Menu\Programs\Sublime Text.lnk"
```

### V0.2.0
- 使用Go语言调用openai接口，实现文本翻译和润色功能
    - Go 版本的 AI 工具 :
    - 使用 gopkg.in/ini.v1 包解析 INI 配置文件
    - 实现了与 Python 版本相同的功能：翻译、润色、摘要等
    - 支持日志记录
    - 处理 API 请求和响应
- 配置文件 :
    - 保持与 Python 版本相同的结构
    - 使用 JSON 格式存储功能配置
- AHK 脚本修改 :
    - 只需修改 ProcessText 函数，将调用从 Python 脚本改为 Go 可执行文件

### V0.2.0
- 使用Python调用openai接口，实现文本翻译和润色功能
- 增加了Sublime Text快速启动功能
- 增加了任务状态标记功能
- 增加了日期输入功能