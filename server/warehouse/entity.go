package warehouse

type Warehouse struct {
	ID          uint   `gorm:"primary_key;autoIncrement"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Address     string `gorm:"column:location"`
	CompanyID   uint   `gorm:"column:company_id"`
}
