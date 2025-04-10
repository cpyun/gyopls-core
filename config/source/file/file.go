package file

import (
	"errors"
	"path/filepath"
	"time"

	"github.com/cpyun/gyopls-core/config/source"
	"github.com/spf13/viper"
)

type file struct {
	viper *viper.Viper
	opts  fileOptions
}

func (f *file) applyOption(opts ...optionFn) {
	for _, o := range opts {
		o(&f.opts)
	}
}

func (f *file) init() {
	if f.opts.file != "" {
		f.viper.SetConfigFile(f.opts.file)
	}
}

func (f *file) Read() (*source.ChangeSet, error) {
	err := f.viper.ReadInConfig()
	if err != nil || errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return nil, err
	}

	cs := &source.ChangeSet{
		Format:    filepath.Ext(f.opts.file),
		Source:    f.String(),
		Timestamp: time.Now(),
		Data:      []byte("viper"),
	}
	cs.Checksum = cs.Sum()

	return cs, nil
}

func (f *file) Watch() (source.Watcher, error) {
	return newWatcher(f)
}

func (f *file) Write(_ *source.ChangeSet) error {
	return nil
}

func (f *file) String() string {
	return "file"
}

func New(opts ...optionFn) source.Source {
	f := &file{
		viper: viper.GetViper(),
		opts:  setDefaultOptions(),
	}
	f.applyOption(opts...)

	f.init()
	return f
}
