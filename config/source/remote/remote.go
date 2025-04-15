package remote

import (
	"time"

	"github.com/cpyun/gyopls-core/config/source"
	"github.com/spf13/viper"
)

type remote struct {
	viper     *viper.Viper
	providers []remoteProvider
}

func (r *remote) applyOptions(opts ...optionFn) {

}

func (r *remote) init() {
	r.viper = viper.GetViper()
	r.viper.SetConfigType("properties")

	if len(r.providers) > 0 {
		for _, p := range r.providers {
			_ = r.viper.AddRemoteProvider(p.name, p.endpoint, p.path)
		}
	}
}

func (r *remote) Read() (*source.ChangeSet, error) {
	err := r.viper.ReadRemoteConfig()

	cs := &source.ChangeSet{
		Format:    "json",
		Source:    r.String(),
		Timestamp: time.Now(),
		Data:      []byte("viper"),
	}
	cs.Checksum = cs.Sum()

	return cs, err
}

func (r *remote) Watch() (source.Watcher, error) {
	// not supported
	return nil, source.ErrWatcherStopped
}

func (r *remote) String() string {
	return "remote"
}

func New(opts ...optionFn) source.Source {
	r := &remote{}
	r.applyOptions(opts...)

	r.init()
	return r
}
