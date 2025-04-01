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

; 确保配置文件存在
if !FileExist("config.ini") {
    MsgBox "配置文件不存在，请确保 config.ini 文件在正确位置", "错误", "48"
    ExitApp
}

; 初始化应用程序
try {
    InitConfig()
    InitGui()
    InitTray()
    InitHotkey()
} catch Error as initError {
    MsgBox "初始化失败: " initError.Message, "错误", "48"
    ExitApp
}