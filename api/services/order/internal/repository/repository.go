package repository

import (
	"fmt"
	"product-service/internal/model"
	"time"

	"github.com/barghest0/reveal/api/packages/cache"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Order) error
	Get(ids []int) (*[]model.Order, error)
	GetByID(id uint) (*model.Order, error)
	Update(product *model.Order) error
	Delete(id uint) error
}

type productRepository struct {
	db    *gorm.DB
	cache cache.CacheService
}

func CreateOrderRepository(db *gorm.DB, cache cache.CacheService) ProductRepository {
	return &productRepository{db, cache}
}

func (r *productRepository) Create(product *model.Order) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("order:%d", product.ID)
	return r.cache.Set(cacheKey, product, time.Hour)
}

func (r *productRepository) Get(ids []int) (*[]model.Order, error) {
	var products []model.Order

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

func (r *productRepository) GetByID(id uint) (*model.Order, error) {
	cacheKey := fmt.Sprintf("order:%d", id)

	var product model.Order
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

func (r *productRepository) Update(product *model.Order) error {
	if err := r.db.Save(product).Error; err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("order:%d", product.ID)
	return r.cache.Set(cacheKey, product, time.Hour)
}

func (r *productRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.Order{}, id).Error; err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("order:%d", id)
	return r.cache.Delete(cacheKey)
}
