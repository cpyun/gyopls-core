package filesystem

import (
	"sync"

	"github.com/cpyun/gyopls-core/examples/internal/config"
	"github.com/cpyun/gyopls-core/filesystem"
	"github.com/spf13/afero"
)

var (
	once              sync.Once
	defaultName       string
	defaultFilesystem *filesystem.FilesystemAdapter
	Fs                afero.Fs
)

func init() {
	once.Do(func() {
		defaultName = config.FilesystemConfig.Driver
		defaultFilesystem = filesystem.NewFile(
			filesystem.WithDriversOpts("local", afero.NewOsFs()),
		)

		Fs = defaultFilesystem.Store(defaultName)
	})
}

func Store() afero.Fs {
	return defaultFilesystem.Store(defaultName)
}
