#Requires AutoHotkey v2.0

; 字符串插入功能的常量定义
B := "((⏱️=2000))"  ; 要添加在其他标点符号后的字符串B
C := "((⏱️=4000))"  ; 要添加在句号后的字符串C

; 热字串定义 - 自动化输入
:C:email::815141681@qq.com
:C:8151::815141681@qq.com
:C:qq::815141681@qq.com
::wc::✅
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
    text := g_editSource.Value
    if (text = "")
        return

    ; 修改为调用Go程序
    result := RunWait(Format('ai_tool.exe "{}" "{}"', function, text), , "Hide")
    if FileExist("result.txt") {
        try {
            processed := FileRead("result.txt", "UTF-8")
            g_editTarget.Value := processed
        } catch Error as processError {
            g_editTarget.Value := "处理文本时出错: " processError.Message
        }
        FileDelete("result.txt")
    }
}