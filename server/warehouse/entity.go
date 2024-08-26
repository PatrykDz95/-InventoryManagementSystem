package warehouse

type Warehouse struct {
	ID          uint   `gorm:"primary_key;autoIncrement"`
	LocationID  uint   `gorm:"column:location_id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Address     string `gorm:"column:location"`
	CompanyID   uint   `gorm:"column:company_id"`
}
