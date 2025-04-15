package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type optionFunc func(*DBManager)

type dbManagerOptions struct {
	dialector map[string]func(string) gorm.Dialector
	logger    logger.Interface
}

func setDefaultDbManagerOptions() dbManagerOptions {
	var dialector = make(map[string]func(string) gorm.Dialector)
	dialector["mysql"] = mysql.Open

	return dbManagerOptions{
		dialector: dialector,
	}
}

func WithConectorOpts(name string, open func(string) gorm.Dialector) optionFunc {
	return func(dm *DBManager) {
		dm.opts.dialector[name] = open
	}
}

func WithLogger(logger logger.Interface) optionFunc {
	return func(dm *DBManager) {
		dm.opts.logger = logger
	}
}
