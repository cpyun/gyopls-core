package db

import (
	"gorm.io/plugin/dbresolver"
)

type OptionFunc func(*Connection)

type ConnectionOptions struct {
	Driver          string              `mapstructure:"driver" json:"driver" yaml:"driver"`                                     //
	Dsn             string              `mapstructure:"dsn" json:"dsn" yaml:"dsn"`                                              //
	ConnMaxIdletime int                 `mapstructure:"conn-max-idle-time" json:"conn-max-idle-time" yaml:"conn-max-idle-time"` //
	ConnMaxLifetime int                 `mapstructure:"conn-max-life-time" json:"conn-max-life-time" yaml:"conn-max-life-time"` //
	MaxIdleConns    int                 `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`             // 空闲中的最大连接数
	MaxOpenConns    int                 `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`             // 打开到数据库的最大连接数
	LoggerMode      string              `mapstructure:"logger-mode" json:"logger-mode" yaml:"logger-mode"`                      //
	Registers       []dbresolver.Config `mapstructure:"registers" json:"registers" yaml:"registers"`                            //
}

func setDefaultOptions() ConnectionOptions {
	return ConnectionOptions{
		Driver:          "mysql",
		ConnMaxIdletime: 5,
		ConnMaxLifetime: 10,
		MaxIdleConns:    10,
		MaxOpenConns:    100,
	}
}

func WithConnectConfig(conf ConnectionOptions) OptionFunc {
	return func(c *Connection) {
		if conf.Driver != "" {
			c.opts.Driver = conf.Driver
		}
		if conf.Dsn != "" {
			c.opts.Dsn = conf.Dsn
		}
		if conf.ConnMaxIdletime != 0 {
			c.opts.ConnMaxIdletime = conf.ConnMaxIdletime
		}
		if conf.ConnMaxLifetime != 0 {
			c.opts.ConnMaxLifetime = conf.ConnMaxLifetime
		}
		if conf.MaxIdleConns != 0 {
			c.opts.MaxIdleConns = conf.MaxIdleConns
		}
		if conf.MaxOpenConns != 0 {
			c.opts.MaxOpenConns = conf.MaxOpenConns
		}
		if len(conf.Registers) > 0 {
			c.opts.Registers = conf.Registers
		}
	}
}

func WithDriver(name string) OptionFunc {
	return func(c *Connection) {
		c.opts.Driver = name
	}
}

func WithDSN(dsn string) OptionFunc {
	return func(c *Connection) {
		c.opts.Dsn = dsn
	}
}

func WithConnMaxIdleTime(connMaxIdletime int) OptionFunc {
	return func(c *Connection) {
		c.opts.ConnMaxIdletime = connMaxIdletime
	}
}

func WithConnMaxLifetime(connMaxLifetime int) OptionFunc {
	return func(c *Connection) {
		c.opts.ConnMaxLifetime = connMaxLifetime
	}
}

func WithMaxIdleConns(maxIdleConns int) OptionFunc {
	return func(c *Connection) {
		c.opts.MaxIdleConns = maxIdleConns
	}
}

func WithMaxOpenConns(maxOpenConns int) OptionFunc {
	return func(c *Connection) {
		c.opts.MaxOpenConns = maxOpenConns
	}
}

func WithRegisters(registers []dbresolver.Config) OptionFunc {
	return func(c *Connection) {
		c.opts.Registers = registers
	}
}
