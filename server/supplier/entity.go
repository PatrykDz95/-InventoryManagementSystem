package supplier

import (
	"time"
)

type Supplier struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Name          string `gorm:"not null"`
	ContactPerson string
	Phone         string
	Email         string
	Address       string
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
