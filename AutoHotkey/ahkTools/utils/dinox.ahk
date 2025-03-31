#Requires AutoHotkey v2.0
; 调用deskAI发送内容到DinoX
; 作者：Lei
; 日期：2025-3-30
ProcessDinox(*) {
    ; 获取输入文本
    g_editTarget.Value := ""
    sourceText := g_editSource.Value
    if (sourceText = "") {
        MsgBox("请输入要发送的内容！", "提示", "48")
        return
    }

    ; 使用全局工作目录构建正确的路径
    deskAIPath := g_workingDir "\deskAI\deskAI.exe"
    
    ; 检查文件是否存在
    if !FileExist(deskAIPath) {
        g_editTarget.Value := "错误：找不到 deskAI.exe，请确保程序已正确编译"
        return
    }

    ; 构建命令行
    command := Format('"{1}" dinoxPost "{2}"', deskAIPath, sourceText)

    ; 执行命令并捕获输出
    try {
        g_editTarget.Value := "正在发送到Dinox……`n"
        ;g_editTarget.Value := "正在执行的命令: " command "`n"
        result := RunWait(command, , "Hide")
        if (result = 0) {
            g_editTarget.Value .= "内容已成功发送到 Dinox！"
        } else {
            g_editTarget.Value := "发送失败，请检查配置和网络连接。"
        }
    } catch as err {
        g_editTarget.Value := "错误：" err.Message
    }
}