package product_entity

type Product struct {
	ID          uint    `gorm:"primary_key"`
	Name        string  `gorm:"column:name"`
	Description string  `gorm:"column:description"`
	Price       float32 `gorm:"column:price"`
	Quantity    int     `gorm:"column:quantity"`
	CompanyID   uint    `gorm:"column:company_id"`
}
