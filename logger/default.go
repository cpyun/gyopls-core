package logger

import (
	"sync"
	"sync/atomic"

	"github.com/cpyun/gyopls-core/contract"
)

var (
	once          sync.Once
	defaultLogger atomic.Pointer[Logger]
)

func init() {
	once.Do(func() {
		// defaultLogger.Store(NewHelper())
	})
}

func SetDefault(l contract.LoggerHandler) {
	defaultLogger.Store(NewHelper(l))
}

func With(args ...any) *Logger {
	return defaultLogger.Load().With(args...)
}

func Debug(msg string, args ...any) {
	defaultLogger.Load().Debug(msg, args...)
}

func Info(msg string, args ...any) {
	defaultLogger.Load().Info(msg, args...)
}

func Warn(msg string, args ...any) {
	defaultLogger.Load().Warn(msg, args...)
}

func Error(msg string, args ...any) {
	defaultLogger.Load().Error(msg, args...)
}

func Fatal(msg string, args ...any) {
	defaultLogger.Load().Fatal(msg, args...)
}
