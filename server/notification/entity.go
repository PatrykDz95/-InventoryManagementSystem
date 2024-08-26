package notification

import (
	"time"
)

type Notification struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	Message   string    `gorm:"not null"`
	Read      bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
