package company_entity

import (
	"gorsk/server/product/product_entity"
)

type Company struct {
	ID        uint                     `gorm:"primary_key"`
	Name      string                   `gorm:"column:name"`
	IsActive  bool                     `gorm:"column:is_active"`
	AddedDate string                   `gorm:"column:added_date"`
	Inventory []product_entity.Product `gorm:"foreignKey:CompanyID"`
}
