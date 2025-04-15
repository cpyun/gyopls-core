package redis

import "context"

type OptionFunc func(opts *redisApt)

type RedisOptions struct {
	ctx    context.Context
	Dsn    string
	Prefix string
}

func setDefaultOptions() RedisOptions {
	return RedisOptions{
		ctx: context.Background(),
	}
}

func WithDSN(dsn string) OptionFunc {
	return func(o *redisApt) {
		o.opts.Dsn = dsn
	}
}

func WithPrefix(prefix string) OptionFunc {
	return func(o *redisApt) {
		o.opts.Prefix = prefix
	}
}
