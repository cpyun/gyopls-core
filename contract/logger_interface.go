package contract

import (
	"github.com/cpyun/gyopls-core/logger/level"
)

// Logger is a generic logging interface
type LoggerHandler interface {
	// String returns the name of logger
	String() string
	// With set fields to always be logged
	With(fields ...any) LoggerHandler
	Log(level level.Level, msg string, v ...any)
	// SetCut(cut *cut.Cut) LoggerHandler
}
