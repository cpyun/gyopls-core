package config

import (
	"log"

	"github.com/cpyun/gyopls-core/database/db"
)

type Config struct {
	Application *Application                     `mapstructure:"application" json:"application" yaml:"application"`
	Logger      *Logger                          `mapstructure:"logger" json:"logger" yaml:"logger"`
	Cache       *Cache                           `mapstructure:"cache" yaml:"cache" json:"cache"`
	Databases   map[string]*db.ConnectionOptions `json:"databases" yaml:"databases"`
	Filesystem  *Filesystem                      `mapstructure:"filesystem" json:"mysql" yaml:"filesystem"`
	Redis       *Redis                           `mapstructure:"redis" json:"redis" yaml:"redis"`
	Server      map[string]Server                `mapstructure:"server" json:"server" yaml:"server"`
	Extend      interface{}                      `yaml:"extend"`
}

func (c *Config) OnChange() {
	//c.init()
	log.Println("config change and reload")
}

func (c *Config) init() {
	c.Logger.Setup()
}
