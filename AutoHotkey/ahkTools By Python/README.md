# ahkTools Python
> 这是一个早先的基础版本，使用的是 AutoHotkey v2.0和Python开发，后续更新，使用AutoHotkey v2.0和Go语言进行了重构，增加了很多新功能。

一个基于 AutoHotkey v2.0 和 Python 的 AI 文本处理工具，支持文本翻译、润色等功能。

## 功能特点

- 支持文本翻译（中英互译）
- 支持文本润色
- 支持快捷键操作
- 系统托盘运行
- GUI 界面操作
- 自动化文本处理

## 系统要求

- Windows 操作系统
- AutoHotkey v2.0
- Python 3.x
- 相关 Python 依赖包

## 安装步骤

1. 安装 [AutoHotkey v2.0](https://www.autohotkey.com/)
2. 安装 [Python](https://www.python.org/downloads/)
3. 安装依赖包：
   ```bash
   pip install -r requirements.txt
   ```
4. 配置 config.ini 文件：
   - 设置热键
   - 配置 API 参数
   - 设置 GUI 界面参数

## 使用方法

### 启动程序

双击运行 `ahkTools Python.ahk` 文件，程序将在系统托盘中运行。

### 快捷键

- `Alt + T`：翻译选中文本（可在配置文件中修改）
- `Alt + P`：润色选中文本（可在配置文件中修改）
- `Win + N`：快速启动/切换 Sublime Text
- `Ctrl + Alt + Q`：退出程序

### 文本替换

- `cr`：插入带有时间标记的文本
- `email`/`8151`/`qq`：自动填充邮箱地址
- `wc`：输入完成标记 ✅
- `db`：输入待办标记 ⏳
- `rq`：输入当前日期（YYYYMMDD格式）

## 配置文件说明

config.ini 文件包含以下配置项：

```ini
[Hotkey]
translate = !t    ; Alt + T
polish = !p      ; Alt + P

[GUI]
width = 400
height = 300

[API]
# API 相关配置
```

## 文件说明

- `ahkTools Python.ahk`：主程序文件
- `ai_tool.py`：Python AI 处理模块
- `config.ini`：配置文件
- `result.txt`：临时结果文件（程序运行时自动创建和删除）
- `ai_tool.log`：日志文件

## 注意事项

1. 确保已正确配置 API 相关参数
2. 使用快捷键前确保已选中要处理的文本
3. 程序退出时会自动清理临时文件

## 作者

Lei

## 版本历史

- v1.0 (2025-03-28)
  - 初始版本发布
  - 实现基本的文本处理功能
  - 添加 GUI 界面
```

这个 README.md 文件包含了项目的主要信息，包括：
1. 功能介绍
2. 安装要求
3. 使用方法
4. 配置说明
5. 文件结构
6. 注意事项等

你可以根据实际需求调整或补充其中的内容。