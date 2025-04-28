package config

import "sync"

type reader struct {
	values map[string]any
	lock   sync.Mutex
}

func newReader() *reader {
	return &reader{
		values: make(map[string]any),
		lock:   sync.Mutex{},
	}
}

func (r *reader) Value(path string) (Value, bool) {
	r.lock.Lock()
	defer r.lock.Unlock()
	av := &atomicValue{}
	// av.Store()
	return av, true
}

func (r *reader) Source() ([]byte, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	return []byte{}, nil
}
