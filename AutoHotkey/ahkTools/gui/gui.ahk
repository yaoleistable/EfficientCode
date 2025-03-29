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

    ; 添加控件
    g_mainGui.Add("Text", , "输入文本:")
    global g_editSource := g_mainGui.Add("Edit", "vSource w" g_guiWidth " h60")
    g_mainGui.Add("Text", , "处理结果:")
    global g_editTarget := g_mainGui.Add("Edit", "vTarget w" g_guiWidth " h60")

    ; 添加功能按钮
    btnTranslate := g_mainGui.Add("Button", "Default", "翻译")
    btnPolish := g_mainGui.Add("Button", , "润色")
    btnCopySource := g_mainGui.Add("Button", , "复制输入")
    btnCopyTarget := g_mainGui.Add("Button", , "复制结果")

    ; 注册事件
    btnTranslate.OnEvent("Click", (*) => ProcessText("translate"))
    btnPolish.OnEvent("Click", (*) => ProcessText("polish"))
    btnCopySource.OnEvent("Click", CopySource)
    btnCopyTarget.OnEvent("Click", CopyTarget)

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

CopyTarget(*) {
    A_Clipboard := g_editTarget.Value
}