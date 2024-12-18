package service

import (
	"encoding/json"
	"log"
	"product-service/internal/messaging"
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
	rmq  messaging.PublisherManager
}

func CreateProductService(repo repository.ProductRepository, rmq messaging.PublisherManager) OrderService {
	return &productService{repo, rmq}
}

func (s *productService) GetOrders(ids []int) (*[]model.Order, error) {
	return s.repo.Get(ids)
}

func (s *productService) CreateOrder(order *model.Order) error {
	// Создание заказа в репозитории
	err := s.repo.Create(order)
	if err != nil {
		return err
	}

	// Создание события и его публикация в RabbitMQ
	event := map[string]interface{}{
		"message": "Order created",
		"title":   "Order notification",
		"token":   "eJNoPPBmTAiyAGUJtGTmbw:APA91bG2B3eBJQyLQnrzQbUX7Js1kFl5qGrb6u4ufr39wKdnoafrIgRjD7mMzweSZ5bSK3tHZYAKjJ8OAszmEwaOjLZVy-R8VQOBZsdjX83jUin6Tx2XPgs",
	}

	eventBody, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal event: %s", err)
		return err
	}

	// Публикация события в RabbitMQ
	err = s.rmq.Publish("notifications", "", eventBody)
	if err != nil {
		log.Printf("Failed to publish message: %s", err)
		return err
	}

	log.Printf("Published message to RabbitMQ: %s", event["message"])
	return nil
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
