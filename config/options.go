package config

type OptionFunc func(opts *options)

type options struct {
	sources   []Source
	callbacks []func()
}

func setDefaultOption() options {
	return options{
		sources:   make([]Source, 0),
		callbacks: make([]func(), 0),
	}
}

func WithSources(s ...Source) OptionFunc {
	return func(t *options) {
		t.sources = append(t.sources, s...)
	}
}
