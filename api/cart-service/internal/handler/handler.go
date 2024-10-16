package handler

import (
	"cart-service/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartHandler struct {
	DB *gorm.DB
}

func (ctrl *CartHandler) GetCartItems(c *gin.Context) {
	userId := c.Param("userId")

	var cart model.Cart

	if err := ctrl.DB.Preload("Items").Where("user_id = ?", userId).First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, cart.Items)
}
