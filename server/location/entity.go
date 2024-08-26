package location

import (
	"gorsk/server/inventory"
	"gorsk/server/warehouse"
	"time"
)

type Location struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"unique;not null"`
	WarehouseID uint
	Warehouse   warehouse.Warehouse   `gorm:"foreignKey:WarehouseID"`
	Inventory   []inventory.Inventory `gorm:"foreignKey:LocationID"`
	CreatedAt   time.Time             `gorm:"autoCreateTime"`
	UpdatedAt   time.Time             `gorm:"autoUpdateTime"`
}
