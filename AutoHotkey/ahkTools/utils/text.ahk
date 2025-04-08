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

    ; 使用全局工作目录构建正确的路径
    deskAIPath := g_workingDir "\deskAI\deskAI.exe"
    
    ; 记录调试信息
    LogDebug("执行路径: " deskAIPath)
    LogDebug("功能: " function)
    LogDebug("文本: " text)
    
    ; 检查文件是否存在
    if !FileExist(deskAIPath) {
        LogDebug("错误：deskAI.exe 不存在")
        g_editTarget.Value := "错误：找不到 deskAI.exe，请确保程序已正确编译"
        return
    }

    ; 构建命令行
    command := Format('"{1}" {2} "{3}"', deskAIPath, function, text)
    LogDebug("执行命令: " command)

    ; 执行命令并捕获输出
    try {
        g_editTarget.Value := "正在处理文本……`n"
        deskAIDir := g_workingDir "\deskAI"
        resultFile := deskAIDir "\result.txt"  ; 修改结果文件路径

        result := RunWait(command, deskAIDir, "Hide")
        LogDebug("执行结果代码: " result)
        
        if (result = 0) {
            LogDebug("结果文件路径: " resultFile)
            
            if FileExist(resultFile) {
                processed := FileRead(resultFile, "UTF-8")
                g_editTarget.Value := processed
                FileDelete(resultFile)
                LogDebug("处理成功")
            } else {
                LogDebug("错误：结果文件不存在")
                g_editTarget.Value := "处理失败：未找到结果文件"
            }
        } else {
            LogDebug("错误：命令执行失败")
            g_editTarget.Value := "处理失败，请检查配置和网络连接。"
        }
    } catch as processError {
        LogDebug("错误：" processError.Message)
        g_editTarget.Value := "错误：" processError.Message
    }
}