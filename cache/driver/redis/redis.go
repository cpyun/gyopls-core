package redis

import (
	"context"
	"time"

	"github.com/cpyun/gyopls-core/contract"
	"github.com/redis/go-redis/v9"
)

// redis cache implement
type redisApt struct {
	ctx     context.Context
	handler *redis.Client
	opts    RedisOptions
}

func (r *redisApt) init() {
	client := redis.NewClient(r.parseRedisConfig())
	if err := client.Ping(r.ctx).Err(); err != nil {
		panic(err.Error())
	}

	r.handler = client
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
	return r.handler.Get(r.ctx, key).Result()
}

// Set value with key and expire time
func (r *redisApt) Set(key string, val any, expire time.Duration) error {
	key = r.getCacheKey(key)
	return r.handler.Set(r.ctx, key, val, expire).Err()
}

// Del delete key in redis
func (r *redisApt) Delete(key string) error {
	key = r.getCacheKey(key)
	return r.handler.Del(r.ctx, key).Err()
}

// HashGet from key
func (r *redisApt) HashGet(key, filed string) (any, error) {
	key = r.getCacheKey(key)
	return r.handler.HGet(r.ctx, key, filed).Result()
}

func (r *redisApt) HashSet(key string, values ...any) (any, error) {
	key = r.getCacheKey(key)
	return r.handler.HSet(r.ctx, key, values...).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *redisApt) HashDelete(key string, fileds ...string) error {
	key = r.getCacheKey(key)
	return r.handler.HDel(r.ctx, key, fileds...).Err()
}

// Increase value
func (r *redisApt) Increase(key string, step int64) error {
	key = r.getCacheKey(key)
	return r.handler.IncrBy(r.ctx, key, step).Err()
}

func (r *redisApt) Decrease(key string, step int64) error {
	key = r.getCacheKey(key)
	return r.handler.DecrBy(r.ctx, key, step).Err()
}

// Expire Set ttl
func (r *redisApt) Expire(key string, dur time.Duration) error {
	key = r.getCacheKey(key)
	return r.handler.Expire(r.ctx, key, dur).Err()
}

func (r *redisApt) Handler() any {
	return r.handler
}

// NewRedis redis模式
func NewRedis(opts ...OptionFunc) contract.CacheHandlerInterface {
	r := &redisApt{
		ctx: context.Background(),
	}
	r.withOptions(opts...)

	r.init()
	return r
}
