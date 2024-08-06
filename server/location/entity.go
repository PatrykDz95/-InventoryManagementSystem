package location

import (
	"time"
)

type Location struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"unique;not null"`
	Address   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
