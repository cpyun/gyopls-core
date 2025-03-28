package cache

type cacheAdapterOptions struct {
	prefix string // 缓存前缀
}

type OptionFunc func(*CacheAdapter)

func setDefaultOptions() cacheAdapterOptions {
	return cacheAdapterOptions{}
}

func WithPrefix(prefix string) OptionFunc {
	return func(t *CacheAdapter) {
		t.opts.prefix = prefix
	}
}
