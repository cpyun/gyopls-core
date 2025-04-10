package filesystem

import (
	"github.com/spf13/afero"
)

type FsManager struct {
	opts fsOptions
}

func (t *FsManager) applyOptions(opts ...OptionFunc) {
	for _, opt := range opts {
		opt(t)
	}
}

// 连接或者切换缓存
func (t *FsManager) Store(name string) afero.Fs {
	if val, ok := t.opts.drivers.Load(name); ok {
		return val.(afero.Fs)
	}

	return nil
}

func New(opts ...OptionFunc) *FsManager {
	t := &FsManager{
		opts: setDefaultOptions(),
	}
	t.applyOptions(opts...)

	// 初始化
	// t.Store(t.getDefaultDriverName())
	return t
}
