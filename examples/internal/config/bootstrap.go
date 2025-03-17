package config

import (
	"log"

	"github.com/cpyun/gyopls-core/config"
	"github.com/cpyun/gyopls-core/config/source/file"
	"github.com/cpyun/gyopls-core/database/db"
)

var (
	ExtendConfig interface{}
	// Settings     *Config
)

type Bootstarp struct {
	Application *Application                     `mapstructure:"application" json:"application" yaml:"application"`
	Logger      *Logger                          `mapstructure:"logger" json:"logger" yaml:"logger"`
	Cache       *Cache                           `mapstructure:"cache" yaml:"cache" json:"cache"`
	Databases   map[string]*db.ConnectionOptions `json:"databases" yaml:"databases"`
	Filesystem  *Filesystem                      `mapstructure:"filesystem" json:"mysql" yaml:"filesystem"`
	Redis       *Redis                           `mapstructure:"redis" json:"redis" yaml:"redis"`
	Server      map[string]Server                `mapstructure:"server" json:"server" yaml:"server"`
	Extend      interface{}                      `yaml:"extend"`
}

func (c *Bootstarp) OnChange() {
	//c.init()
	log.Println("config change and reload")
}

func (c *Bootstarp) init() {
	c.Logger.Setup()
}

// Setup
// @description   Setup 载入配置文件
// @auth      cpYun             时间（2022/7/22   10:00 ）
// @param     s         string        "配置文件路径"
// @param     fs        func          "回调函数"
// @return
func Setup() {
	Settings := &Bootstarp{
		Application: ApplicationConfig,
		Databases:   DatabasesConfig,
		Cache:       CacheConfig,
		Logger:      LoggerConfig,
		Filesystem:  FilesystemConfig,
		Redis:       RedisConfig,
		Extend:      ExtendConfig,
	}

	//
	c := config.New(config.WithSourceOpts(
		file.NewSourceFile(file.WithPath("test/data/settings.yaml")),
	))
	if err := c.Load(); err != nil {
		// panic(err)
	}
	if err := c.Scan(Settings); err != nil {
		// panic(err)
	}
	//

	// 初始化配置
	Settings.init()
}
