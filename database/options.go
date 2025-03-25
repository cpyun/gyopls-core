package database

import (
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type optionFunc func(*DBManager)

type dbManagerOptions struct {
	dialector sync.Map
	logger    logger.Interface
}

func setDefaultDbManagerOptions() dbManagerOptions {
	return dbManagerOptions{}
}

func WithConectorOpts(name string, open func(string) gorm.Dialector) optionFunc {
	return func(dm *DBManager) {
		dm.opts.dialector.Store(name, open)
	}
}

func WithLogger(logger logger.Interface) optionFunc {
	return func(dm *DBManager) {
		dm.opts.logger = logger
	}
}
