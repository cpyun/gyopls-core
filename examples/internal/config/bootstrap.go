package config

import (
	"github.com/cpyun/gyopls-core/config"
	"github.com/cpyun/gyopls-core/config/source/file"
)

var (
	ExtendConfig interface{}
	// Settings     *Config
)

// @title Setup
// @description   Setup 载入配置文件
// @auth	cpYun	2022/7/22 10:00
func Setup() {
	Settings := Config{
		Application: ApplicationConfig,
		Databases:   DatabasesConfig,
		Cache:       CacheConfig,
		Logger:      LoggerConfig,
		Filesystem:  FilesystemConfig,
		Redis:       RedisConfig,
		Extend:      ExtendConfig,
	}

	//
	c := config.New(config.WithSources(
		file.New(file.WithFile("test/data/settings.yaml")),
	))
	if err := c.Load(); err != nil {
		// panic(err)
	}
	if err := c.Scan(&Settings); err != nil {
		// panic(err)
	}
	//

	// 初始化配置
	Settings.init()
}
