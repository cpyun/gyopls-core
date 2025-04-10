package file

import (
	"github.com/cpyun/gyopls-core/config/source"
)

type optionFn func(*fileOptions)

type fileOptions struct {
	source.Option
	file string
	path []string
}

func setDefaultOptions() fileOptions {
	return fileOptions{
		file: "./config/settings.yaml",
		path: make([]string, 0),
	}
}

func WithFile(f string) optionFn {
	return func(o *fileOptions) {
		o.file = f
	}
}
