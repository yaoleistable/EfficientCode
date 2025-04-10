package ai

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "time"
)

var logger *log.Logger

func init() {
    // 创建日志目录
    logDir := "logs"
    if err := os.MkdirAll(logDir, 0755); err != nil {
        log.Fatal("创建日志目录失败:", err)
    }

    // 创建日志文件
    logFile := filepath.Join(logDir, fmt.Sprintf("deskAI_%s.log", time.Now().Format("20060102")))
    f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal("创建日志文件失败:", err)
    }

    // 初始化日志记录器
    logger = log.New(f, "", log.LstdFlags)
}

// LogInfo 记录信息日志
func LogInfo(format string, v ...interface{}) {
    if logger != nil {
        logger.Printf("[INFO] "+format, v...)
    }
}

// LogError 记录错误日志
func LogError(format string, v ...interface{}) {
    if logger != nil {
        logger.Printf("[ERROR] "+format, v...)
    }
}