package redis

import (
	"context"
	"time"

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

func (r *redisApt) close() error {
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
	return r.handler.Set(r.ctx, key, val, time.Duration(expire)*time.Second).Err()
}

// Del delete key in redis
func (r *redisApt) Delete(key string) error {
	key = r.getCacheKey(key)
	return r.handler.Del(r.ctx, key).Err()
}

// HashGet from key
func (r *redisApt) HashGet(hk, key string) (any, error) {
	key = r.getCacheKey(key)
	return r.handler.HGet(r.ctx, hk, key).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *redisApt) HashDelete(hk, key string) error {
	key = r.getCacheKey(key)
	return r.handler.HDel(r.ctx, hk, key).Err()
}

// Increase value
func (r *redisApt) Increase(key string, step int) error {
	key = r.getCacheKey(key)
	return r.handler.Incr(r.ctx, key).Err()
}

func (r *redisApt) Decrease(key string, step int) error {
	key = r.getCacheKey(key)
	return r.handler.Decr(r.ctx, key).Err()
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
func NewRedis(opts ...OptionFunc) *redisApt {
	r := &redisApt{
		ctx: context.Background(),
	}

	r.withOptions(opts...)
	r.init()

	return r
}
