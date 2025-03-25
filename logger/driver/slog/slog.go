package slog

import (
	"context"
	"log/slog"

	"github.com/cpyun/gyopls-core/contract"
	"github.com/cpyun/gyopls-core/logger/level"
)

type slogApt struct {
	ctx    context.Context
	hanler *slog.Logger
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
	c.hanler = c.hanler.With(args...)

	return c
}

func (t *slogApt) Log(level level.Level, msg string, args ...any) {
	lvl := levlToSlogLeve(level)

	t.hanler.Log(t.ctx, lvl, msg, args...)
}

func (t *slogApt) String() string {
	return "slog"
}

func New() contract.LoggerHandler {
	return &slogApt{
		hanler: slog.Default(),
	}
}

func levlToSlogLeve(lvl level.Level) slog.Level {
	return slog.LevelDebug
}
