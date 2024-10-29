package repository

import (
	"cart-service/internal/model"

	"gorm.io/gorm"
)

type CartRepository interface {
	Create(cart *model.Cart) error
	GetByID(userId uint) (*model.Cart, error)
	AddItemToCart(cartId uint, product *model.CartItem) error
	UpdateCart(cart *model.Cart) error
	RemoveItemToCart(cartId uint, product_id uint) error
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
	if err := r.db.Preload("Products").Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil

}

func (r *cartRepository) AddItemToCart(cartId uint, product *model.CartItem) error {
	product.CartId = cartId
	return r.db.Save(product).Error
}

func (r *cartRepository) UpdateCart(cart *model.Cart) error {
	return r.db.Save(cart).Error
}

func (r *cartRepository) RemoveItemToCart(cartId uint, product_id uint) error {
	if err := r.db.Where("id = ? AND cart_id = ?", product_id, cartId).Delete(&model.CartItem{}).Error; err != nil {
		return err
	}
	return nil
}
