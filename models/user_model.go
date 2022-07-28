package models

import (
	"time"
)

//用户session表
type UserSession struct {
	SessionId string    `gorm:"primary_key;column:session_id"`
	Guid      string    `gorm:"column:guid"`
	UserId    int64     `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (UserSession) TableName() string {
	return "user_session"
}
