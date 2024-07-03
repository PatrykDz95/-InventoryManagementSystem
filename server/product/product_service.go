package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorsk/server/database"
	"gorsk/server/product/product_entity"
	"net/http"
)

func AddProduct(c *gin.Context) {
	db := database.GetDB()
	product, err := bindProductFromJSON(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}
	created := db.Create(&product)
	if created.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": created.Error.Error()})
		return
	}
	sendResponse(c, created, http.StatusOK)
}

func GetAllProducts(c *gin.Context) {
	db := database.GetDB()
	var products []product_entity.Product
	found := db.Find(&products)
	if found.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": found.Error.Error()})
		return
	}
	sendResponse(c, found, http.StatusOK)
}

func GetProductById(c *gin.Context) {
	db := database.GetDB()
	var product product_entity.Product
	found := db.First(&product, c.Param("id"))
	if found.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": found.Error.Error()})
		return
	}
	sendResponse(c, found, http.StatusOK)
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	product, err := bindProductFromJSON(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}
	updated := db.Model(&product).Where("id = ?", c.Param("id")).Updates(product)
	if updated.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": updated.Error.Error()})
		return
	}
	sendResponse(c, updated, http.StatusOK)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	var product product_entity.Product
	//result := db.First(&product, c.Param("id"))
	//if result.Error != nil {
	//	c.JSON(http.StatusNotFound, gin.H{"errorMessage": result.Error.Error()})
	//	return
	//}
	deleted := db.Clauses(clause.Returning{}).Where("id = ?", c.Param("id")).Delete(&product)
	if deleted.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": deleted.Error.Error()})
		return
	}
	if deleted.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": "No product found"})
		return

	}
	sendResponse(c, deleted, http.StatusOK)

}

func bindProductFromJSON(c *gin.Context) (product_entity.Product, error) {
	product := product_entity.Product{}
	err := c.ShouldBindJSON(&product)
	return product, err
}

func sendResponse(c *gin.Context, dbResponse *gorm.DB, status int) {
	if dbResponse.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": dbResponse.Error.Error()})
	} else {
		c.JSON(status, gin.H{
			"status":  status,
			"Product": dbResponse.Statement.Model,
		})
	}
}
