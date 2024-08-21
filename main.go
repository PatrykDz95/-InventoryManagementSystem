package main

import (
	"github.com/gin-gonic/gin"
	"gorsk/server/common"
	"gorsk/server/company"
	"gorsk/server/database"
	"gorsk/server/product"
	"gorsk/server/user"
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
	mongoDB := database.InitMongoDB("mongodb://localhost:27017")

	db := database.InitDB()
	services := database.InitServices(db, mongoDB)
	productHandler := product.NewProductHandler(services.ProductService)
	companyHandler := company.NewCompanyHandler(services.CompanyService)
	userHandler := user.NewUserHandler(services.UserService)

	r := gin.Default()
	api := r.Group("/api")
	api.Use(common.JWTMiddleware())
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

	r.POST("/user/register", userHandler.RegisterUser)
	r.POST("/user/login", userHandler.LoginUser)

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
