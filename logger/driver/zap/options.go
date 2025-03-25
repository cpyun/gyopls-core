package zap

import (
	"github.com/cpyun/gyopls-core/logger/cut"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapOption struct {
	mode             string
	level            zapcore.Level
	callerSkipKey    int
	configKey        zap.Config
	encoderConfigKey zapcore.EncoderConfig
	namespaceKey     string
	timeFormat       string
	outputPaths      []string
	cutter           *cut.Cut
}

type OptionFunc func(o *zapOption)

func setDefaultOptions() zapOption {
	return zapOption{
		mode:          "release",
		level:         zapcore.InfoLevel,
		callerSkipKey: 1,
		outputPaths:   []string{"stdout"},
		cutter:        cut.Newcut(),
	}
}

func WithModel(name string) OptionFunc {
	return func(o *zapOption) {
		o.mode = name
	}
}

func WithLevel(level zapcore.Level) OptionFunc {
	return func(o *zapOption) {
		o.level = level
	}
}

func WithCallerSkip(i int) OptionFunc {
	return func(o *zapOption) {
		if i < 1 {
			i = 1
		}
		o.callerSkipKey = i
	}
}

// WithConfig pass zap.Config to logger
func WithConfig(c zap.Config) OptionFunc {
	return func(o *zapOption) {
		o.configKey = c
	}
}

// WithEncoderConfig pass zapcore.EncoderConfig to logger
func WithEncoderConfig(c zapcore.EncoderConfig) OptionFunc {
	return func(o *zapOption) {
		o.encoderConfigKey = c
	}
}

func WithNamespace(namespace string) OptionFunc {
	return func(o *zapOption) {
		o.namespaceKey = namespace
	}
}

func WithOutputPath(out ...string) OptionFunc {
	return func(o *zapOption) {
		o.outputPaths = out
	}
}

func WithTimeFormat(timeFormat string) OptionFunc {
	return func(o *zapOption) {
		o.timeFormat = timeFormat
	}
}
