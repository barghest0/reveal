package repository

import (
	"gorm.io/gorm"
	"product-service/internal/model"
	"time"
)

type ProductRepository interface {
	Create(product *model.Product) error
	GetAll() (*[]model.Product, error)
	GetByID(id uint) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id uint) error
	Purchase(id uint, buyerID uint) error
}

type productRepository struct {
	db *gorm.DB
}

func CreateProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetAll() (*[]model.Product, error) {
	var products []model.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

func (r *productRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}

func (r *productRepository) Purchase(id uint, buyerID uint) error {
	var product model.Product

	if err := r.db.First(&product, id).Error; err != nil {
		return err
	}

	if product.BuyerID != nil {
		return gorm.ErrInvalidData
	}

	buyer := buyerID
	soldAt := time.Now()

	return r.db.Model(&product).Updates(map[string]interface{}{
		"buyer_id": buyer,
		"sold_at":  soldAt,
	}).Error
}
