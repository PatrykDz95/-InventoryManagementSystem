package database

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorsk/server/audit_log"
	"gorsk/server/category"
	"gorsk/server/common"
	"gorsk/server/company"
	"gorsk/server/inventory"
	"gorsk/server/location"
	"gorsk/server/notification"
	"gorsk/server/order"
	"gorsk/server/orderItem"
	"gorsk/server/product"
	"gorsk/server/product_tag"
	"gorsk/server/supplier"
	"gorsk/server/tag"
	"gorsk/server/user"
	"gorsk/server/warehouse"
	"log"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=localhost user=gorm password=gorm dbname=%s port=9920 sslmode=disable TimeZone=Asia/Shanghai", "my_database")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	err = db.AutoMigrate(&audit_log.AuditLog{}, &category.Category{},
		&company.Company{}, location.Location{}, &notification.Notification{},
		&tag.Tag{}, &orderItem.OrderItem{}, &product.Product{},
		&product_tag.ProductTag{}, &warehouse.Warehouse{}, &supplier.Supplier{},
		&order.Order{}, &inventory.Inventory{})

	if err != nil {
		log.Fatal("Failed to auto-migrate the database:", err)
	}

	common.SetDB(db)
	return db
}

type Services struct {
	CompanyService *company.Service
	ProductService *product.Service
	UserService    *user.Service
}

func InitServices(db *gorm.DB, mongoDB *mongo.Database) *Services {
	return &Services{
		CompanyService: company.NewCompanyService(db),
		ProductService: product.NewProductService(db),
		UserService:    user.NewUserService(mongoDB),
	}
}
