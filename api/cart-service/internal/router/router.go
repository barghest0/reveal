package router

import (
	"cart-service/internal/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func InitRoutes(router *gin.Engine, db *gorm.DB) {
	cartHandler := &handler.CartHandler{DB: db}

	router.GET("/cart/:userId/items", cartHandler.GetCartItems)
}
