package model

type Product struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"column:name"`
	Quantity  int    `gorm:"column:quantity"`
	CompanyID uint   `gorm:"column:company_id"`
}
