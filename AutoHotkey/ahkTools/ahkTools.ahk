#Requires AutoHotkey v2.0
#SingleInstance Force
#Warn
FileEncoding "UTF-8"

; 导入模块
#Include "core\config.ahk"
#Include "gui\gui.ahk"
#Include "gui\tray.ahk"
#Include "core\hotkey.ahk"
#Include "utils\text.ahk"
#Include "utils\dinox.ahk"
#Include "core\sublime.ahk"

; 初始化全局变量
InitGlobalVars() {
    global g_workingDir := A_ScriptDir
    global g_debug := true
    global g_codePage := "CP65001"
    global g_logDir := g_workingDir "\logs"
}

; 初始化日志系统
InitLogging() {
    if !DirExist(g_logDir)
        DirCreate(g_logDir)
}

; 日志函数
LogDebug(message) {
    if g_debug {
        FileAppend FormatTime() ": " message "`n", g_logDir "\debug.log"
    }
}

; 检查必要文件
CheckRequiredFiles() {
    if !FileExist("config.ini") {
        MsgBox "配置文件不存在，请确保 config.ini 文件在正确位置", "错误", "48"
        ExitApp
    }
}

; 设置系统编码
SetSystemEncoding() {
    DllCall("SetConsoleOutputCP", "Int", 65001)
}

; 主程序初始化
InitApplication() {
    try {
        InitGlobalVars()
        InitLogging()
        CheckRequiredFiles()
        SetSystemEncoding()
        
        InitConfig()
        InitGui()
        InitTray()
        InitHotkey()
        
        LogDebug("应用程序初始化成功")
    } catch Error as initError {
        LogDebug("初始化失败: " initError.Message)
        MsgBox "初始化失败: " initError.Message, "错误", "48"
        ExitApp
    }
}

; 启动应用程序
InitApplication()