package cache

import (
	"sync"
	"time"

	"github.com/cpyun/gyopls-core/contract"
)

type CacheAdapter struct {
	handler contract.CacheHandlerInterface
	lock    sync.RWMutex
	opts    cacheAdapterOptions
}

func (t *CacheAdapter) applyOptions(opts ...OptionFunc) {
	for _, opt := range opts {
		opt(t)
	}
}

func (r *CacheAdapter) getCacheKey(key string) string {
	return r.opts.prefix + key
}

func (t *CacheAdapter) Get(name string) (any, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return t.handler.Get(t.getCacheKey(name))
}

func (t *CacheAdapter) Set(name string, value any, dur time.Duration) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	return t.handler.Set(t.getCacheKey(name), value, dur)
}

func (t *CacheAdapter) Delete(name string) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	return t.handler.Delete(t.getCacheKey(name))
}

func (t *CacheAdapter) Increase(key string, step int64) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	return t.handler.Increase(t.getCacheKey(key), step)
}

func (t *CacheAdapter) Decrease(key string, step int64) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	return t.handler.Decrease(t.getCacheKey(key), step)
}

func (t *CacheAdapter) Expire(key string, dur time.Duration) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	return t.handler.Expire(t.getCacheKey(key), dur)
}

func NewCache(handler contract.CacheHandlerInterface, opts ...OptionFunc) *CacheAdapter {
	t := &CacheAdapter{
		handler: handler,
		opts:    setDefaultOptions(),
	}

	t.applyOptions(opts...)
	return t
}
