#Requires AutoHotkey v2.0
; 全局变量
global g_guiWidth := 400

; 全局GUI变量
global g_mainGui := ""
global g_editSource := ""
global g_editTarget := ""

; 初始化GUI
InitGui() {
    global g_mainGui := Gui()
    g_mainGui.Icon := "*"
    g_mainGui.Icon := A_ScriptDir "/ai.ico"
    g_mainGui.SetFont("s10", "Microsoft YaHei")
    g_mainGui.Title := "AI文本工具"

    ; 创建水平布局
    g_mainGui.Add("Text", , "输入文本:")
    inputGroup := g_mainGui.Add("GroupBox", "w" g_guiWidth + 5 " h120")
    global g_editSource := g_mainGui.Add("Edit", "xp+10 yp+20 w" g_guiWidth - 10 " h90")
    btnClearSource := g_mainGui.Add("Button", "x+5 yp w24 h24", "🗑️")
    
    g_mainGui.Add("Text", "xm", "处理结果:")
    resultGroup := g_mainGui.Add("GroupBox", "w" g_guiWidth + 5 " h120")
    global g_editTarget := g_mainGui.Add("Edit", "xp+10 yp+20 w" g_guiWidth - 10 " h90 ReadOnly VScroll")
    btnCopyTarget := g_mainGui.Add("Button", "x+5 yp w24 h24", "📋")

    ; 添加功能按钮
    btnTranslate := g_mainGui.Add("Button", "xm Default", "翻译")
    btnPolish := g_mainGui.Add("Button", "x+10", "润色")
    btnDinox := g_mainGui.Add("Button", "x+10", "Dinox")

    ; 注册事件
    btnTranslate.OnEvent("Click", (*) => ProcessText("translate"))
    btnPolish.OnEvent("Click", (*) => ProcessText("polish"))
    btnDinox.OnEvent("Click", ProcessDinox)
    btnClearSource.OnEvent("Click", ClearSource)
    btnCopyTarget.OnEvent("Click", CopyTarget)

    ; 添加复制按钮提示
    btnClearSource.ToolTip := "清空输入文本"
    btnCopyTarget.ToolTip := "复制处理结果"

    ; 设置窗口关闭事件
    g_mainGui.OnEvent("Close", (*) => g_mainGui.Hide())

    ; 初始化时最小化到托盘
    g_mainGui.Show("AutoSize Hide")
}

; 显示主窗口函数
ShowMainWindow(*) {
    g_mainGui.Show()
    WinActivate("AI文本工具")
}

; 复制函数
CopySource(*) {
    A_Clipboard := g_editSource.Value
}
; 清空函数
ClearSource(*) {
    g_editSource.Value := ""
}
CopyTarget(*) {
    A_Clipboard := g_editTarget.Value
}
