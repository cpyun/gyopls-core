package flag

import (
	"encoding/json"
	"time"

	"github.com/cpyun/gyopls-core/config/source"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type flag struct {
	viper *viper.Viper
	sets  *pflag.FlagSet
}

func (f *flag) init() {
	f.viper = viper.GetViper()

	if f.sets != nil {
		f.viper.BindPFlags(f.sets)
	}
}

func (f *flag) Read() (*source.ChangeSet, error) {
	var changes map[string]any
	f.sets.VisitAll(func(flag *pflag.Flag) {
		changes[flag.Name] = flag.Value.String()
	})
	b, err := json.Marshal(changes)

	cs := &source.ChangeSet{
		Format:    "json",
		Source:    f.String(),
		Timestamp: time.Now(),
		Data:      b,
	}
	cs.Checksum = cs.Sum()

	return cs, err
}

func (f *flag) Watch() (source.Watcher, error) {
	return nil, source.ErrWatcherStopped
}

func (f *flag) String() string {
	return "flag"
}

func New(opts ...source.Option) source.Source {
	f := &flag{}

	f.init()
	return f
}
