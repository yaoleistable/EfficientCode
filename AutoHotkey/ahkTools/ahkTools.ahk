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

; 初始化应用程序
InitConfig()
InitGui()
InitTray()
InitHotkey()