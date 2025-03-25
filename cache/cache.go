package cache

import (
	"time"

	"github.com/cpyun/gyopls-core/contract"
)

type CacheAdapter struct {
	handler contract.CacheHandlerInterface
	opts    cacheAdapterOptions
}

func (t *CacheAdapter) withOptions(opts ...OptionFunc) {
	for _, opt := range opts {
		opt(t)
	}
}

func (r *CacheAdapter) getCacheKey(key string) string {
	return r.opts.prefix + key
}

func (t *CacheAdapter) Get(name string) (any, error) {
	return t.handler.Get(t.getCacheKey(name))
}

func (t *CacheAdapter) Set(name string, value any, dur time.Duration) error {
	return t.handler.Set(t.getCacheKey(name), value, dur)
}

func (t *CacheAdapter) Delete(name string) error {
	return t.handler.Delete(t.getCacheKey(name))
}

// 连接或者切换缓存
func (t *CacheAdapter) Store(name string) contract.CacheHandlerInterface {
	return t.handler
}

func NewCache(handler contract.CacheHandlerInterface, opts ...OptionFunc) *CacheAdapter {
	t := &CacheAdapter{
		handler: handler,
		opts:    setDefaultOptions(),
	}

	t.withOptions(opts...)
	return t
}
