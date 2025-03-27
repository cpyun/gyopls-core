package logger

import (
	"github.com/cpyun/gyopls-core/contract"
	"github.com/cpyun/gyopls-core/logger/level"
)

func NewHelper(log contract.LoggerHandler) *Logger {
	return NewLogger(log, WithLevel(level.InfoLevel))
}
