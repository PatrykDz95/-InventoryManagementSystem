package inventory

import (
	"gorsk/server/product"
	"time"
)

type Inventory struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	ProductID  uint
	Product    product.Product `gorm:"foreignKey:ProductID"`
	LocationID uint
	Quantity   int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
