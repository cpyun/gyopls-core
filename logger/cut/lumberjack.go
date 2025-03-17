package cut

import (
	"io"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Cut struct {
	handler *lumberjack.Logger
	opts    CutOption
}

func (t *Cut) init() {
	t.handler = &lumberjack.Logger{
		Filename:   t.opts.FileName,
		MaxSize:    t.opts.MaxSize,
		MaxAge:     t.opts.MaxAge,
		MaxBackups: t.opts.MaxBackups,
		Compress:   t.opts.Compress,
		LocalTime:  true,
	}
}

func (t *Cut) withOption(opts ...OptionFunc) {
	for _, opt := range opts {
		opt(t)
	}
}

func (t *Cut) SetFile(name string) io.Writer {
	if name == "stdout" || name == "stderr" {
		return os.Stdout
	}
	t.handler.Filename = name
	return t
}

func (t *Cut) Write(p []byte) (int, error) {
	return t.handler.Write(p)
}

func Newcut(opts ...OptionFunc) *Cut {
	c := &Cut{
		opts: setDefault(),
	}
	c.withOption(opts...)

	return c
}
