package company

import (
	"gorsk/server/product"
)

type Company struct {
	ID        uint              `gorm:"primaryKey;autoIncrement"`
	Name      string            `gorm:"column:name"`
	IsActive  bool              `gorm:"column:is_active"`
	AddedDate string            `gorm:"column:added_date"`
	Products  []product.Product `gorm:"foreignKey:ID"`
}
