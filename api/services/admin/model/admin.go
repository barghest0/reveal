package model

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"name" gorm: "not null"`
	Email    string `json:"email" gorm:"unique;not null`
	Password string `json:"password" gorm:"not null`
	Role     string `gorm:"default:'admin'"`
}

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

type CartProduct struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	CartId    uint    `json:"cart_id" gorm:"index"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
	Product   Product `gorm:"-" json:"product,omitempty"`
}

type Cart struct {
	ID        uint          `json:"id"`
	UserId    uint          `json:"user_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Products  []CartProduct `json:"products" gorm:"foreignKey:CartId;constraint:OnDelete:CASCADE;"`
}
