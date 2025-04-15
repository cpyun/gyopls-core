package zap

import (
	"os"

	"github.com/cpyun/gyopls-core/logger/output"
	"go.uber.org/zap/zapcore"
)

type zapOption struct {
	mode          string
	level         zapcore.Level
	callerSkipKey int
	namespaceKey  string
	timeFormat    string
	output        []zapcore.WriteSyncer // 默认console writer
}

type OptionFunc func(o *zapOption)

func setDefaultOptions() zapOption {
	return zapOption{
		mode:          "production",
		level:         zapcore.DebugLevel,
		callerSkipKey: 3,
		timeFormat:    "2006-01-02T15:04:05.000Z07:00",
		output:        []zapcore.WriteSyncer{zapcore.Lock(os.Stdout)},
	}
}

func WithModel(name string) OptionFunc {
	return func(o *zapOption) {
		o.mode = name
	}
}

func WithCallerSkip(i int) OptionFunc {
	return func(o *zapOption) {
		o.callerSkipKey += i
	}
}

func WithNamespace(namespace string) OptionFunc {
	return func(o *zapOption) {
		o.namespaceKey = namespace
	}
}

func WithTimeFormat(timeFormat string) OptionFunc {
	return func(o *zapOption) {
		o.timeFormat = timeFormat
	}
}

func WithOutput(output ...output.Output) OptionFunc {
	return func(o *zapOption) {
		if len(output) == 0 {
			return
		}

		o.output = o.output[:0]
		for _, out := range output {
			o.output = append(o.output, zapcore.AddSync(out))
		}
	}
}
