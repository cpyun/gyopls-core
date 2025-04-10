package filesystem

import (
	"sync"

	"github.com/spf13/afero"
)

type fsOptions struct {
	drivers sync.Map
}

type OptionFunc func(*FsManager)

func setDefaultOptions() fsOptions {
	return fsOptions{}
}

func WithDriver(name string, fs afero.Fs) OptionFunc {
	return func(o *FsManager) {
		o.opts.drivers.Store(name, fs)
	}
}
