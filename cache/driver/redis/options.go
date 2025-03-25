package redis

type RedisOptions struct {
	// config.Redis
	Dsn    string
	Prefix string
}

type OptionFunc func(opts *redisApt)

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
