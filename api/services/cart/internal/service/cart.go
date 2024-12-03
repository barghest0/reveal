package service

import (
	"cart-service/internal/messaging"
	"cart-service/internal/model"
	"cart-service/internal/repository"
)

type CartService interface {
	CreateCart(userId uint) error
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

// func (s *cartService) CreateCartAfterRegistration() {
// 	messages, err := s.rmq.Consume("user.registered")
// 	if err != nil {
// 		log.Fatalf("Failed to consume messages: %v", err)
// 	}
//
// 	for msg := range messages {
// 		var userData Data
//
// 		receivedStr := string(msg.Body)
//
// 		cleanedStr := strings.Trim(receivedStr, "[]")
//
// 		parts := strings.Split(cleanedStr, " ")
//
// 		byteArray := make([]byte, len(parts))
//
// 		for i, part := range parts {
// 			num, err := strconv.Atoi(part)
// 			if err != nil {
// 				fmt.Println("Error converting:", err)
// 				return
// 			}
// 			byteArray[i] = byte(num)
// 		}
//
// 		log.Println(byteArray, &userData, json.Unmarshal(byteArray, &userData))
// 		log.Printf("CONSUME MESSAGE: %s", msg.Body)
//
// 		// Декодируем сообщение
// 		if err := json.Unmarshal(byteArray, &userData); err != nil {
// 			log.Printf("Failed to unmarshal user_id: %v", err)
// 			continue
// 		}
//
// 		log.Println(userData.Id, "CONSUME USER DATA")
//
// 		// Создаем корзину для пользователя
// 		cart := &model.Cart{UserId: userData.Id}
// 		if err := s.repo.Create(cart); err != nil {
// 			log.Printf("Failed to create cart for user %v: %v", userData.Id, err)
// 		} else {
// 			log.Printf("Created cart for user %v", userData.Id)
// 		}
// 	}
// }
