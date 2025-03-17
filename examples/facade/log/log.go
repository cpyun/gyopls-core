package log

import (
	"sync"

	"github.com/cpyun/gyopls-core/contract"
	"github.com/cpyun/gyopls-core/logger"
	"github.com/cpyun/gyopls-core/logger/driver/zap"
)

var (
	once          sync.Once
	defaultLogger *logger.Logger
	defaultName   = "zap"
)

func init() {
	once.Do(func() {
		defaultLogger = logger.NewLogger(zap.NewZap(zap.WithOutputPath("stdout", "./logs/logs.log")))
	})
}

func SetDefault(l contract.LoggerHandler) {
	defaultLogger = logger.NewHelper(l)
}

func Debug(msg string, args ...any) {
	defaultLogger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	defaultLogger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	defaultLogger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	defaultLogger.Error(msg, args...)
}

func Fatal(msg string, args ...any) {
	defaultLogger.Fatal(msg, args...)
}
