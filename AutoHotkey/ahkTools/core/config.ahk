#Requires AutoHotkey v2.0
#Include "hotkey.ahk"


; 全局配置变量
global g_configFile := "config.ini"
global g_translateHotkey := ""
global g_polishHotkey := ""
global g_guiWidth := 400
global g_guiHeight := 300

; 初始化配置
InitConfig() {
    global g_translateHotkey := IniRead(g_configFile, "Hotkey", "translate", "!t")
    global g_polishHotkey := IniRead(g_configFile, "Hotkey", "polish", "!p")
    global g_dinoxHotkey := IniRead(g_configFile, "Hotkey", "dinoxHotkey", "!a")
    global g_guiWidth := IniRead(g_configFile, "GUI", "width", "400")
    global g_guiHeight := IniRead(g_configFile, "GUI", "height", "300")
    
    ; 确保热键格式正确
    global g_translateHotkey := FormatHotkey(g_translateHotkey)
    global g_polishHotkey := FormatHotkey(g_polishHotkey)
}

