package cache

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/cpyun/gyopls-core/cache/driver/memory"
	"github.com/cpyun/gyopls-core/contract"
)

var (
	once         sync.Once
	defaultCache atomic.Pointer[CacheAdapter]
)

func init() {
	once.Do(func() {
		defaultCache.Store(NewCache(memory.NewMemory()))
	})
}

func Default() *CacheAdapter {
	return defaultCache.Load()
}

func SetDefault(c contract.CacheHandlerInterface) {
	defaultCache.Store(NewCache(c))
}

func Get(key string) (any, error) {
	return defaultCache.Load().Get(key)
}

func Set(key string, value any, ttl time.Duration) error {
	return defaultCache.Load().Set(key, value, ttl)
}

func Delete(key string) error {
	return defaultCache.Load().Delete(key)
}

func Increase(key string, step int64) error {
	return defaultCache.Load().Increase(key, step)
}

func Decrease(key string, step int64) error {
	return defaultCache.Load().Decrease(key, step)
}

func Expire(key string, dur time.Duration) error {
	return defaultCache.Load().Expire(key, dur)
}
