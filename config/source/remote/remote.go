package remote

import (
	"time"

	"github.com/cpyun/gyopls-core/config"
	"github.com/spf13/viper"
)

type remote struct {
	viper     *viper.Viper
	providers []remoteProvider
}

func (r *remote) applyOptions(opts ...optionFn) {

}

func (r *remote) init() {
	r.viper.SetConfigType("properties")

	if len(r.providers) > 0 {
		for _, p := range r.providers {
			_ = r.viper.AddRemoteProvider(p.name, p.endpoint, p.path)
		}
	}
}

func (r *remote) Load() (*config.ChangeSet, error) {
	err := r.viper.ReadRemoteConfig()

	cs := &config.ChangeSet{
		Format:    "json",
		Source:    r.String(),
		Timestamp: time.Now(),
		Data:      []byte("viper"),
	}
	cs.Checksum = cs.Sum()

	return cs, err
}

func (r *remote) Watch() (config.Watcher, error) {
	// not supported
	return nil, config.ErrWatcherStopped
}

func (r *remote) String() string {
	return "remote"
}

func New(opts ...optionFn) config.Source {
	r := &remote{
		viper: viper.GetViper(),
	}
	r.applyOptions(opts...)

	r.init()
	return r
}
