package service

import (
	"product-service/internal/model"
	"product-service/internal/repository"
)

type OrderService interface {
	CreateOrder(product *model.Order) error
	GetOrder(id uint) (*model.Order, error)
	GetOrders(ids []int) (*[]model.Order, error)
	UpdateOrder(product *model.Order) error
	DeleteOrder(id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func CreateProductService(repo repository.ProductRepository) OrderService {
	return &productService{repo}
}

func (s *productService) GetOrders(ids []int) (*[]model.Order, error) {
	return s.repo.Get(ids)
}

func (s *productService) CreateOrder(product *model.Order) error {
	return s.repo.Create(product)
}

func (s *productService) GetOrder(id uint) (*model.Order, error) {
	return s.repo.GetByID(id)
}

func (s *productService) UpdateOrder(product *model.Order) error {
	return s.repo.Update(product)
}

func (s *productService) DeleteOrder(id uint) error {
	return s.repo.Delete(id)
}
