package main

import (
	"github.com/gin-gonic/gin"
	"gorsk/server/company"
	"gorsk/server/initDB"
	"gorsk/server/product"
)

func main() {
	//router := gin.Default()
	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))

	db := initDB.InitDB()
	services := initDB.InitServices(db)
	productHandler := product.NewProductHandler(services.ProductService)
	companyHandler := company.NewCompanyHandler(services.CompanyService)

	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/products", productHandler.AddProducts)
		api.GET("/products", productHandler.GetAllProducts)
		api.GET("/products/:id", productHandler.GetProductById)
		api.PUT("/products/:id", productHandler.UpdateProduct)
		api.DELETE("/products/:id", productHandler.DeleteProduct)

		api.POST("/company", companyHandler.AddCompany)
		api.GET("/companies", companyHandler.GetAllCompanies)
		api.GET("/company/:id", companyHandler.GetCompanyById)
		api.PUT("/company/:id", companyHandler.UpdateCompany)
		api.DELETE("/company/:id", companyHandler.DeleteCompany)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//v1.POST("/company", company.AddCompany)
	//v1.GET("/companies", company.GetAllCompanies)
	//v1.GET("/company/:id", company.GetCompanyById)
	//v1.GET("/companies/products", company.GetCompaniesWithProducts)
	//v1.GET("/companies/:name", company.GetCompanyByName)
	//v1.PUT("/company/:id", company.UpdateCompany)
	//v1.DELETE("/company/:id", company.DeleteCompany)

	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
