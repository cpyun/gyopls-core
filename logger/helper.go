package logger

import "github.com/cpyun/gyopls-core/contract"

func NewHelper(log contract.LoggerHandler) *Logger {
	return &Logger{
		handler: log,
	}
}
