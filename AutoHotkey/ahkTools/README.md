我来帮你编写一个完整的说明文件：

```markdown:f:\Code\EfficientCode\AutoHotkey\ahkTools\README.md
# AI文本工具使用说明

## 项目简介
AI文本工具（ahkTools）是一个基于 AutoHotkey v2.0 的智能文本处理工具，集成了 OpenAI 服务，提供文本翻译、润色等 AI 功能，同时包含多个实用的快捷输入功能。

## 系统要求
- Windows 操作系统
- AutoHotkey v2.0
- Go 1.20 或更高版本
- 有效的 OpenAI API 密钥

## 安装步骤
1. 安装 AutoHotkey v2.0
2. 安装 Go 环境
3. 克隆或下载本项目
4. 编译 AI 工具：
   ```bash
   cd ahkTools
   go build -o ai_tool.exe ai_tool.go
   ```
5. 配置 config.ini 文件

## 配置说明
### 配置文件：config.ini
```ini
[API]
base_url = https://api.openai.com/v1
api_key = your_api_key_here
model = gpt-3.5-turbo
timeout = 30

[Hotkey]
translate = !t    ; Alt+T：翻译
polish = !p      ; Alt+P：润色

[GUI]
width = 400      ; 界面宽度
height = 300     ; 界面高度

[Functions]
translate = {"system_prompt":"你是一个专业的翻译助手","user_prompt":"将文本翻译成{target_lang}：\n\n{text}","temperature":0.3,"stream":false}
polish = {"system_prompt":"你是一个专业的文字润色助手","user_prompt":"润色以下文本：\n\n{text}","temperature":0.7,"stream":false}
```

## 功能说明

### 1. AI 文本处理
#### 文本翻译
- 热键：Alt+T（可自定义）
- 使用方法：
  1. 选中要翻译的文本
  2. 按下热键
  3. 或打开主界面，粘贴文本后点击"翻译"按钮

#### 文本润色
- 热键：Alt+P（可自定义）
- 使用方法：
  1. 选中要润色的文本
  2. 按下热键
  3. 或打开主界面，粘贴文本后点击"润色"按钮

### 2. 快捷工具

#### Sublime Text 快速启动
- 热键：Win+N
- 功能：
  - 已运行：激活窗口
  - 最小化：恢复窗口
  - 未运行：启动新实例

#### 语音合成标记插入
- 触发：输入 "cr"
- 功能：自动在标点符号后添加语音合成时间标记
  - 句号(./。)后：`((⏱️=4000))`
  - 其他标点(,，;:?!)后：`((⏱️=2000))`
- 使用方法：
  1. 复制需要处理的文本
  2. 输入 "cr"
  3. 自动插入时间标记

#### 快捷输入
- 邮箱地址
  - `email` → 815141681@qq.com
  - `8151` → 815141681@qq.com
  - `qq` → 815141681@qq.com

- 任务状态标记
  - `wc` → ✅（完成标记）
  - `db` → ⏳（待办标记）

- 日期输入
  - `rq` → 当前日期（YYYYMMDD格式）

### 3. 界面功能
- 系统托盘
  - 启动后最小化到托盘
  - 双击托盘图标：显示主窗口
  - 右键菜单：显示主窗口/退出

- 主界面
  - 文本输入框
  - 结果输出框
  - 功能按钮：翻译/润色/复制
  - 窗口关闭：最小化到托盘

### 4. 退出程序
- 热键：Ctrl+Alt+Q
- 或通过托盘菜单退出

## 常见问题
1. 翻译/润色失败
   - 检查 API 密钥是否正确
   - 确认网络连接正常
   - 查看 ai_tool.log 日志文件

2. 热键无效
   - 确认热键格式正确（如 !t, ^p）
   - 检查是否与其他程序热键冲突
   - 重启程序尝试

3. 文本获取失败
   - 确保已选中文本
   - 检查目标窗口是否允许复制
   - 确认剪贴板未被其他程序占用

## 更新日志

### V0.2.0
- 使用 Go 语言重写 AI 服务调用
- 优化配置文件格式
- 改进错误处理和日志记录

### V0.1.0
- 初始版本发布
- 实现基本的翻译和润色功能
- 添加快捷输入功能
- 实现系统托盘支持