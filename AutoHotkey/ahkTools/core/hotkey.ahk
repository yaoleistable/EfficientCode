#Requires AutoHotkey v2.0

; 转换热键格式
FormatHotkey(key) {
    static validModifiers := Map(
        "!", "Alt",
        "^", "Ctrl",
        "+", "Shift",
        "#", "Win"
    )

    ; 转换为小写
    key := StrLower(key)

    ; 验证格式
    if (StrLen(key) < 2)
        return ""

    ; 检查修饰键和键名
    modifier := SubStr(key, 1, 1)
    keyName := SubStr(key, 2)

    if !validModifiers.Has(modifier) || keyName = ""
        return ""

    return key
}

; 初始化热键
InitHotkey() {
    ; 确保热键格式正确
; 使用全局配置变量
global g_translateHotkey
global g_polishHotkey
global g_dinoxHotkey
    ; 注册热键
    if (g_translateHotkey || g_polishHotkey || g_dinoxHotkey) {
        HotIfWinNotActive "AI文本工具"
        try {
            if (g_translateHotkey)
                Hotkey g_translateHotkey, (*) => ShowTool("translate")
            if (g_polishHotkey)
                Hotkey g_polishHotkey, (*) => ShowTool("polish")
            if (g_dinoxHotkey)
                Hotkey g_dinoxHotkey, (*) => ShowTool("dinox")
        } catch Error as hotkeyError {
            MsgBox Format("热键注册失败: {}`n配置的热键: {}, {}, {}",
                hotkeyError.Message, g_translateHotkey, g_polishHotkey, g_dinoxHotkey)
        }
    } else {
        MsgBox "热键格式无效，请检查配置文件。`n应使用如 !d (Alt+D) 或 ^d (Ctrl+D) 的格式。"
    }
}




; 显示工具
ShowTool(function, *) {
    ; 保存当前活动窗口句柄
    sourceWin := WinExist("A")
    
    ; 检查窗口状态并确保正确显示
    if !WinExist("AI文本工具") {
        g_mainGui.Show()
    }
    
    ; 获取选中的文本
    savedClip := ClipboardAll()  ; 保存完整的剪贴板内容
    A_Clipboard := ""  ; 清空剪贴板
    
    try {
        ; 确保源窗口处于活动状态
        if (sourceWin) {
            WinActivate("ahk_id " sourceWin)
            if !WinWaitActive("ahk_id " sourceWin, , 1) {
                throw Error("无法激活源窗口")
            }
            Sleep(100)  ; 等待窗口激活
        }
        
        ; 发送复制命令并等待
        BlockInput("On")  ; 临时阻止用户输入
        Send "^c"
        BlockInput("Off")
        Sleep(200)  ; 等待复制命令执行
        
        ; 等待剪贴板内容，最多重试5次
        success := false
        loop 5 {
            if ClipWait(0.5) {
                text := A_Clipboard
                if (text != "") {
                    success := true
                    break
                }
            }
            Sleep(300)  ; 增加等待时间
        }
        
        ; 激活主窗口并处理文本
        WinActivate("AI文本工具")
        if !WinWaitActive("AI文本工具", , 1) {
            throw Error("无法激活主窗口")
        }
        
        if (success) {
            g_editSource.Value := text
            if (function = "dinox") {
                ProcessDinox()  ; 直接调用 Dinox 处理函数
            } else {
                ProcessText(function)
            }
        } else {
            errorMsg := "无法获取选中的文本。请确保："
            errorMsg .= "`n1. 已选中要处理的文本"
            errorMsg .= "`n2. 当前窗口允许复制操作"
            errorMsg .= "`n3. 没有其他程序占用剪贴板"
            errorMsg .= "`n4. 选中的文本不为空"
            MsgBox(errorMsg, "错误", "48")
        }
    } finally {
        ; 恢复原始剪贴板内容
        A_Clipboard := savedClip
    }
}