package model

import (
	"time"
)

type Model struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"column:create_time; autoCreateTime; <-:create"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"column:update_time; autoUpdateTime; <-;" `
	//DeleteTime	gorm.DeleteAt 	`json:"-" gorm:"index;comment:删除时间"`
}

type ControlBy struct {
	CreateBy int `json:"create_by" gorm:"index;comment:创建者"`
	UpdateBy int `json:"update_by" gorm:"index;comment:更新者"`
}

// SetCreateBy 设置创建人id
func (e *ControlBy) SetCreateBy(createBy int) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy int) {
	e.UpdateBy = updateBy
}
