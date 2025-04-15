package env

import (
	"time"

	"github.com/cpyun/gyopls-core/config/source"
	"github.com/spf13/viper"
)

type env struct {
	viper *viper.Viper
	opts  envOptions
}

func (e *env) applyOption(opts ...optionFn) {
	for _, o := range opts {
		o(&e.opts)
	}
}

func (e *env) init() {
	// 设置前缀
	if e.opts.prefix != "" {
		e.viper.SetEnvPrefix(e.opts.prefix)
	}

	// 使用替代符替换
	if e.opts.replacer != nil {
		e.viper.SetEnvKeyReplacer(e.opts.replacer)
	}
}

func (e *env) Read() (*source.ChangeSet, error) {
	// 自动加载环境变量
	e.viper.AutomaticEnv()

	cs := &source.ChangeSet{
		Format:    "json",
		Source:    e.String(),
		Timestamp: time.Now(),
		Data:      []byte("viper"),
	}
	cs.Checksum = cs.Sum()

	return cs, nil
}

func (e *env) Watch() (source.Watcher, error) {
	return nil, source.ErrWatcherStopped
}

func (e *env) String() string {
	return "env"
}

func New(opts ...optionFn) source.Source {
	e := &env{
		viper: viper.GetViper(),
		opts:  setDefaultOption(),
	}
	e.applyOption(opts...)

	e.init()
	return e
}
