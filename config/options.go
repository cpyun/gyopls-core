package config

import (
	"github.com/cpyun/gyopls-core/config/source"
)

type configOptions struct {
	sources   []source.Source
	callbacks []func()
}

type OptionFunc func(opts *configOptions)

func WithSourceOpts(s ...source.Source) OptionFunc {
	return func(t *configOptions) {
		t.sources = append(t.sources, s...)
	}
}

func WithCallbackOpts(fs ...func()) OptionFunc {
	return func(t *configOptions) {
		t.callbacks = append(t.callbacks, fs...)
	}
}
