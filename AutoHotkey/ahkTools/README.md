# ahkTools

ahkTools是一个基于AutoHotkey v2.0、Go语言开发的AI文本处理工具，提供了便捷的文本翻译、润色等功能，支持全局热键和图形界面操作。

## 主要功能

### 1. AI文本处理

- **文本翻译**
  - 支持全局热键快速调用（默认Alt+T）
  - 支持GUI界面操作
  - 自动识别源语言
  - 支持多种目标语言选择

- **文本润色**
  - 支持全局热键快速调用（默认Alt+P）
  - 支持GUI界面操作
  - 智能优化文本表达
  - 提供多种润色风格
- **DinoAI笔记**
  - 支持全局热键快速调用（默认Alt+D）
  - 支持GUI界面操作
  - 一键发送文本到DinoAI
  - 支持命令行调用

### 2. 便捷操作

- 全局热键支持
- 简洁的图形界面
- 系统托盘图标
- 快速配置选项
- 命令行工具支持

## 系统要求

- Windows 10/11
- AutoHotkey v2.0及以上版本
- 联网环境（用于AI服务调用）

## 安装说明

1. 安装[AutoHotkey v2.0](https://www.autohotkey.com/)
2. 下载本项目代码
3. 运行`ahkTools.ahk`主程序
4. 配置`config.json`中的DinoAI token

## 配置说明

### AI服务配置

在`config.json`和`config.ini`中配置：

- API密钥
- DinoAI Token
- 服务地址
- 其他AI相关参数

### 热键配置

在`config.ini`中可自定义：

- 翻译热键（默认Alt+T）
- 润色热键（默认Alt+P）
- DinoAI热键（默认Alt+D）

## 使用方法

1. **文本翻译**
   - 选中需要翻译的文本
   - 按下Alt+T（或自定义热键）
   - 或打开GUI界面，粘贴文本后点击"翻译"按钮

2. **文本润色**
   - 选中需要润色的文本
   - 按下Alt+P（或自定义热键）
   - 或打开GUI界面，粘贴文本后点击"润色"按钮
3. **DinoAI笔记**
   - 选中需要发送的文本
   - 按下Alt+D（或自定义热键）
   - 或打开GUI界面，粘贴文本后点击"Dinox"按钮
   - 或使用命令行：`.\deskAI.exe dinoxPost "你的笔记内容"`(PowerShell)或`deskAI.exe dinoxPost "你的笔记内容"`

## 项目结构

```
ahkTools.ahk     # 主程序入口
core/            # 核心功能模块
  config.ahk     # 配置管理
  hotkey.ahk     # 热键管理
gui/             # 界面相关
  gui.ahk        # 主界面
  tray.ahk       # 托盘图标
utils/           # 工具函数
  text.ahk       # 文本处理
```

## 贡献指南

欢迎提交Issue和Pull Request来帮助改进项目。在提交代码时请确保：

1. 遵循现有的代码风格
2. 添加必要的注释和文档
3. 测试新功能或修复的bug

## 许可证

本项目采用MIT许可证。