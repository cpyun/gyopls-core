package db

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

type Connection struct {
	// dialector func(string) gorm.Dialector
	opts ConnectionOptions
}

func (t *Connection) GetDriverName() string {
	return t.opts.Driver
}

// connect
func (t *Connection) Connect(dialector func(string) gorm.Dialector) *gorm.DB {
	if dialector == nil {
		return nil
	}
	config := t.getGormOption()
	return t.createConnect(dialector(t.opts.Dsn), config)
}

// create connection
func (t *Connection) createConnect(dialector gorm.Dialector, config *gorm.Config) *gorm.DB {
	db, err := gorm.Open(dialector, config)
	if err != nil {
		panic(err)
	}
	//
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	register := dbresolver.Register(dbresolver.Config{})
	// 设置连接池
	if t.opts.ConnMaxIdletime > 0 {
		sqlDB.SetConnMaxIdleTime(time.Duration(t.opts.ConnMaxIdletime) * time.Second)
		register.SetConnMaxIdleTime(time.Duration(t.opts.ConnMaxIdletime) * time.Second)
	}
	if t.opts.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxIdleTime(time.Duration(t.opts.ConnMaxLifetime) * time.Second)
		register.SetConnMaxLifetime(time.Duration(t.opts.ConnMaxLifetime) * time.Second)
	}
	if t.opts.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(t.opts.MaxOpenConns)
		register.SetMaxOpenConns(t.opts.MaxOpenConns)
	}
	if t.opts.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(t.opts.MaxIdleConns)
		register.SetMaxIdleConns(t.opts.MaxIdleConns)
	}
	// 分布式
	for _, v := range t.opts.Registers {
		register.Register(v)
	}

	//
	if register != nil {
		db.Use(register)
	}
	return db
}

func (t *Connection) getGormOption() *gorm.Config {
	gc := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "sd_",
			SingularTable: true,
			//NoLowerCase: true, // 关闭小写转换
			//NameReplacer: nil,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		//SkipDefaultTransaction: true,		//跳过默认事务
		// Logger: nil,
	}
	return gc
}

func (t *Connection) withOptionFunc(opts ...OptionFunc) {
	for _, v := range opts {
		v(t)
	}
}

func NewConnection(opts ...OptionFunc) *Connection {
	c := &Connection{
		opts: setDefaultOptions(),
	}
	c.withOptionFunc(opts...)
	return c
}
