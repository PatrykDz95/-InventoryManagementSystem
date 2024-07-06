package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorsk/server/company/company_entity"
	"gorsk/server/product/product_entity"
)

var DB *gorm.DB

func Init(dbName string) {
	// Connect to postgresql database
	dsn := fmt.Sprintf("host=localhost user=gorm password=gorm dbname=%s port=9920 sslmode=disable TimeZone=Asia/Shanghai", dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	rows, err := db.Raw("SELECT datname FROM pg_database WHERE datname = ?", dbName).Rows()
	if err != nil {
		panic("Failed to execute query")
	}

	if !rows.Next() {
		db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
	}

	// AutoMigrate the Product struct
	err = db.AutoMigrate(&product_entity.Product{})
	if err != nil {
		panic("Failed to migrate database!")
	}

	// AutoMigrate the Company struct
	err = db.AutoMigrate(&company_entity.Company{})
	if err != nil {
		panic("Failed to migrate database!")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get SQL database from GORM")
	}
	sqlDB.Close()

	dsn = fmt.Sprintf("host=localhost user=gorm password=gorm dbname=%s port=9920 sslmode=disable TimeZone=Asia/Shanghai", dbName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
}

func GetDB() *gorm.DB {
	return DB
}
