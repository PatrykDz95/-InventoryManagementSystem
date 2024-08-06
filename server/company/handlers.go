package company

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Handler struct {
	Service *Service
}

func NewCompanyHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) AddCompany(c *gin.Context) {
	company, err := bindCompanyFromJSON(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}

	// Start a transaction
	tx := h.Service.DB.Begin()
	if err := tx.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}

	// Create the company
	if err := tx.Create(&company).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}

	// Create the products and associate them with the company
	for i := range company.Products {
		//company.Products[i].ID = company.ID
		if err := tx.Create(&company.Products[i]).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
			return
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}

	log.Println("Created company:", company.Name)
	c.JSON(http.StatusCreated, company)
}

func (h *Handler) GetAllCompanies(c *gin.Context) {
	companies, err := h.Service.GetAllCompanies(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}
	log.Println("Found companies:", companies)
	c.JSON(http.StatusOK, companies)
}

func (h *Handler) GetCompanyById(c *gin.Context) {
	company, err := h.Service.GetCompanyById(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": err.Error()})
		return
	}
	log.Println("Found company by id:", company.Name)
	c.JSON(http.StatusOK, company)
}

func (h *Handler) UpdateCompany(c *gin.Context) {
	company, err := bindCompanyFromJSON(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}
	if err := h.Service.UpdateCompany(c.Request.Context(), c.Param("id"), &company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}
	log.Println("Updated company:", company.Name)
	c.JSON(http.StatusOK, company)
}

func (h *Handler) DeleteCompany(c *gin.Context) {
	company, err := h.Service.DeleteCompany(c.Request.Context(), c.Param("id"))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"errorMessage": "No company found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		}
		return
	}
	log.Println("Deleted company:", company.Name)
	c.JSON(http.StatusOK, company)
}

func bindCompanyFromJSON(c *gin.Context) (Company, error) {
	company := Company{}
	err := c.ShouldBindJSON(&company)
	return company, err
}
