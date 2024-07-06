package entity

import (
	"gorsk/server/product/product_entity"
)

type Supplier struct {
	ID               uint                     `gorm:"primary_key"`
	Name             string                   `gorm:"column:name"`
	Description      string                   `gorm:"column:description"`
	ContactInfo      string                   `gorm:"column:contact_info"`
	Address          string                   `gorm:"column:location"`
	CompanyID        uint                     `gorm:"column:company_id"`
	ProductsSupplied []product_entity.Product `gorm:"foreignKey:CompanyID"`
}
