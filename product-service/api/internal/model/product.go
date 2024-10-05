package model

import "time"

type Product struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	SellerID    uint       `json:"seller_id"`
	BuyerID     *uint      `json:"buyer_id,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
	SoldAt      *time.Time `json:"sold_at,omitempty"`
}
