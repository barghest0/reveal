package repository

import (
	"cart-service/internal/model"

	"gorm.io/gorm"
)

type CartRepository interface {
	Create(cart *model.Cart) error
	GetByID(userId uint) (*model.Cart, error)
	AddItemToCart(cartId uint, product *model.CartProduct) error
	UpdateCart(cart *model.Cart) error
	RemoveItemFromCart(cartId uint, product_id uint) error
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

func (r *cartRepository) AddItemToCart(cartId uint, product *model.CartProduct) error {
	product.CartId = cartId
	return r.db.Save(product).Error
}

func (r *cartRepository) UpdateCart(cart *model.Cart) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(cart).Error
}

func (r *cartRepository) RemoveItemFromCart(cartId uint, product_id uint) error {
	if err := r.db.Where("cart_id = ? AND product_id = ?", cartId, product_id).Delete(&model.CartProduct{}).Error; err != nil {
		return err
	}

	return nil
}
