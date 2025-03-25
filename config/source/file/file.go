package file

import (
	"errors"
	"path/filepath"
	"time"

	"github.com/cpyun/gyopls-core/config/source"
	"github.com/spf13/viper"
)

type file struct {
	path string
	opts source.Options
}

func (f *file) Read() (*source.ChangeSet, error) {
	viper.SetConfigFile(f.path)

	err := viper.ReadInConfig()
	if err != nil || errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return nil, err
	}

	cs := &source.ChangeSet{
		Format:    filepath.Ext(f.path),
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

func NewSourceFile(opts ...source.Option) source.Source {
	options := source.NewOptions(opts...)
	path := "./config/"

	if fk, ok := options.Context.Value(filePathKey{}).(string); ok {
		path = fk
	}

	return &file{
		path: path,
		opts: options,
	}
}
