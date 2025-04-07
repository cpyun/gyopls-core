package database

import (
	"sync"

	"github.com/cpyun/gyopls-core/database/db"
	"gorm.io/gorm"
)

type DBManager struct {
	instance sync.Map         // *gorm.DB
	lock     sync.RWMutex     //
	opts     dbManagerOptions //
}

func (t *DBManager) init() {}

func (t *DBManager) applyOption(opts ...optionFunc) {
	for _, v := range opts {
		v(t)
	}
}

func (t *DBManager) getDialector(name string) func(string) gorm.Dialector {
	if val, ok := t.opts.dialector[name]; ok {
		return val
	}
	return nil
}

func (t *DBManager) createConnect(conn *db.Connection) *gorm.DB {
	dialector := t.getDialector(conn.GetDriverName())
	c := conn.Connect(dialector)
	// 设置日志
	if t.opts.logger != nil {
		c.Session(&gorm.Session{Logger: t.opts.logger})
	}

	return c
}

// connect 创建数据库连接查询.
func (t *DBManager) Connect(key string, conn *db.Connection) *DBManager {
	db := t.createConnect(conn)
	t.instance.Store(key, db)
	return t
}

// store 切换数据库连接查询.
func (t *DBManager) Store(name string) *gorm.DB {
	// log.Debug("Database at [%s] => %s", name, pkg.Green(c.Source))
	if val, ok := t.instance.Load(name); ok {
		return val.(*gorm.DB)
	}

	return nil
}

func NewDBManager(opts ...optionFunc) *DBManager {
	m := &DBManager{
		opts: setDefaultDbManagerOptions(),
	}
	//
	m.applyOption(opts...)

	m.init()
	return m
}
