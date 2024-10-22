package repository

import (
	"cart-service/internal/model"

	"gorm.io/gorm"
)

type CartRepository interface {
	Create(cart *model.Cart) error
	GetByID(userId uint) (*model.Cart, error)
	AddItemToCart(cartId uint, item *model.CartItem) error
	UpdateCart(cart *model.Cart) error
}

type cartRepository struct {
	db *gorm.DB
}

func CreateCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) Create(cart *model.Cart) error {
	return r.db.Create(cart).Error
}

func (r *cartRepository) GetByID(userId uint) (*model.Cart, error) {
	var cart model.Cart
	// Используем Preload для загрузки связанных элементов
	if err := r.db.Preload("Items").Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil

}

func (r *cartRepository) AddItemToCart(cartId uint, item *model.CartItem) error {
	item.CartId = cartId
	return r.db.Save(item).Error
}

func (r *cartRepository) UpdateCart(cart *model.Cart) error {
	return r.db.Save(cart).Error
}
