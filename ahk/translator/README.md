# AI文本工具

一个基于AutoHotkey v2.0和AI服务的文本处理工具，支持文本翻译、润色等功能。

## 功能特点

- 支持文本翻译和润色功能
- 全局快捷键支持
- 系统托盘图标
- 可自定义AI服务配置
- 界面简洁易用

## 安装要求

- AutoHotkey v2.0
- Python 3.x
- 有效的AI服务API密钥

## 配置说明

### config.ini 配置文件

```ini
[API]
base_url = 你的API基础URL
api_key = 你的API密钥
model = 使用的AI模型
timeout = 请求超时时间（秒）

[Hotkey]
translate = !t  # Alt+T
polish = ^p    # Ctrl+P

[GUI]
width = 400    # 界面宽度
height = 300   # 界面高度
```

## 使用方法

1. 运行translator.ahk启动程序
2. 程序会最小化到系统托盘
3. 使用以下方式处理文本：
   - 通过快捷键：选中文本后按快捷键
   - 通过界面：点击托盘图标打开主界面

### 快捷键

- Alt+T：翻译选中的文本
- Ctrl+P：润色选中的文本

### 界面操作

1. 在输入框中输入或粘贴文本
2. 点击相应功能按钮（翻译/润色）
3. 结果会显示在下方文本框
4. 可以使用复制按钮快速复制结果

## 注意事项

- 确保config.ini中配置了正确的API信息
- 使用快捷键功能时，需要先选中要处理的文本
- 程序会在后台保持运行，可以通过托盘图标管理