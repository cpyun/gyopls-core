package contract

import "time"

type CacheHandlerInterface interface {
	Get(key string) (any, error)
	Set(key string, val any, expire time.Duration) error
	Delete(key string) error
	Increase(key string, step int) error
	Decrease(key string, step int) error
	Expire(key string, dur time.Duration) error
	Handler() any
}
