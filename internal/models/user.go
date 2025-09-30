package models

import "time"

type (
	User struct {
		ID             int64         `gorm:"primaryKey"`
		ChatID         int64         `gorm:"column:chat_id"`
		Nickname       string        `gorm:"column:nickname"`
		CreatedAt      time.Time     `gorm:"column:created_at"`
		NotifyInterval time.Duration `gorm:"column:notify_interval"`
		NotifiedAt     time.Time     `gorm:"column:notified_at"`
	}
)
