package handler

import (
	"cart-service/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *CartHandler) CreateCart(c *gin.Context) {
	var cart model.Cart

	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, cart)
}
