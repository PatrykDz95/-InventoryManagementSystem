package product

import (
	"github.com/gin-gonic/gin"
	"gorsk/server/database"
	"gorsk/server/product/product_entity"
	"net/http"
)

func AddProduct(c *gin.Context) {
	db := database.GetDB()
	product := product_entity.Product{}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}
	created := db.Create(&product)
	if created.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": created.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"Product": created.Statement.Model,
	})
}
