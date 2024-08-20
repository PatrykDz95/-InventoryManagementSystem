package category

import (
	"time"
)

type Category struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	ProductID uint      `gorm:"not null"`
	Name      string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
