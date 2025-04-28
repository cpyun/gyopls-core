package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type config struct {
	handler *viper.Viper
	reader  *reader
	opts    options
}

func (t *config) Load() error {
	for _, src := range t.opts.sources {
		if _, err := src.Load(); err != nil {
			return err
		}

	}

	return nil
}

func (t *config) Scan(v any) error {
	return t.handler.Unmarshal(v)
}

func (t *config) Value(name string) Value {
	av := new(atomicValue)
	av.Store(t.handler.Get(name))
	return av
}

func (t *config) Watch(name string, f func(key string, value Value)) error {
	if t.handler.Get(name) == nil {
		return ErrNotFound
	}

	t.handler.WatchConfig()
	t.handler.OnConfigChange(func(in fsnotify.Event) {
		// v := t.handler.AllSettings()
		f(name, t.Value(name))
	})
	return nil
}

func (t *config) Close() error {
	return nil
}

func New(opts ...OptionFunc) *config {
	c := &config{
		handler: viper.GetViper(),
		reader:  newReader(),
		opts:    setDefaultOption(),
	}
	//
	for _, f := range opts {
		f(&c.opts)
	}
	return c
}
