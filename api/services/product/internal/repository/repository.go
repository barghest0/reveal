package repository

import (
	"fmt"
	"product-service/internal/model"
	"time"

	"github.com/barghest0/reveal/api/packages/cache"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) error
	Get(ids []int) (*[]model.Product, error)
	GetByID(id uint) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id uint) error
	Purchase(id uint, buyerID uint) error
}

type productRepository struct {
	db    *gorm.DB
	cache cache.CacheService
}

<<<<<<< HEAD
func CreateProductRepository(db *gorm.DB, cache packages.CacheService) ProductRepository {
=======
func CreateProductRepository(db *gorm.DB, cache cache.CacheService) ProductRepository {
>>>>>>> cart-ui
	return &productRepository{db, cache}
}

func (r *productRepository) Create(product *model.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("product:%d", product.ID)
	return r.cache.Set(cacheKey, product, time.Hour)
}

func (r *productRepository) Get(ids []int) (*[]model.Product, error) {
	var products []model.Product

	if len(ids) > 0 {
		if err := r.db.Where("id IN ?", ids).Find(&products).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.db.Find(&products).Error; err != nil {
			return nil, err
		}
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
