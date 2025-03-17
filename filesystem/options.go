package filesystem

import (
	"sync"

	"github.com/spf13/afero"
)

type filesystemAdapterOptions struct {
	drivers sync.Map
}

type OptionFunc func(*FilesystemAdapter)

func setDefaultOptions() filesystemAdapterOptions {
	return filesystemAdapterOptions{}
}

func WithDriversOpts(name string, fs afero.Fs) OptionFunc {
	return func(o *FilesystemAdapter) {
		o.opts.drivers.Store(name, fs)
	}
}
