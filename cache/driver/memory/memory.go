package memory

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/spf13/cast"
)

type item struct {
	Value   string
	Expired time.Time
}

type memory struct {
	items  sync.Map
	mutex  sync.RWMutex
	prefix string
}

// NewMemory memory模式
func NewMemory() *memory {
	return &memory{}
}

func (*memory) String() string {
	return "memory"
}

func (r *memory) getCacheKey(key string) string {
	return r.prefix + key
}

func (m *memory) getItem(key string) (*item, error) {
	var err error
	i, ok := m.items.Load(key)
	if !ok {
		return nil, nil
	}
	switch i.(type) {
	case *item:
		item := i.(*item)
		if item.Expired.Before(time.Now()) {
			//过期
			_ = m.del(key)
			//过期后删除
			return nil, nil
		}
		return item, nil
	default:
		err = fmt.Errorf("value of %s type error", key)
		return nil, err
	}
}

func (m *memory) Get(key string) (any, error) {
	key = m.getCacheKey(key)

	item, err := m.getItem(key)
	if err != nil || item == nil {
		return "", err
	}
	return item.Value, nil
}

func (m *memory) Set(key string, val any, expire time.Duration) error {
	key = m.getCacheKey(key)

	s, err := cast.ToStringE(val)
	if err != nil {
		return err
	}
	item := &item{
		Value:   s,
		Expired: time.Now().Add(time.Duration(expire) * time.Second),
	}
	return m.setItem(key, item)
}

func (m *memory) setItem(key string, item *item) error {
	m.items.Store(key, item)
	return nil
}

func (m *memory) Delete(key string) error {
	key = m.getCacheKey(key)

	return m.del(key)
}

func (m *memory) del(key string) error {
	m.items.Delete(key)
	return nil
}

func (m *memory) HashGet(hk, key string) (any, error) {
	key = m.getCacheKey(key)

	item, err := m.getItem(hk + key)
	if err != nil || item == nil {
		return "", err
	}
	return item.Value, err
}

func (m *memory) HashDelete(hk, key string) error {
	key = m.getCacheKey(key)

	return m.del(hk + key)
}

func (m *memory) Increase(key string, step int) error {
	key = m.getCacheKey(key)

	return m.calculate(key, step)
}

func (m *memory) Decrease(key string, step int) error {
	key = m.getCacheKey(key)

	return m.calculate(key, step)
}

func (m *memory) calculate(key string, num int) error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	item, err := m.getItem(key)
	if err != nil {
		return err
	}

	if item == nil {
		err = fmt.Errorf("%s not exist", key)
		return err
	}
	var n int
	n, err = cast.ToIntE(item.Value)
	if err != nil {
		return err
	}
	n += num
	item.Value = strconv.Itoa(n)
	return m.setItem(key, item)
}

func (m *memory) Expire(key string, dur time.Duration) error {
	key = m.getCacheKey(key)

	m.mutex.RLock()
	defer m.mutex.RUnlock()
	item, err := m.getItem(key)
	if err != nil {
		return err
	}
	if item == nil {
		err = fmt.Errorf("%s not exist", key)
		return err
	}
	item.Expired = time.Now().Add(dur)
	return m.setItem(key, item)
}

func (m *memory) Handler() any {
	return m
}
