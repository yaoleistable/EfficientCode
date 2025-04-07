#Requires AutoHotkey v2.0  ; 指定需要 AutoHotkey v2.0 版本
#Include "hotkey.ahk"      ; 导入热键处理模块

; 全局配置变量定义
global g_configFile := "config.ini"    ; 配置文件路径
global g_translateHotkey := ""         ; 翻译功能的热键
global g_polishHotkey := ""           ; 润色功能的热键
global g_guiWidth := 400              ; GUI 窗口宽度
global g_guiHeight := 300             ; GUI 窗口高度

; 初始化配置函数
InitConfig() {
    ; 从配置文件读取热键设置，如果没有配置则使用默认值
    global g_translateHotkey := IniRead(g_configFile, "Hotkey", "translate", "!t")    ; Alt+T
    global g_polishHotkey := IniRead(g_configFile, "Hotkey", "polish", "!p")         ; Alt+P
    global g_dinoxHotkey := IniRead(g_configFile, "Hotkey", "dinoxHotkey", "!a")     ; Alt+A
    
    ; 从配置文件读取 GUI 窗口尺寸设置
    global g_guiWidth := IniRead(g_configFile, "GUI", "width", "400")                 ; 默认宽度 400
    global g_guiHeight := IniRead(g_configFile, "GUI", "height", "300")               ; 默认高度 300
    
    ; 确保热键格式正确，通过 FormatHotkey 函数处理热键字符串
    ; FormatHotkey 函数在 hotkey.ahk 中定义，用于标准化热键格式
    global g_translateHotkey := FormatHotkey(g_translateHotkey)                       ; 格式化翻译热键
    global g_polishHotkey := FormatHotkey(g_polishHotkey)                            ; 格式化润色热键
}

