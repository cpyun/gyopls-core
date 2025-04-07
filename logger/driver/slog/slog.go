package slog

import (
	"context"
	"io"
	"log/slog"

	"github.com/cpyun/gyopls-core/contract"
	"github.com/cpyun/gyopls-core/logger/level"
)

type slogApt struct {
	ctx     context.Context
	handler *slog.Logger
	opts    slogOption
}

func (t *slogApt) init() {
	handler := slog.NewJSONHandler(t.getIoWriter(), t.getHandlerConfig())

	t.handler = slog.New(handler)
}

func (t *slogApt) applyOptions(opts ...optionFunc) {
	for _, opt := range opts {
		opt(&t.opts)
	}
}

func (t *slogApt) getIoWriter() io.Writer {
	return io.MultiWriter(t.opts.output...)
}

func (t *slogApt) getHandlerConfig() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: true,
		Level:     t.opts.level,
	}
}

func (t *slogApt) clone() *slogApt {
	c := *t
	return &c
}

func (t *slogApt) With(args ...any) contract.LoggerHandler {
	if len(args) == 0 {
		return t
	}

	c := t.clone()
	c.handler = c.handler.With(args...)

	return c
}

func (t *slogApt) Log(level level.Level, msg string, args ...any) {
	lvl := levlToSlogLeve(level)

	t.handler.Log(t.ctx, lvl, msg, args...)
}

func (t *slogApt) String() string {
	return "slog"
}

func New(opts ...optionFunc) contract.LoggerHandler {
	sl := &slogApt{
		opts: setDefaultOptions(),
	}
	sl.applyOptions(opts...)

	sl.init()
	return sl
}

func levlToSlogLeve(lvl level.Level) slog.Level {
	switch lvl {
	case level.TraceLevel, level.DebugLevel:
		return slog.LevelDebug
	case level.InfoLevel:
		return slog.LevelInfo
	case level.WarnLevel:
		return slog.LevelWarn
	case level.ErrorLevel:
		return slog.LevelError
	case level.FatalLevel:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
