#Requires AutoHotkey v2.0
#Include "..\core\config.ahk"

; 字符串插入功能的常量定义
B := "((⏱️=2000))"  ; 要添加在其他标点符号后的字符串B
C := "((⏱️=4000))"  ; 要添加在句号后的字符串C

; 从配置文件读取邮箱
global g_userEmail := IniRead(g_configFile, "User", "email", "")

; 热字串定义 - 自动化输入，快捷输入邮箱
:C:email::
{
    SendText(g_userEmail)
}
:C:qq::
{
    SendText(g_userEmail)
}

; 完成
::wc::✅
; 待办
::db::⏳

; 输入 rq 自动生成当天日期（格式：YYYYMMDD）
::rq::
{
    ; 获取当前日期并格式化
    currentDate := FormatTime(, "yyyyMMdd")
    
    ; 发送格式化后的日期
    SendText(currentDate)
}

; 热字串触发 - 字符串插入功能
::cr::
{
    ; 保存当前剪贴板内容
    ClipSaved := ClipboardAll()
    
    ; 获取剪贴板文本内容
    TextContent := A_Clipboard

    if !TextContent
    {
        MsgBox "剪贴板为空或无法访问。"
        return
    }

    ; 1. 在句号 (英文和中文) 后添加字符串C
    ModifiedText := RegExReplace(TextContent, "([.。])", "$1" . C)

    ; 2. 在其他标点符号后添加字符串B
    ModifiedText := RegExReplace(ModifiedText, "([,，;:?!])", "$1" . B)

    ; 发送修改后的文本
    SendText(ModifiedText)

    ; 恢复原始剪贴板内容
    A_Clipboard := ClipSaved

    ; 清空变量，释放内存
    TextContent := ""
    ModifiedText := ""
    ClipSaved := ""
}

; 处理文本
ProcessText(function) {
    ; 清空目标文本框
    g_editTarget.Value := ""
    text := g_editSource.Value
    if (text = "")
        return

    deskAIPath := g_workingDir "\deskAI\deskAI.exe"
    
    if !FileExist(deskAIPath) {
        g_editTarget.Value := "错误：找不到 deskAI.exe，请确保程序已正确编译"
        return
    }

    try {
        ; 设置命令
        cmd := Format('"{1}" {2} "{3}"', deskAIPath, function, text)
        
        g_editTarget.Value := "正在处理文本...`n"
        
        ; 创建临时管道文件
        tempFile := A_Temp "\ahk_output.txt"
        
        ; 使用 Run 执行命令并重定向输出
        RunWait(A_ComSpec ' /c chcp 65001 >nul && ' cmd ' > "' tempFile '"',, "Hide")
        
        ; 读取输出文件
        if FileExist(tempFile) {
            output := FileRead(tempFile, "UTF-8")
            g_editTarget.Value := output
            FileDelete(tempFile)
        } else {
            g_editTarget.Value := "处理完成，但无输出结果"
        }
        
    } catch as err {
        g_editTarget.Value := "执行错误: " err.Message
        LogDebug("执行错误: " err.Message)
    }
}