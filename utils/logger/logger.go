package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// Init 初始化日志
func Init() {
	log = logrus.New()
	
	// 设置日志格式
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	
	// 设置日志级别
	level := os.Getenv("LOG_LEVEL")
	switch level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}
	
	// 设置输出
	log.SetOutput(os.Stdout)
}

// Info 记录信息日志
func Info(args ...interface{}) {
	if log != nil {
		log.Info(args...)
	}
}

// Error 记录错误日志
func Error(args ...interface{}) {
	if log != nil {
		log.Error(args...)
	}
}

// Debug 记录调试日志
func Debug(args ...interface{}) {
	if log != nil {
		log.Debug(args...)
	}
}

// Warn 记录警告日志
func Warn(args ...interface{}) {
	if log != nil {
		log.Warn(args...)
	}
}

// Fatal 记录致命错误日志
func Fatal(args ...interface{}) {
	if log != nil {
		log.Fatal(args...)
	}
} 