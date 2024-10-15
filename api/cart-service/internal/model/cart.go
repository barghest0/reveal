package model

import "time"

type Cart struct {
	ID        int        `json:"id"`
	UserId    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	Items     []CartItem `gorm:"foreignKey:CartID"`
}

type CartItem struct {
	ID        int     `json:"id"`
	CartId    int     `json:"cart_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
