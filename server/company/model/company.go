package model

import "gorsk/server/product/model"

type Company struct {
	ID        uint            `gorm:"primary_key"`
	Name      string          `gorm:"column:name"`
	IsActive  bool            `gorm:"column:is_active"`
	AddedDate string          `gorm:"column:added_date"`
	Inventory []model.Product `gorm:"foreignKey:CompanyID"`
}
