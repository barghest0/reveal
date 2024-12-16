package handler

import (
	"log"
	"product-service/internal/messaging"
	"product-service/internal/service"
)

func RegisterSubscribers(consumerManager *messaging.ConsumerManager, service *service.NotificationService) {
	// Объявление обменника
	err := consumerManager.DeclareExchange("notifications", "fanout")
	if err != nil {
		log.Fatalf("Failed to declare exchange: %s", err)
	}

	queue, err := consumerManager.DeclareQueue("orders")
	if err != nil {
		log.Fatalf("Failed to declare queue: %s", err)
	}

	// Привязка очереди к обменнику
	err = consumerManager.BindQueue(queue.Name, "", "notifications")
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
			go service.HandleNotification(msg)
		}
	}()
}
