package model

import (
	"time"
)

type Cart struct {
	ID        uint       `json:"id"`
	UserId    uint       `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Products  []CartItem `gorm:"foreignKey:CartId;constraint:OnDelete:CASCADE;"`
}

type CartItem struct {
	ID        uint    `gorm:"primaryKey"`
	CartId    uint    `json:"cart_id" gorm:"index"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
}
