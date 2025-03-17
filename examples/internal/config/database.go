package config

import (
	"github.com/cpyun/gyopls-core/database/db"
)

var (
	// DatabaseConfig  = new(db.ConnectionOptions)
	DatabasesConfig = make(map[string]*db.ConnectionOptions)
)
