package remote

import (
	"time"

	"github.com/cpyun/gyopls-core/config/source"
	"github.com/spf13/viper"
)

type remote struct {
	opts source.Options
}

func (r *remote) Read() (*source.ChangeSet, error) {
	v := viper.GetViper()
	v.SetConfigType("properties")

	provider := r.opts.Context.Value(remoteProvider{}).(remoteProvider)
	err := v.AddRemoteProvider(provider.name, provider.endpoint, provider.path)
	if err != nil {
		return nil, err
	}

	err = v.ReadRemoteConfig()

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

func NewSourceRemote(opts ...source.Option) source.Source {
	options := source.NewOptions(opts...)

	return &remote{
		opts: options,
	}
}
