package cache

import (
	"testing"
	"time"

	"github.com/cpyun/gyopls-core/cache/driver/memory"
)

const (
	key   = "key"
	value = "value"
	ttl   = 10 * time.Second
)

var (
	handler = memory.NewMemory()
)

func TestCache_SetAndGet(t *testing.T) {
	cache := NewCache(handler)
	if err := cache.Set(key, value, ttl); err != nil {
		t.Error(err.Error())
	}

	val, err := cache.Get(key)
	if err != nil {
		t.Error(err.Error())
	}

	if val != value {
		t.Errorf("Expected value1, got %v", val)
	}
}

func TestCache_Delete(t *testing.T) {
	cache := NewCache(handler)
	cache.Set(key, value, ttl)

	cache.Delete(key)
	if val, err := cache.Get(key); err == nil || val != nil {
		t.Error(err.Error())
	}
}

func TestCache_Expire(t *testing.T) {
	cache := NewCache(handler)
	cache.Set(key, value, ttl)

	cache.Expire(key, 1*time.Millisecond)
	time.Sleep(2 * time.Millisecond)
	if val, err := cache.Get(key); err != nil || val != nil {
		t.Errorf("Expected nothing, got %v", val)
	}
}

func TestCache_Increase(t *testing.T) {
	cache := NewCache(handler)
	cache.Set(key, 1, ttl)

	cache.Increase(key, 1)
	if val, err := cache.Get(key); err != nil || val.(int64) != 2 {
		t.Errorf("Expected %d, got %v", val, err)
	}
}

func TestCache_Decrease(t *testing.T) {
	cache := NewCache(handler)
	cache.Set(key, 1, ttl)

	cache.Decrease(key, -1)
	if val, err := cache.Get(key); err != nil || val.(int64) != 0 {
		t.Errorf("Expected %d, got %v", val, err)
	}
}
