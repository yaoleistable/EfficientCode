#Requires AutoHotkey v2.0

; 注册Sublime Text快捷键
#n::
{
    SublimeTextHotkey()
}

; 定义处理 Sublime Text 热键的函数
SublimeTextHotkey() {

        ; 尝试查找 Sublime Text 窗口
        sublimeWindow := WinExist("ahk_class PX_WINDOW_CLASS")
        
        if sublimeWindow  ; 如果窗口存在
        {
            if WinGetMinMax("ahk_id " sublimeWindow) = -1
                WinRestore("ahk_id " sublimeWindow)
            WinActivate("ahk_id " sublimeWindow)
        }
        else  ; 如果窗口不存在，运行新实例
        {
            ; 检查两个可能的开始菜单路径
            startMenuPaths := [
                EnvGet("APPDATA") "\Microsoft\Windows\Start Menu\Programs\",
                EnvGet("ProgramData") "\Microsoft\Windows\Start Menu\Programs\"
            ]

            shortcutFound := false
            for basePath in startMenuPaths {
                ; 使用 Loop Files 来查找匹配的快捷方式
                Loop Files, basePath "sublime*.lnk", "F"  ; F 表示仅文件
                {
                    Run A_LoopFilePath
                    shortcutFound := true
                    break
                }
                if shortcutFound
                    break
            }

            if !shortcutFound {
                ; 如果快捷方式不存在，尝试直接运行程序
                sublimePaths := [
                    A_ProgramFiles "\Sublime Text\sublime_text.exe",
                    A_ProgramFiles "\Sublime Text 3\sublime_text.exe",
                    EnvGet("LOCALAPPDATA") "\Programs\Sublime Text\sublime_text.exe"
                ]
                
                for path in sublimePaths {
                    if FileExist(path) {
                        Run path
                        return
                    }
                }
                
                MsgBox("未找到 Sublime Text，请确保已正确安装。", "错误", "48")
            }
        }
    }

