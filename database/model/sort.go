package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// FieldSortDest 自定义字段排序
func FieldSortDest(field string, desc bool) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: field,
			},
			Desc: desc,
		})
	}
}
