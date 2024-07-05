package company

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorsk/server/company/company_entite"
	"gorsk/server/database"
	"net/http"
)

func AddCompany(c *gin.Context) {
	db := database.GetDB()
	company, err := bindProductFromJSON(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}
	created := db.Create(&company)
	if created.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": created.Error.Error()})
		return
	}
	sendResponse(c, created)
}

func GetAllCompanies(c *gin.Context) {
	db := database.GetDB()
	var companies []company_entite.Company
	found := db.Find(&companies)
	if found.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": found.Error.Error()})
		return
	}
	sendResponse(c, found)
}

func GetCompaniesWithProducts(c *gin.Context) {
	db := database.GetDB()
	var companies []company_entite.Company
	found := db.Preload("Inventory").Find(&companies)
	if found.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": found.Error.Error()})
		return
	}
	sendResponse(c, found)
}

func bindProductFromJSON(c *gin.Context) (company_entite.Company, error) {
	company := company_entite.Company{}
	err := c.ShouldBindJSON(&company)
	return company, err
}

func sendResponse(c *gin.Context, dbResponse *gorm.DB) {
	if dbResponse.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": dbResponse.Error.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"Company": dbResponse.Statement.Model,
		})
	}
}
