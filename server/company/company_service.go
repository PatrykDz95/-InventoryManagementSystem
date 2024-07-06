package company

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorsk/server/company/company_entity"
	"gorsk/server/database"
	"log"
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
	log.Println("Created company: ", company.Name)
	sendResponse(c, created)
}

func GetAllCompanies(c *gin.Context) {
	db := database.GetDB()
	var companies []company_entity.Company
	found := db.Find(&companies)
	if found.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": found.Error.Error()})
		return
	}
	log.Println("Found companies: ", companies)
	sendResponse(c, found)
}

func GetCompanyById(c *gin.Context) {
	db := database.GetDB()
	var company company_entity.Company
	found := db.First(&company, c.Param("id"))
	if found.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": found.Error.Error()})
		return
	}
	log.Println("Found company by id: ", company.Name)
	sendResponse(c, found)
}

func GetCompaniesWithProducts(c *gin.Context) {
	db := database.GetDB()
	var companies []company_entity.Company
	found := db.Preload("Inventory").Find(&companies)
	if found.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": found.Error.Error()})
		return
	}
	log.Println("Found companies with products: ", companies)
	sendResponse(c, found)
}

func GetCompanyByName(c *gin.Context) {
	db := database.GetDB()
	var company []company_entity.Company
	found := db.Where("name LIKE ?", "%"+c.Param("name")+"%").Find(&company)
	if found.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": found.Error.Error()})
		return
	}
	log.Println("Found companies: ", company)
	sendResponse(c, found)
}

func UpdateCompany(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")
	company, err := bindProductFromJSON(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}

	// Remove the inventory items if present
	company.Inventory = nil

	updated := db.Model(&company).Where("id = ?", id).Updates(company)
	if updated.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": updated.Error.Error()})
		return
	}
	log.Println("Updated company: ", company.Name)
	sendResponse(c, updated)
}

func DeleteCompany(c *gin.Context) {
	db := database.GetDB()
	var company company_entity.Company
	result := db.First(&company, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": result.Error.Error()})
		return
	}

	// Delete first associated products in one batch
	deleteInventory := db.Where("company_id = ?", &company.ID).Delete(&company.Inventory)
	if deleteInventory.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": deleteInventory.Error.Error()})
		return
	}
	log.Println("Deleted inventory items: ", deleteInventory.RowsAffected, " for company id: ", company.ID, " - ", company.Name)

	deleted := db.Delete(&company, c.Param("id"))
	if deleted.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": deleted.Error.Error()})
		return
	}
	log.Println("Deleted company: ", company.Name)
	sendResponse(c, deleted)
}

func bindProductFromJSON(c *gin.Context) (company_entity.Company, error) {
	company := company_entity.Company{}
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
