package handler

import (
	"cart-service/internal/messaging"
	"log"
)

func RegisterSubscribers(consumerManager *messaging.ConsumerManager, handler *CartHandler) {
	// Объявление обменника
	err := consumerManager.DeclareExchange("user_events", "fanout")
	if err != nil {
		log.Fatalf("Failed to declare exchange: %s", err)
	}

	// Объявление очереди
	queue, err := consumerManager.DeclareQueue("create_cart_queue")
	if err != nil {
		log.Fatalf("Failed to declare queue: %s", err)
	}

	// Привязка очереди к обменнику
	err = consumerManager.BindQueue(queue.Name, "", "user_events")
	if err != nil {
		log.Fatalf("Failed to bind queue: %s", err)
	}

	msgs, err := consumerManager.Consume(queue.Name)
	if err != nil {
		log.Fatalf("Failed to register consumer: %s", err)
	}

	// Обработка сообщений
	go func() {
		log.Printf("Listening to queue %s...", queue.Name)
		for msg := range msgs {
			go handler.HandleUserCreatedEvent(msg)
		}
	}()
}
