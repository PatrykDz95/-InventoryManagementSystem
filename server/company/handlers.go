package company

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorsk/server/category"
	"log"
	"net/http"
	"time"
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

	tx := h.Service.DB.Begin()
	if err := tx.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}

	// Create the company without its products
	company.AddedDate = time.Now()
	products := company.Products
	company.Products = nil
	if err := tx.Create(&company).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"Error while creating a company": err.Error()})
		return
	}

	// Create or validate the categories and products
	for i := range products {
		// Check if the productCategory exists
		var productCategory category.Category
		// TODO change to find by id
		if err := tx.Where("name = ?", products[i].Category.Name).First(&productCategory).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Create the productCategory if it does not exist
				productCategory = products[i].Category
				if err := tx.Create(&productCategory).Error; err != nil {
					tx.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{"Error while creating a productCategory": err.Error()})
					return
				}
			} else {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"Error while searching for a productCategory": err.Error()})
				return
			}
		}

		// Set the CategoryID to the valid productCategory
		products[i].CategoryID = productCategory.ID
		products[i].ID = 0 // Ensure ID is zero to avoid primary key conflict
		products[i].CompanyID = company.ID
		products[i].Category = category.Category{}
		if err := tx.Create(&products[i]).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}

	log.Println("Created company:", company.Name)
	c.JSON(http.StatusCreated, company.Name)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
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
