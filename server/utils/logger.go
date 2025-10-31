package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	// Logger 全局日志实例
	Logger *log.Logger
	// ErrorLogger 错误日志实例
	ErrorLogger *log.Logger
	// InfoLogger 信息日志实例
	InfoLogger *log.Logger
)

// InitLogger 初始化日志系统
func InitLogger(logDir string) error {
	// 创建日志目录
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	// 获取当前日期作为日志文件名
	dateStr := time.Now().Format("2006-01-02")

	// 通用日志文件
	logFile, err := os.OpenFile(
		filepath.Join(logDir, "app-"+dateStr+".log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return err
	}

	// 错误日志文件
	errorLogFile, err := os.OpenFile(
		filepath.Join(logDir, "error-"+dateStr+".log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return err
	}

	// 创建多写入器，同时写入文件和控制台
	logWriter := io.MultiWriter(os.Stdout, logFile)
	errorWriter := io.MultiWriter(os.Stderr, errorLogFile)

	// 创建日志实例
	Logger = log.New(logWriter, "[LOG] ", log.LstdFlags|log.Lshortfile)
	ErrorLogger = log.New(errorWriter, "[ERROR] ", log.LstdFlags|log.Lshortfile)
	InfoLogger = log.New(logWriter, "[INFO] ", log.LstdFlags|log.Lshortfile)

	return nil
}

// LogInfo 记录信息日志
func LogInfo(format string, v ...interface{}) {
	if InfoLogger != nil {
		InfoLogger.Printf(format, v...)
	} else {
		log.Printf("[INFO] "+format, v...)
	}
}

// LogError 记录错误日志
func LogError(format string, v ...interface{}) {
	if ErrorLogger != nil {
		ErrorLogger.Printf(format, v...)
	} else {
		log.Printf("[ERROR] "+format, v...)
	}
}

// LogWarning 记录警告日志
func LogWarning(format string, v ...interface{}) {
	if Logger != nil {
		Logger.Printf("[WARN] "+format, v...)
	} else {
		log.Printf("[WARN] "+format, v...)
	}
}

// LogDebug 记录调试日志（仅在开发环境）
func LogDebug(format string, v ...interface{}) {
	if Logger != nil {
		Logger.Printf("[DEBUG] "+format, v...)
	}
}
