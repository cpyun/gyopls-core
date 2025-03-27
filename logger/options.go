package logger

import (
	"context"

	"github.com/cpyun/gyopls-core/logger/level"
)

type OptionFunc func(*Logger)

type LoggerOptions struct {
	context context.Context //
	level   level.Level     //
}

func setDefaultOptions() LoggerOptions {
	return LoggerOptions{
		context: context.Background(),
		level:   level.InfoLevel,
	}
}

// WithLevel set default level for the logger
func WithLevel(lvl level.Level) OptionFunc {
	return func(o *Logger) {
		o.opts.level = lvl
	}
}
