package file

import (
	"github.com/cpyun/gyopls-core/config"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type watcher struct {
	f    *file
	exit chan bool
}

func newWatcher(f *file) (config.Watcher, error) {
	w := &watcher{
		f:    f,
		exit: make(chan bool, 1),
	}
	go w.watch()

	return w, nil
}
func (w *watcher) Next() (set *config.ChangeSet, err error) {
	select {
	case <-w.exit:
		set, err = w.f.Load()
		return
	}
}

func (w *watcher) Stop() error {
	close(w.exit)
	return nil
}

func (w *watcher) watch() error {
	viper.OnConfigChange(func(in fsnotify.Event) {
		//log.Println(in.String())
		w.exit <- true
		return
	})
	viper.WatchConfig()
	return nil
}
