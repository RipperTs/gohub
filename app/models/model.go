// Package models 模型通用属性和方法
package models

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID        uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
	DeletedAt gorm.DeletedAt
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
	DeletedAt time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
