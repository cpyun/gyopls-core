package file

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

type fileOutput struct {
	handler *lumberjack.Logger
	opts    fileOption
}

func (t *fileOutput) init() {
	t.handler = &lumberjack.Logger{
		Filename:   t.opts.FileName,
		MaxSize:    t.opts.MaxSize,
		MaxAge:     t.opts.MaxAge,
		MaxBackups: t.opts.MaxBackups,
		Compress:   t.opts.Compress,
		LocalTime:  true,
	}
}

func (t *fileOutput) applyOption(opts ...OptionFunc) {
	for _, opt := range opts {
		opt(&t.opts)
	}
}

func (t *fileOutput) Write(p []byte) (int, error) {
	return t.handler.Write(p)
}

func New(opts ...OptionFunc) *fileOutput {
	c := &fileOutput{
		opts: setDefaultOptions(),
	}
	c.applyOption(opts...)

	c.init()
	return c
}
