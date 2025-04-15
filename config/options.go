package config

import (
	"github.com/cpyun/gyopls-core/config/source"
)

type OptionFunc func(opts *configOptions)

type configOptions struct {
	sources   []source.Source
	callbacks []func()
}

func setDefaultOption() configOptions {
	return configOptions{
		sources:   make([]source.Source, 0),
		callbacks: make([]func(), 0),
	}
}

func WithSources(s ...source.Source) OptionFunc {
	return func(t *configOptions) {
		t.sources = append(t.sources, s...)
	}
}
