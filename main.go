package main

import (
	"github.com/gin-gonic/gin"
	"gorsk/server/database"
)

func main() {
	router := gin.Default()
	database.Init("my_database")

	v1 := router.Group("/api/v1")

	// User routes
	//v1.POST("/product/", product.AddProduct())

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
