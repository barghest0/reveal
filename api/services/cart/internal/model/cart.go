package model

import (
	"time"
)

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Cart struct {
	ID        uint          `json:"id"`
	UserId    uint          `json:"user_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Products  []CartProduct `json:"products" gorm:"foreignKey:CartId;constraint:OnDelete:CASCADE;"`
}

type CartProduct struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	CartId    uint    `json:"cart_id" gorm:"index"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
	Product   Product `gorm:"-" json:"product,omitempty"`
}
