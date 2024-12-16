package service

import (
	"encoding/json"
	"log"
	"product-service/internal/model"

	"github.com/streadway/amqp"
)

type NotificationService struct{}

func CreateNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) HandleNotification(d amqp.Delivery) {
	var notification model.Notification

	err := json.Unmarshal(d.Body, &notification)
	if err != nil {
		log.Printf("Error decoding event: %s", err)
		return
	}

	log.Printf("Received notification: %s", notification.Message)
}
