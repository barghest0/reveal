package handler

import (
	"cart-service/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *CartHandler) AddItemToCart(c *gin.Context) {
	var cartItem model.CartItem

	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, cartItem)
}
