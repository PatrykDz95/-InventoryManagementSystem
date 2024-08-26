package product

import (
	"gorsk/server/category"
	"time"
)

type Product struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	CompanyID   uint   `gorm:"not null"`
	CategoryID  uint   `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string
	Price       float64           `gorm:"not null"`
	Category    category.Category `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time         `gorm:"autoCreateTime"`
	UpdatedAt   time.Time         `gorm:"autoUpdateTime"`
}
