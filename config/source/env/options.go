package env

import (
	"strings"
)

type optionFn func(*envOptions)

type envOptions struct {
	prefix   string
	replacer *strings.Replacer
}

func setDefaultOption() envOptions {
	return envOptions{
		replacer: strings.NewReplacer(".", "_"),
	}
}

// WithPrefix sets the prefix for the environment variable.
func WithPrefix(p string) optionFn {
	return func(o *envOptions) {
		o.prefix = p
	}
}

// WithReplace sets the replacer for the environment variable.
func WithReplace(oldNew ...string) optionFn {
	return func(o *envOptions) {
		o.replacer = strings.NewReplacer(oldNew...)
	}
}
