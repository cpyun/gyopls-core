package database

import (
	"sync"

	"github.com/cpyun/gyopls-core/database/db"
	"gorm.io/gorm"
)

type DBManager struct {
	instance sync.Map         // *gorm.DB
	mux      sync.RWMutex     //
	opts     dbManagerOptions //
}

func (t *DBManager) init() {}

func (t *DBManager) withOptionFunc(opts ...optionFunc) {
	for _, v := range opts {
		v(t)
	}
}

func (t *DBManager) getDialector(name string) func(string) gorm.Dialector {
	if val, ok := t.opts.dialector.Load(name); ok {
		return val.(func(string) gorm.Dialector)
	}
	return nil
}

// connect 创建数据库连接查询.
func (t *DBManager) Connect(name string, conn *db.Connection) *DBManager {
	dialector := t.getDialector(conn.GetDriverName())
	c := conn.Connect(dialector)
	// 设置日志
	if t.opts.logger != nil {
		c.Session(&gorm.Session{Logger: t.opts.logger})
	}

	t.instance.Store(name, c)
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
	m.withOptionFunc(opts...)

	m.init()
	return m
}
