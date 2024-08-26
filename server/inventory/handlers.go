package inventory

import (
	"github.com/gin-gonic/gin"
	"gorsk/server/product"
	"net/http"
)

type Handler struct {
	Service        *Service
	ProductHandler *product.Handler
}

func NewInventoryHandler(service *Service, productHandler *product.Handler) *Handler {
	return &Handler{Service: service, ProductHandler: productHandler}
}

func (h *Handler) AddInventories(c *gin.Context) {
	inventory := &Inventory{}
	if err := c.ShouldBindJSON(inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.ProductHandler.Service.GetProductById(c.Request.Context(), c.Param("product.id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	if product == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}

	if err := h.Service.CreateInventory(c.Request.Context(), inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
