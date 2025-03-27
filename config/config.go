package config

import (
	"github.com/spf13/viper"
)

type config struct {
	handler *viper.Viper
	opts    configOptions
}

func (t *config) Load() error {
	for _, src := range t.opts.sources {
		if _, err := src.Read(); err != nil {
			return err
		}
	}

	return nil
}

func (t *config) Scan(v any) error {
	return t.handler.Unmarshal(v)
}

// func (t *config) Value(name string) reader.Value {
// 	return t.handler.Get(name)
// }

// func (t *config) Watch(f func()) error {
// 	t.handler.OnConfigChange(func(in fsnotify.Event) {

// 	})
// 	t.handler.WatchConfig()
// 	return nil
// }

func (t *config) Close() error {
	return nil
}

func (t *config) Handler() *viper.Viper {
	return t.handler
}

func New(opts ...OptionFunc) *config {
	c := &config{
		handler: viper.GetViper(),
		opts:    setDefaultOption(),
	}
	//
	for _, f := range opts {
		f(&c.opts)
	}
	return c
}
