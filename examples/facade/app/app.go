package app

import (
	"sync"

	"github.com/cpyun/gyopls-core/application"
)

var (
	once       sync.Once
	defaultApp *application.App
)

func init() {
	once.Do(func() {
		defaultApp = application.New()
	})
}

// 版本
func Version() string {
	return defaultApp.Version()
}

func IsDebug() bool {
	return defaultApp.IsDebug()
}

func Initialized() bool {
	return defaultApp.Initialized()
}

// 初始化
func Initialize() {
	defaultApp.Initialize()
}
