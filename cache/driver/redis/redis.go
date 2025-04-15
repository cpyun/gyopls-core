package redis

import (
	"time"

	"github.com/cpyun/gyopls-core/contract"
	"github.com/redis/go-redis/v9"
)

// redis cache implement
type redisApt struct {
	handler *redis.Client
	opts    RedisOptions
}

func (r *redisApt) init() {
	rdb := redis.NewClient(r.parseRedisConfig())
	if err := rdb.Ping(r.opts.ctx).Err(); err != nil {
		panic(err)
	}

	r.handler = rdb
}

func (r *redisApt) withOptions(opts ...OptionFunc) {
	for _, opt := range opts {
		opt(r)
	}
}

// 解析redis配置
func (r *redisApt) parseRedisConfig() *redis.Options {
	val, _ := redis.ParseURL(r.opts.Dsn)
	return val
}

func (r *redisApt) Close() error {
	return r.handler.Close()
}

func (*redisApt) String() string {
	return "redis"
}

func (r *redisApt) getCacheKey(key string) string {
	return r.opts.Prefix + key
}

// Get from key
func (r *redisApt) Get(key string) (any, error) {
	key = r.getCacheKey(key)
	return r.handler.Get(r.opts.ctx, key).Result()
}

// Set value with key and expire time
func (r *redisApt) Set(key string, val any, expire time.Duration) error {
	key = r.getCacheKey(key)
	return r.handler.Set(r.opts.ctx, key, val, expire).Err()
}

// Del delete key in redis
func (r *redisApt) Delete(key string) error {
	key = r.getCacheKey(key)
	return r.handler.Del(r.opts.ctx, key).Err()
}

// HashGet from key
func (r *redisApt) HashGet(key, filed string) (any, error) {
	key = r.getCacheKey(key)
	return r.handler.HGet(r.opts.ctx, key, filed).Result()
}

func (r *redisApt) HashSet(key string, values ...any) (any, error) {
	key = r.getCacheKey(key)
	return r.handler.HSet(r.opts.ctx, key, values...).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *redisApt) HashDelete(key string, fileds ...string) error {
	key = r.getCacheKey(key)
	return r.handler.HDel(r.opts.ctx, key, fileds...).Err()
}

// Increase value
func (r *redisApt) Increase(key string, step int64) error {
	key = r.getCacheKey(key)
	return r.handler.IncrBy(r.opts.ctx, key, step).Err()
}

func (r *redisApt) Decrease(key string, step int64) error {
	key = r.getCacheKey(key)
	return r.handler.DecrBy(r.opts.ctx, key, step).Err()
}

// Expire Set ttl
func (r *redisApt) Expire(key string, dur time.Duration) error {
	key = r.getCacheKey(key)
	return r.handler.Expire(r.opts.ctx, key, dur).Err()
}

func (r *redisApt) Handler() any {
	return r.handler
}

// NewRedis redis模式
func NewRedis(opts ...OptionFunc) contract.CacheHandlerInterface {
	r := &redisApt{
		opts: setDefaultOptions(),
	}
	r.withOptions(opts...)

	r.init()
	return r
}
