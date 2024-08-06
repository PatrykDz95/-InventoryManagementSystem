package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Handler struct {
	Service *Service
}

func NewProductHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) AddProducts(c *gin.Context) {
	var products []Product
	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}

	for _, product := range products {
		if err := h.Service.CreateProduct(c.Request.Context(), &product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
			return
		}
		log.Println("Created product:", product.Name)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Products created successfully"})
}

func (h *Handler) GetAllProducts(c *gin.Context) {
	products, err := h.Service.GetAllProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}
	log.Println("Found products:", products)
	c.JSON(http.StatusOK, products)
}

func (h *Handler) GetProductById(c *gin.Context) {
	product, err := h.Service.GetProductById(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": err.Error()})
		return
	}
	log.Println("Found product by id:", product.Name)
	c.JSON(http.StatusOK, product)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	product, err := bindProductFromJSON(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}
	if err := h.Service.UpdateProduct(c.Request.Context(), c.Param("id"), &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}
	log.Println("Updated product:", product.Name)
	c.JSON(http.StatusOK, product)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	product, err := h.Service.DeleteProduct(c.Request.Context(), c.Param("id"))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"errorMessage": "No product found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		}
		return
	}
	log.Println("Deleted product:", product.Name)
	c.JSON(http.StatusOK, product)
}

func bindProductFromJSON(c *gin.Context) (Product, error) {
	product := Product{}
	err := c.ShouldBindJSON(&product)
	return product, err
}
