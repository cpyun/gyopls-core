package filesystem

import (
	"github.com/spf13/afero"
)

type FilesystemAdapter struct {
	client afero.Fs
	opts   filesystemAdapterOptions
}

func (t *FilesystemAdapter) withOptions(opts ...OptionFunc) {
	for _, opt := range opts {
		opt(t)
	}
}

// 连接或者切换缓存
func (t *FilesystemAdapter) Store(name string) afero.Fs {
	if val, ok := t.opts.drivers.Load(name); ok {
		return val.(afero.Fs)
	}

	return nil
}

func NewFile(opts ...OptionFunc) *FilesystemAdapter {
	t := &FilesystemAdapter{}
	t.opts = setDefaultOptions()
	t.withOptions(opts...)

	// 初始化
	// t.Store(t.getDefaultDriverName())
	return t
}
