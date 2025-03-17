package cache

import (
	"sync"
	"time"

	"github.com/cpyun/gyopls-core/cache"
	"github.com/cpyun/gyopls-core/cache/driver/memory"
	"github.com/cpyun/gyopls-core/contract"
	"github.com/cpyun/gyopls-core/examples/internal/config"
)

var (
	once         sync.Once
	defaultCache *cache.CacheAdapter
	defaultName  string
)

func init() {
	once.Do(func() {
		defaultName = config.CacheConfig.Driver
		defaultCache = cache.NewCache(memory.NewMemory(),
			cache.WithPrefix("test:"),
		)
	})
}

func Get(key string) (any, error) {
	return defaultCache.Store(defaultName).Get(key)
}

func Set(key string, value any, ttl time.Duration) error {
	return defaultCache.Store(defaultName).Set(key, value, ttl)
}

func Delete(key string) error {
	return defaultCache.Store(defaultName).Delete(key)
}

// Store 切换缓存实例
func Store(name string) contract.CacheHandlerInterface {
	return defaultCache.Store(name)
}
