package service

import (
	"product-service/internal/model"
	"product-service/internal/repository"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	GetProduct(id uint) (*model.Product, error)
	GetProducts() (*[]model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(id uint) error
	PurchaseProduct(id uint, buyerID uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func (s *productService) GetProducts() (*[]model.Product, error) {
	return s.repo.GetAll()
}

func CreateProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) CreateProduct(product *model.Product) error {
	return s.repo.Create(product)
}

func (s *productService) GetProduct(id uint) (*model.Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) UpdateProduct(product *model.Product) error {
	return s.repo.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}

func (s *productService) PurchaseProduct(id uint, buyerID uint) error {
	return s.repo.Purchase(id, buyerID)
}
