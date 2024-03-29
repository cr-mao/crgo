package models

import (
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at,omitempty"`
}
