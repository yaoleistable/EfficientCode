# ahkTools使用说明

ahkTools是一个基于AutoHotkey v2.0和Go语言开发的AI文本处理工具，提供了便捷的文本翻译、润色等功能，支持全局热键和图形界面操作。

## 主要功能

### 1. AI文本处理

- **文本翻译**
  - 支持全局热键快速调用（默认Alt+T）
  - 支持GUI界面操作
  - 自动识别中英文
  - 中英互译

- **文本润色**
  - 支持全局热键快速调用（默认Alt+P）
  - 支持GUI界面操作
  - 智能优化文本表达

- **DinoAI笔记**
  - 支持全局热键快速调用（默认Alt+A）
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
3. 编译 deskAI 目录下的 Go 程序
4. 运行`ahkTools.ahk`主程序
5. 配置`deskAI\config.json`中Dinox Token、AI Token
6. 配置`config.ini`中的设置

## 配置说明

### config.ini 配置

```ini
[Hotkey]
translate = !t
polish = !p
dinoxHotkey = !a

[GUI]
width = 400
height = 300
```

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
   - 按下Alt+A（或自定义热键）
   - 或打开GUI界面，粘贴文本后点击"Dinox"按钮
   - 或使用命令行：`deskAI.exe dinoxPost "你的笔记内容"`（PowerShell中需加`.\deskAI.exe ...`）

## 项目结构

```
## 项目结构

ahkTools/                   # 项目根目录
├─ ahkTools.ahk             # 主程序入口
├─ config.ini               # 配置文件
├─ deskAI                   # DinoAI相关
│  └─ ai                    # deskAI功能实现（Go）
│     ├─ ai.go              # ai文本调用
│     └─ config.go          # 配置文件处理
│     └─ text_utils.go      # ai翻译、总结等具体功能实现
│  └─ dinox                 # Dinox相关功能
│     └─ dinox.go           # 调用dinox相关功能
│  └─ deskAI.go             # 主程序入口
├─ core                     # 核心功能模块
│  ├─ config.ahk            # 配置管理
│  ├─ hotkey.ahk            # 热键管理
│  └─ sublime.ahk           # Sublime快捷键
├─ gui                      # 界面相关
│  ├─ gui.ahk               # 主界面
│  └─ tray.ahk              # 托盘图标
└─ utils                    # 工具函数
   ├─ text.ahk              # 文本处理，调用deskAI.exe
   └─ dinox.ahk             # DinoAI处理,调用deskAI.exe
```

## 开发说明

- AutoHotkey v2.0 负责界面和热键处理
- Go 语言实现 AI 功能调用
- 使用 INI 文件进行配置管理
- 模块化设计，便于扩展

## 许可证

本项目采用MIT许可证。
```

主要更新：
1. 简化了功能描述，更符合实际实现
2. 更新了安装步骤，添加了Go程序编译说明
3. 添加了具体的配置文件示例
4. 更新了项目结构，反映当前的文件组织
5. 添加了开发说明部分
6. 移除了过时的功能描述
7. 统一了热键说明（DinoAI热键改为Alt+A）
8. 简化了命令行示例