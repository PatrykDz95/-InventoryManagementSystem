package order

import (
	"gorsk/server/supplier"
	"time"
)

type Order struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	UserID      uint
	SupplierID  uint
	Supplier    supplier.Supplier `gorm:"foreignKey:SupplierID"`
	Status      string            `gorm:"not null"`
	TotalAmount float64           `gorm:"not null"`
	CreatedAt   time.Time         `gorm:"autoCreateTime"`
	UpdatedAt   time.Time         `gorm:"autoUpdateTime"`
}
