package company

import (
	"gorsk/server/product"
	"time"
)

type Company struct {
	ID        uint              `gorm:"primaryKey;autoIncrement"`
	Name      string            `gorm:"column:name"`
	IsActive  bool              `gorm:"column:is_active"`
	AddedDate time.Time         `gorm:"autoCreateTime"`
	Products  []product.Product `gorm:"foreignKey:CompanyID"`
}
