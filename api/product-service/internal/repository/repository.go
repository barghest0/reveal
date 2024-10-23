package repository

import (
	"fmt"
	"product-service/internal/model"
	"time"

	"gorm.io/gorm"
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
	db    *gorm.DB
	cache packages.CacheService
}

func CreateProductRepository(db *gorm.DB, cache packages.CacheService) ProductRepository {
	return &productRepository{db, cache}
}

func (r *productRepository) Create(product *model.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("product:%d", product.ID)
	return r.cache.Set(cacheKey, product, time.Hour)
}

func (r *productRepository) GetAll() (*[]model.Product, error) {
	cacheKey := "products:all"

	var products []model.Product
	if err := r.cache.Get(cacheKey, &products); err == nil {
		return &products, nil
	}

	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	if err := r.cache.Set(cacheKey, &products, time.Hour); err != nil {
		return nil, err
	}

	return &products, nil
}

func (r *productRepository) GetByID(id uint) (*model.Product, error) {
	cacheKey := fmt.Sprintf("product:%d", id)

	var product model.Product
	if err := r.cache.Get(cacheKey, &product); err == nil {
		return &product, nil
	}

	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	if err := r.cache.Set(cacheKey, &product, time.Hour); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Update(product *model.Product) error {
	if err := r.db.Save(product).Error; err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("product:%d", product.ID)
	return r.cache.Set(cacheKey, product, time.Hour)
}

func (r *productRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.Product{}, id).Error; err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("product:%d", id)
	return r.cache.Delete(cacheKey)
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
