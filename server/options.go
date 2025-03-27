package server

import (
	"time"

	"github.com/cpyun/gyopls-core/logger"
)

type OptionFunc func(*options)

type options struct {
	gracefulShutdownTimeout time.Duration
	logger                  *logger.Logger
}

func setDefaultOptions() options {
	return options{
		gracefulShutdownTimeout: 5 * time.Second,
	}
}

func WithGracefulShutdownTimeout(timeout int64) OptionFunc {
	return func(o *options) {
		o.gracefulShutdownTimeout = time.Duration(timeout) * time.Second
	}
}
