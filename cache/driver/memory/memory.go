package memory

import (
	"fmt"
	"sync"
	"time"

	"github.com/cpyun/gyopls-core/contract"
	"github.com/spf13/cast"
)

type item struct {
	Value   any
	Expired time.Time
}

type memory struct {
	items map[string]*item
	mutex sync.RWMutex
	opt   memoryOptions
}

func (*memory) String() string {
	return "memory"
}

func (r *memory) getCacheKey(key string) string {
	return r.opt.prefix + key
}

func (m *memory) setItem(key string, item *item) error {
	if m.opt.maxEntries <= 0 || len(m.items) >= m.opt.maxEntries {
		return fmt.Errorf("cache is full")
	}

	m.items[key] = item
	return nil
}

func (m *memory) getItem(key string) (*item, error) {
	i, ok := m.items[key]
	if !ok {
		return nil, fmt.Errorf("%s not exist", key)
	}

	if i.Expired.Before(time.Now()) {
		//过期后删除
		err := m.del(key)
		return nil, err
	}
	return i, nil
}

func (m *memory) Set(key string, val any, expire time.Duration) error {
	key = m.getCacheKey(key)
	m.mutex.Lock()
	defer m.mutex.Unlock()

	item := &item{
		Value:   val,
		Expired: time.Now().Add(expire),
	}
	return m.setItem(key, item)
}

func (m *memory) Get(key string) (any, error) {
	key = m.getCacheKey(key)
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	item, err := m.getItem(key)
	if err != nil || item == nil {
		return nil, err
	}

	return item.Value, nil
}

func (m *memory) Delete(key string) error {
	key = m.getCacheKey(key)

	return m.del(key)
}

func (m *memory) del(key string) error {
	delete(m.items, key)
	return nil
}

func (m *memory) Increase(key string, step int64) error {
	key = m.getCacheKey(key)
	return m.calculate(key, step)
}

func (m *memory) Decrease(key string, step int64) error {
	key = m.getCacheKey(key)
	return m.calculate(key, step)
}

func (m *memory) calculate(key string, num int64) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	item, err := m.getItem(key)
	if err != nil {
		return err
	}
	if item == nil {
		return fmt.Errorf("%s not exist", key)
	}

	var n int64
	n, err = cast.ToInt64E(item.Value)
	if err != nil {
		return err
	}

	item.Value = n + num
	return m.setItem(key, item)
}

func (m *memory) Expire(key string, dur time.Duration) error {
	key = m.getCacheKey(key)
	item, err := m.getItem(key)
	if err != nil {
		return err
	}
	if item == nil {
		return fmt.Errorf("%s not exist", key)
	}

	// 修改操作
	m.mutex.Lock()
	defer m.mutex.Unlock()

	item.Expired = time.Now().Add(dur)
	return nil
}

func (m *memory) Handler() any {
	return m
}

// NewMemory memory模式
func NewMemory() contract.CacheHandlerInterface {
	return &memory{
		items: make(map[string]*item),
		opt:   setDefaultOption(),
	}
}
