package main

import (
	"log"
	"product-service/internal/config"
	"product-service/internal/messaging"
	"product-service/internal/service"
)

func main() {
	appConfig := config.LoadConfig()

	// Подключаемся к RabbitMQ
	rabbitMQManager, err := messaging.CreateConsumerManager(appConfig.BrokerURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitMQManager.Close()

	// Объявляем обменник и очередь
	err = rabbitMQManager.DeclareExchange(appConfig.Exchange, "fanout")
	if err != nil {
		log.Fatalf("Failed to declare exchange: %v", err)
	}

	err = rabbitMQManager.DeclareAndBindQueue(appConfig.OrderQueue, "", appConfig.Exchange)
	if err != nil {
		log.Fatalf("Failed to declare or bind queue: %v", err)
	}

	// Инициализируем сервис уведомлений
	notificationService := service.CreateNotificationService()

	// Начинаем обрабатывать сообщения
	go messaging.StartConsumer(rabbitMQManager, appConfig.OrderQueue, notificationService)

	// Логируем успешный запуск
	log.Printf("Notification service is running on RabbitMQ exchange: '%s', queue: '%s'", appConfig.Exchange, appConfig.OrderQueue)

	// Блокируем выполнение основного потока
	select {}
}
