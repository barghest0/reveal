package service

import (
	"encoding/json"
	"log"
	"product-service/internal/model"
)

type NotificationService struct{}

func CreateNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) HandleNotification(body []byte) {
	var notification model.Notification

	log.Printf("Received notification: %s", json.Unmarshal(body, &notification))
}
