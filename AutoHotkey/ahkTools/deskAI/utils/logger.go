package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var logger *log.Logger

func init() {
	// 使用上级目录的 logs 文件夹
	execPath, _ := os.Executable()
	execDir := filepath.Dir(execPath)
	parentDir := filepath.Dir(execDir)
	logDir := filepath.Join(parentDir, "logs")

	// 使用同一个日志文件
	logFile := filepath.Join(logDir, "debug.log")

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("创建日志文件失败:", err)
	}

	// 使用与 AutoHotkey 相同的时间格式
	logger = log.New(f, "", 0)
}

// LogInfo 记录信息日志
func LogInfo(format string, v ...interface{}) {
	if logger != nil {
		timeStr := time.Now().Format("15:04 2006-01-02")
		logger.Printf("%s: %s", timeStr, fmt.Sprintf(format, v...))
	}
}

// LogError 记录错误日志
func LogError(format string, v ...interface{}) {
	if logger != nil {
		timeStr := time.Now().Format("15:04 2006-01-02")
		logger.Printf("%s: ERROR: %s", timeStr, fmt.Sprintf(format, v...))
	}
}
