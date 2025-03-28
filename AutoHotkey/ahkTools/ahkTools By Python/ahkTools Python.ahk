#Requires AutoHotkey v2.0

; 作者：Lei
; 日期：2025-03-28
; 版本：1.0
; 功能：AI文本工具一个基于AutoHotkey v2.0和AI服务的文本处理工具，支持文本翻译、润色等功能。
; 1. 安装AutoHotkey v2.0。
; 2. 将脚本保存为 .ahk 文件。
; 3. 配置脚本中的AI服务地址和API密钥。
; 4. 运行脚本。
; 5. 按下热键（默认是 !t）来翻译文本。
; 6. 按下热键（默认是 !p）来润色文本。


#SingleInstance Force
#Warn

; 读取配置
config := IniRead("config.ini", "Hotkey", "translate", "!t")
polishHotkey := IniRead("config.ini", "Hotkey", "polish", "!p")
width := IniRead("config.ini", "GUI", "width", "400")
height := IniRead("config.ini", "GUI", "height", "300")

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

; 确保热键格式正确
config := FormatHotkey(config)
polishHotkey := FormatHotkey(polishHotkey)

; 创建GUI
MyGui := Gui()
MyGui.SetFont("s10", "Microsoft YaHei")
MyGui.Title := "AI文本工具"

; 创建托盘图标
TraySetIcon("*")  ; 使用默认图标
A_TrayMenu.Delete()  ; 清除默认菜单项
A_TrayMenu.Add("显示主窗口", ShowMainWindow)
A_TrayMenu.Add("退出", (*) => ExitApp())
A_TrayMenu.Default := "显示主窗口"  ; 设置双击托盘图标的默认动作

; 添加控件
MyGui.Add("Text", , "输入文本:")
editSource := MyGui.Add("Edit", "vSource w" width " h60")
MyGui.Add("Text", , "处理结果:")
editTarget := MyGui.Add("Edit", "vTarget w" width " h60")

; 添加功能按钮
btnTranslate := MyGui.Add("Button", "Default", "翻译")
btnPolish := MyGui.Add("Button", , "润色")
btnCopySource := MyGui.Add("Button", , "复制输入")
btnCopyTarget := MyGui.Add("Button", , "复制结果")

; 注册事件
btnTranslate.OnEvent("Click", (*) => ProcessText("translate"))
btnPolish.OnEvent("Click", (*) => ProcessText("polish"))
btnCopySource.OnEvent("Click", CopySource)
btnCopyTarget.OnEvent("Click", CopyTarget)

; 初始化时最小化到托盘
MyGui.Show("AutoSize Hide")

; 设置窗口关闭事件
MyGui.OnEvent("Close", (*) => MyGui.Hide())

; 显示主窗口函数
ShowMainWindow(*) {
    MyGui.Show()
    WinActivate("AI文本工具")
}

; 注册热键
if (config) {
    HotIfWinNotActive "AI文本工具"
    try {
        if (config)
            Hotkey config, (*) => ShowTool("translate")
        if (polishHotkey)
            Hotkey polishHotkey, (*) => ShowTool("polish")
    } catch Error as err {
        MsgBox Format("热键注册失败: {}`n配置的热键: {}, {}",
            err.Message, config, polishHotkey)
    }
} else {
    MsgBox "热键格式无效，请检查配置文件。`n应使用如 !t (Alt+T) 或 ^p (Ctrl+P) 的格式。"
}

; 处理文本
ProcessText(function) {
    text := editSource.Value
    if (text = "")
        return

    result := RunWait(Format('python ai_tool.py "{}" "{}"', function, text), , "Hide")
    if FileExist("result.txt") {
        try {
            processed := FileRead("result.txt", "UTF-8")
            editTarget.Value := processed
        } catch Error as processError {
            editTarget.Value := "处理文本时出错: " processError.Message
        }
        FileDelete("result.txt")
    }
}

; 复制函数
CopySource(*) {
    A_Clipboard := editSource.Value
}

CopyTarget(*) {
    A_Clipboard := editTarget.Value
}

; 显示工具
ShowTool(function, *) {
    ; 保存当前活动窗口句柄
    sourceWin := WinExist("A")
    
    ; 检查窗口状态并确保正确显示
    if !WinExist("AI文本工具") {
        MyGui.Show()
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
            editSource.Value := text
            ProcessText(function)
        } else {
            errorMsg := "无法获取选中的文本。请确保："
            errorMsg .= "`n1. 已选中要处理的文本"
            errorMsg .= "`n2. 当前窗口允许复制操作"
            errorMsg .= "`n3. 没有其他程序占用剪贴板"
            errorMsg .= "`n4. 选中的文本不为空"
            MsgBox(errorMsg, "错误", "48")
        }
    } catch Error as processError {
        MsgBox("处理文本时出错: " processError.Message, "错误", "48")
    } finally {
        A_Clipboard := savedClip  ; 恢复剪贴板内容
        BlockInput("Off")  ; 确保输入未被阻止
    }
}

; 退出时清理
OnExit ExitFunc

ExitFunc(ExitReason, ExitCode) {
    if FileExist("result.txt")
        FileDelete("result.txt")
}



; Win+Z 打开 AutoHotkey 官网
; #z::Run "https://www.autohotkey.com"

; 功能一：Win+N 处理 Sublime Text
#n::
{
    ; 尝试查找 Sublime Text 窗口
    ; ahk_class PX_WINDOW_CLASS 是 Sublime Text 的窗口类名
    sublimeWindow := WinExist("ahk_class PX_WINDOW_CLASS")
    
    if sublimeWindow  ; 如果窗口存在
    {
        ; 如果窗口最小化，恢复窗口
        if WinGetMinMax("ahk_id " sublimeWindow) = -1
            WinRestore("ahk_id " sublimeWindow)
        
        ; 激活窗口
        WinActivate("ahk_id " sublimeWindow)
    }
    else  ; 如果窗口不存在，运行新实例
    {
        Run "C:\Users\Lei\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\sublime_text.lnk"
    }
}

; 功能二：字符串插入
; 1. 复制一段文字A之后，进行以下操作：
;    - 在英文句号(.)或中文句号(。)后面添加字符串C。
;    - 在其他标点符号后面添加字符串B。
; 2. 按下cr字母后，粘贴修改后的文本

B := "((⏱️=2000))"  ; 要添加在其他标点符号后的字符串B，注意字符串两边的空格
C := "((⏱️=4000))"  ; 要添加在句号后的字符串C，注意字符串两边的空格

; 热字串触发
::cr::
{
    ; 保存当前剪贴板内容
    ClipSaved := ClipboardAll()
    
    ; 获取剪贴板文本内容
    TextContent := A_Clipboard  ; 直接使用 A_Clipboard 而不是 ClipboardAll

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

; 功能三：自动化输入
; 自动填充email
::email::815141681@qq.com
::8151::815141681@qq.com
::qq::815141681@qq.com
;任务完成
::wc::✅
;待办任务
::db::⏳
; 输入 rq 自动生成当天日期（格式：YYYYMMDD）
::rq::
{
    ; 获取当前日期并格式化
    currentDate := FormatTime(, "yyyyMMdd")
    
    ; 发送格式化后的日期
    SendText(currentDate)
}

; 退出脚本的热键 (可选)
^!q::ExitApp()
