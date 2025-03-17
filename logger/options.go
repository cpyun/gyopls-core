package logger

import (
	"context"
	"sync"

	"github.com/cpyun/gyopls-core/logger/level"
)

type LoggerOptions struct {
	drivers sync.Map        //
	Context context.Context //
	level   level.Level     //
	Name    string          // logger name
}

type OptionFunc func(*Logger)

// WithLevel set default level for the logger
func WithLevel(lvl level.Level) OptionFunc {
	return func(o *Logger) {
		o.opts.level = lvl
	}
}

// WithName set name for logger
func WithName(name string) OptionFunc {
	return func(o *Logger) {
		o.opts.Name = name
	}
}

func WithContext(k, v interface{}) OptionFunc {
	return func(o *Logger) {
		if o.opts.Context == nil {
			o.opts.Context = context.Background()
		}
		o.opts.Context = context.WithValue(o.opts.Context, k, v)
	}
}
