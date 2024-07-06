package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gorsk/server/company"
	"gorsk/server/database"
	"gorsk/server/product"
)

func main() {
	router := gin.Default()
	database.Init("my_database")
	router.Use(cors.Default())

	v1 := router.Group("/api/v1")

	// User routes
	v1.POST("/product", product.AddProduct)
	v1.GET("/products", product.GetAllProducts)
	v1.GET("/product/:id", product.GetProductById)
	v1.PUT("/product/:id", product.UpdateProduct)
	v1.DELETE("/product/:id", product.DeleteProduct)

	v1.POST("/company", company.AddCompany)
	v1.GET("/companies", company.GetAllCompanies)
	v1.GET("/company/:id", company.GetCompanyById)
	v1.GET("/companies/products", company.GetCompaniesWithProducts)
	v1.GET("/companies/:name", company.GetCompanyByName)
	v1.PUT("/company/:id", company.UpdateCompany)
	v1.DELETE("/company/:id", company.DeleteCompany)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
