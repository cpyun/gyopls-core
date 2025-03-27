package logger

import (
	"os"

	"github.com/cpyun/gyopls-core/contract"
	"github.com/cpyun/gyopls-core/logger/level"
)

type Logger struct {
	handler contract.LoggerHandler // handler
	fields  []any                  // 运行时的公共字段（非初始化）
	filter  filterOptions          // filter
	// filter  []func(lvl level.Level, keyVals ...any) bool // filter
	opts LoggerOptions
}

func (t *Logger) init() {}

func (t *Logger) applyOption(opts ...OptionFunc) {
	for _, o := range opts {
		o(t)
	}
}

func (t *Logger) clone() *Logger {
	c := *t
	return &c
}

func (t *Logger) With(fields ...any) *Logger {
	if len(fields) == 0 {
		return t
	}

	c := t.clone()
	c.fields = append(c.fields, fields...)
	return c
}

func (t *Logger) Log(level level.Level, msg string, args ...any) {
	t.log(level, msg, args...)
}

// trace
func (t *Logger) Trace(msg string, args ...any) {
	t.log(level.TraceLevel, msg, args...)
}

// debug
func (t *Logger) Debug(msg string, args ...any) {
	t.log(level.DebugLevel, msg, args...)
}

// info
func (t *Logger) Info(msg string, args ...any) {
	t.log(level.InfoLevel, msg, args...)
}

// warn
func (t *Logger) Warn(msg string, args ...any) {
	t.log(level.WarnLevel, msg, args...)
}

func (t *Logger) Error(msg string, args ...any) {
	t.log(level.ErrorLevel, msg, args...)
}

func (t *Logger) Fatal(msg string, args ...any) {
	t.log(level.FatalLevel, msg, args...)
	os.Exit(1)
}

func (t *Logger) log(level level.Level, msg string, args ...any) {
	field := t.fields
	field = append(field, args...)

	// 字段过滤
	data := t.checkFilter(level, field...)

	t.handler.Log(level, msg, data...)
}

func NewLogger(log contract.LoggerHandler, opts ...OptionFunc) *Logger {
	if log == nil {
		panic("nil loggerHandler")
	}

	l := &Logger{
		handler: log,
		opts:    setDefaultOptions(),
	}
	l.applyOption(opts...)

	l.init()
	return l
}
