package cache

type cacheAdapterOptions struct {
	prefix string // 缓存前缀
	// drivers sync.Map // 缓存驱动
}

type OptionFunc func(*CacheAdapter)

func setDefaultOptions() cacheAdapterOptions {
	// opts.drivers["redis"] = driver.NewRedis()
	// t.drivers["memory"] = driver.NewMemory()
	// t.drivers["memcache"] = driver.NewRedis()
	return cacheAdapterOptions{}
}

func WithPrefix(prefix string) OptionFunc {
	return func(t *CacheAdapter) {
		t.opts.prefix = prefix
	}
}
