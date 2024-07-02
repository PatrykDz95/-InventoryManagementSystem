package company

import (
	"github.com/gin-gonic/gin"
	"gorsk/server/company/company_entite"
	"gorsk/server/database"
	"net/http"
)

func AddCompany(c *gin.Context) {
	db := database.GetDB()
	company := company_entite.Company{}
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}
	created := db.Create(&company)
	if created.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": created.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"Product": created.Statement.Model,
	})
}
