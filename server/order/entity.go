package order

import (
	"gorsk/server/supplier"
	"gorsk/server/user"
	"time"
)

type Order struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	UserID      uint
	User        user.User `gorm:"foreignKey:UserID"`
	SupplierID  uint
	Supplier    supplier.Supplier `gorm:"foreignKey:SupplierID"`
	Status      string            `gorm:"not null"`
	TotalAmount float64           `gorm:"not null"`
	CreatedAt   time.Time         `gorm:"autoCreateTime"`
	UpdatedAt   time.Time         `gorm:"autoUpdateTime"`
}
