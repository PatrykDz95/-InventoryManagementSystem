package notification

import (
	"gorsk/server/user"
	"time"
)

type Notification struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint
	User      user.User `gorm:"foreignKey:UserID"`
	Message   string    `gorm:"not null"`
	Read      bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
