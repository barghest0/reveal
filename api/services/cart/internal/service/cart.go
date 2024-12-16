package service

import (
	"cart-service/internal/messaging"
	"cart-service/internal/model"
	"cart-service/internal/repository"
)

type CartService interface {
	CreateCart(userId uint) error
	GetCarts(ids []int) ([]model.Cart, error)
	GetCart(userId uint) (*model.Cart, error)
	AddProductToCart(cartId uint, product *model.CartProduct) error
	UpdateCart(cart *model.Cart) error
	RemoveProductFromCart(cartId uint, product_id uint) error
}

type cartService struct {
	repo repository.CartRepository
	rmq  messaging.ConsumerManager
}

func CreateCartService(repo repository.CartRepository, rmq *messaging.ConsumerManager) CartService {
	return &cartService{repo, *rmq}
}

func (s *cartService) CreateCart(userId uint) error {
	existingCart, err := s.repo.GetByID(userId)
	if err == nil && existingCart != nil {
		return nil
	}

	cart := &model.Cart{
		UserId: userId,
	}

	return s.repo.Create(cart)
}

func (s *cartService) GetCarts(ids []int) ([]model.Cart, error) {
	return s.repo.Get(ids)
}

func (s *cartService) GetCart(userId uint) (*model.Cart, error) {
	return s.repo.GetByID(userId)
}

func (s *cartService) AddProductToCart(cartId uint, product *model.CartProduct) error {
	return s.repo.AddItemToCart(cartId, product)
}

func (s *cartService) UpdateCart(cart *model.Cart) error {
	return s.repo.UpdateCart(cart)
}

func (s *cartService) RemoveProductFromCart(cartId uint, product_id uint) error {
	return s.repo.RemoveItemFromCart(cartId, product_id)
}

type Data struct {
	Id uint `json:"id"`
}
