package product_tag

type ProductTag struct {
	ProductID uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey"`
}
