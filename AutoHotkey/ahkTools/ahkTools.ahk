#Requires AutoHotkey v2.0

; 作者：Lei
; 日期：2025-03-28
; 版本：0.2.0
; 功能：AI文本工具一个基于AutoHotkey v2.0和AI服务的文本处理工具，支持文本翻译、润色等功能。

#SingleInstance Force
#Warn
FileEncoding "UTF-8"  ; 设置文件操作的默认编码为UTF-8

; 导入模块
#Include "core\\config.ahk"
#Include "gui\gui.ahk"
#Include "gui\tray.ahk"
#Include "core\hotkey.ahk"
#Include "utils\text.ahk"
#Include "utils\dinox.ahk"
#Include "core\sublime.ahk"

; 全局变量
global g_workingDir := A_ScriptDir
global g_debug := true
global g_codePage := "CP65001"  ; 添加 UTF-8 编码页设置

; 创建日志目录
global g_logDir := g_workingDir "\logs"
if !DirExist(g_logDir)
    DirCreate(g_logDir)

; 日志函数
LogDebug(message) {
    if g_debug {
        FileAppend FormatTime() ": " message "`n", g_logDir "\debug.log"
    }
}

; 确保配置文件存在
if !FileExist("config.ini") {
    MsgBox "配置文件不存在，请确保 config.ini 文件在正确位置", "错误", "48"
    ExitApp
}

; 初始化应用程序
try {
    ; 设置控制台输出编码为 UTF-8
    DllCall("SetConsoleOutputCP", "Int", 65001)
    
    InitConfig()
    InitGui()
    InitTray()
    InitHotkey()
} catch Error as initError {
    MsgBox "初始化失败: " initError.Message, "错误", "48"
    ExitApp
}