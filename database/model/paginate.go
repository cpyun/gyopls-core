package model

import "gorm.io/gorm"

// Paginate 生成分页
func Paginate(pageSize, pageIndex int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (pageIndex - 1) * pageSize
		if offset < 0 {
			offset = 0
		}
		return db.Offset(offset).Limit(pageSize)
	}
}
