package filesystem

import (
	"sync"

	"github.com/cpyun/gyopls-core/filesystem"
	"github.com/spf13/afero"
)

var (
	once      sync.Once
	defaultFs *filesystem.FsManager
)

func init() {
	once.Do(func() {
		defaultFs = filesystem.New(
			filesystem.WithDriver("local", afero.NewOsFs()),
		)
	})
}

func Default() afero.Fs {
	return Store("local")
}

// 切换文件系统
func Store(name string) afero.Fs {
	return defaultFs.Store(name)
}
