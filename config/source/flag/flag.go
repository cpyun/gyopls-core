package flag

import (
	"encoding/json"
	"time"

	"github.com/cpyun/gyopls-core/config"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type flag struct {
	viper *viper.Viper
	sets  *pflag.FlagSet
}

func (f *flag) init() {

	if f.sets != nil {
		f.viper.BindPFlags(f.sets)
	}
}

func (f *flag) Load() (*config.ChangeSet, error) {
	var changes map[string]any
	f.sets.VisitAll(func(flag *pflag.Flag) {
		changes[flag.Name] = flag.Value.String()
	})
	b, err := json.Marshal(changes)

	cs := &config.ChangeSet{
		Format:    "json",
		Source:    f.String(),
		Timestamp: time.Now(),
		Data:      b,
	}
	cs.Checksum = cs.Sum()

	return cs, err
}

func (f *flag) Watch() (config.Watcher, error) {
	return nil, config.ErrWatcherStopped
}

func (f *flag) String() string {
	return "flag"
}

func New(opts ...optionFn) config.Source {
	f := &flag{
		viper: viper.GetViper(),
	}

	f.init()
	return f
}
