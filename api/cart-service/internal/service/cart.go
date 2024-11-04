package service

import (
	"cart-service/internal/model"
	messaging "cart-service/internal/rabbitmq"
	"cart-service/internal/repository"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type CartService interface {
	CreateCart(cart *model.Cart) error
	GetCart(userId uint) (*model.Cart, error)
	AddItem(cartId uint, product *model.CartItem) error
	UpdateCart(cart *model.Cart) error
	RemoveItemToCart(cartId uint, product_id uint) error
	UserRegistrationListener()
}

type cartService struct {
	repo repository.CartRepository
	rmq  messaging.RabbitMQ
}

func CreateCartService(repo repository.CartRepository, rmq *messaging.RabbitMQ) CartService {
	return &cartService{repo, *rmq}
}

type Data struct {
	Id uint `json:"id"`
}

// Прослушивание очереди и создание корзины для зарегистрированных пользователей
func (s *cartService) UserRegistrationListener() {
	messages, err := s.rmq.Consume("user.registered")
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	for msg := range messages {
		var userData Data

		receivedStr := string(msg.Body)

		// Удаляем квадратные скобки
		cleanedStr := strings.Trim(receivedStr, "[]")

		// Разделяем строку на части
		parts := strings.Split(cleanedStr, " ")

		// Создаем массив байтов
		byteArray := make([]byte, len(parts))

		// Логируем полученное сообщение

		for i, part := range parts {
			num, err := strconv.Atoi(part) // Преобразование строки в целое число
			if err != nil {
				fmt.Println("Error converting:", err)
				return
			}
			byteArray[i] = byte(num) // Преобразование числа в байт
		}

		log.Println(byteArray, &userData, json.Unmarshal(byteArray, &userData))
		log.Printf("CONSUME MESSAGE: %s", msg.Body)

		// Декодируем сообщение
		if err := json.Unmarshal(byteArray, &userData); err != nil {
			log.Println("Failed to unmarshal user_id: %v", err)
			continue
		}

		log.Println(userData.Id, "CONSUME USER DATA")

		// Создаем корзину для пользователя
		cart := &model.Cart{UserId: userData.Id}
		if err := s.repo.Create(cart); err != nil {
			log.Printf("Failed to create cart for user %v: %v", userData.Id, err)
		} else {
			log.Printf("Created cart for user %v", userData.Id)
		}
	}
}

func (s *cartService) CreateCart(cart *model.Cart) error {
	return s.repo.Create(cart)
}

func (s *cartService) GetCart(userId uint) (*model.Cart, error) {
	return s.repo.GetByID(userId)
}

func (s *cartService) AddItem(cartId uint, product *model.CartItem) error {
	return s.repo.AddItemToCart(cartId, product)
}

func (s *cartService) UpdateCart(cart *model.Cart) error {
	return s.repo.UpdateCart(cart)
}

func (s *cartService) RemoveItemToCart(cartId uint, product_id uint) error {
	return s.repo.RemoveItemToCart(cartId, product_id)
}
