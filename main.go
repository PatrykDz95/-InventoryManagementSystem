package main

import (
	"github.com/gin-gonic/gin"
	"gorsk/server/company"
	"gorsk/server/database"
	"gorsk/server/product"
)

func main() {
	router := gin.Default()
	database.Init("my_database")

	v1 := router.Group("/api/v1")

	// User routes
	v1.POST("/product/", product.AddProduct)
	v1.POST("/company/", company.AddCompany)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
