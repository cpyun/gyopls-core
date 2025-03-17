package logger

import (
	"github.com/cpyun/gyopls-core/contract"
	"github.com/cpyun/gyopls-core/logger/level"
)

type Logger struct {
	handler contract.LoggerHandler                       // handler
	filter  []func(lvl level.Level, keyVals ...any) bool // filter
	fields  []any                                        // 运行时的公共字段（非初始化）
	opts    LoggerOptions
}

func (t *Logger) init() {}

func (t *Logger) clone() *Logger {
	c := *t
	return &c
}

// checkFilter 校验字段过滤
func (t *Logger) checkFilter(lvl level.Level, args ...any) []any {
	for _, f := range t.filter {
		if f != nil && f(lvl, args...) {
			break
			// return nil
		}
	}
	return args
}

func (t *Logger) Log(level level.Level, msg string, args ...any) {
	field := args

	// 公共field
	for _, v := range t.fields {
		field = append(field, v)
	}

	// 字段过滤
	data := t.checkFilter(level, field...)

	t.handler.Log(level, msg, data...)
}

func (t *Logger) With(fields ...any) *Logger {
	if len(fields) == 0 {
		return t
	}

	c := t.clone()
	c.fields = append(c.fields, fields...)
	return c
}

// trace
// func (t *Logger) Trace(msg string, args ...any) {
// 	t.Log(level.TraceLevel, msg, args...)
// }

// debug
func (t *Logger) Debug(msg string, args ...any) {
	t.Log(level.DebugLevel, msg, args...)
}

// info
func (t *Logger) Info(msg string, args ...any) {
	t.Log(level.InfoLevel, msg, args...)
}

// warn
func (t *Logger) Warn(msg string, args ...any) {
	t.Log(level.WarnLevel, msg, args...)
}

func (t *Logger) Error(msg string, args ...any) {
	t.Log(level.ErrorLevel, msg, args...)
}

func (t *Logger) Fatal(msg string, args ...any) {
	t.Log(level.FatalLevel, msg, args...)
}

func (t *Logger) Store(name string) contract.LoggerHandler {
	if v, ok := t.opts.drivers.Load(name); ok {
		if h, ok := v.(contract.LoggerHandler); ok {
			return h
		}
		return nil
	}

	return nil
}

func NewLogger(log contract.LoggerHandler, opts ...OptionFunc) *Logger {
	l := &Logger{
		handler: log,
	}

	for _, o := range opts {
		o(l)
	}

	l.init()
	return l
}
