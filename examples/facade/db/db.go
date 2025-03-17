package db

import (
	"sync"

	"github.com/cpyun/gyopls-core/database"
	"github.com/cpyun/gyopls-core/database/db"
	"github.com/cpyun/gyopls-core/examples/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once      sync.Once
	defaultDB *database.DBManager
	DB        *gorm.DB
)

func init() {
	once.Do(func() {
		defaultDB = database.NewDBManager(
			database.WithConectorOpts("mysql", mysql.Open),
			// database.WithLogger(logger.New(gormLogger.Config{})),
		)

		for k, v := range config.DatabasesConfig {
			defaultDB.Connect(k, db.NewConnection(
				db.WithConnectConfig(*v),
			))
		}

		// 默认mysql
		DB = defaultDB.Store("mysql")
	})
}

func Store(name string) *gorm.DB {
	return defaultDB.Store(name)
}

func Model(name any) *gorm.DB {
	return DB.Model(name)
}

func Tables(name string, args ...any) *gorm.DB {
	return DB.Table(name, args...)
}

// func GetDB() *gorm.DB {
// 	return DB
// }
