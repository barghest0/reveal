package main

import (
	"log"
	"product-service/internal/config"
	"product-service/internal/handler"
	"product-service/internal/messaging"
	"product-service/internal/service"
)

func main() {
	appConfig := config.LoadConfig()

	// Подключаемся к RabbitMQ
	rmq, err := messaging.CreateConsumerManager(appConfig.BrokerURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rmq.Close()

	// Объявляем обменник и очередь
	err = rmq.DeclareExchange(appConfig.Exchange, "fanout")
	if err != nil {
		log.Fatalf("Failed to declare exchange: %v", err)
	}

	// Инициализируем сервис уведомлений
	service := service.CreateNotificationService()

	handler.RegisterSubscribers(rmq, service)

	// Логируем успешный запуск
	log.Printf("Notification service is running on RabbitMQ exchange: '%s', queue: '%s'", appConfig.Exchange, appConfig.OrderQueue)

	// Блокируем выполнение основного потока
	select {}
}
