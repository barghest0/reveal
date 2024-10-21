package service

import (
	"cart-service/internal/model"
	"cart-service/internal/repository"
)

type CartService interface {
	CreateCart(cart *model.Cart) error
	GetCart(userId uint) (*model.Cart, error)
	AddItem(cartId uint, item *model.CartItem) error
}

type cartService struct {
	repo repository.CartRepository
}

func CreateCartService(repo repository.CartRepository) CartService {
	return &cartService{repo}
}

func (s *cartService) CreateCart(cart *model.Cart) error {
	return s.repo.Create(cart)
}

func (s *cartService) GetCart(userId uint) (*model.Cart, error) {
	return s.repo.GetByID(userId)
}

func (s *cartService) AddItem(cartId uint, item *model.CartItem) error {
	return s.repo.AddItemToCart(cartId, item)
}
