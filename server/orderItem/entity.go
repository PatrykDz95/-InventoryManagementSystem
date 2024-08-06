package orderItem

import (
	"gorsk/server/order"
	"gorsk/server/product"
	"time"
)

type OrderItem struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	OrderID   uint
	Order     order.Order `gorm:"foreignKey:OrderID"`
	ProductID uint
	Product   product.Product `gorm:"foreignKey:ProductID"`
	Quantity  int             `gorm:"not null"`
	Price     float64         `gorm:"not null"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime"`
}
