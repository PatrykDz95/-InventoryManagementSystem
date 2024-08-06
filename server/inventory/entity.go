package inventory

import (
	"gorsk/server/location"
	"gorsk/server/product"
	"time"
)

type Inventory struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	ProductID  uint
	Product    product.Product `gorm:"foreignKey:ProductID"`
	LocationID uint
	Location   location.Location `gorm:"foreignKey:LocationID"`
	Quantity   int               `gorm:"not null"`
	CreatedAt  time.Time         `gorm:"autoCreateTime"`
	UpdatedAt  time.Time         `gorm:"autoUpdateTime"`
}
