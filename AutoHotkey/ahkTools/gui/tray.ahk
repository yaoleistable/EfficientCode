#Requires AutoHotkey v2.0

; 初始化系统托盘
InitTray() {
    ; 创建托盘图标
    TraySetIcon("ai.ico")  ; 使用AI图标
    A_TrayMenu.Delete()  ; 清除默认菜单项
    
    ; 添加托盘菜单项
    A_TrayMenu.Add("显示主窗口", ShowMainWindow)
    A_TrayMenu.Add("退出", (*) => ExitApp())
    
    ; 设置双击托盘图标的默认动作
    A_TrayMenu.Default := "显示主窗口"
}