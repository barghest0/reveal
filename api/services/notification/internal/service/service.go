package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"product-service/internal/model"

	"firebase.google.com/go/v4/messaging"
	"github.com/appleboy/go-fcm"
	"github.com/streadway/amqp"
)

type NotificationService struct {
}

func CreateNotificationService() *NotificationService {
	return &NotificationService{}
}

const TelegramToken string = ""

// FIXME: move api request to handlers
func (s *NotificationService) HandleTelegramNotification(d amqp.Delivery) error {
	var notification model.Notification

	err := json.Unmarshal(d.Body, &notification)
	// Формируем URL для API запроса
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", TelegramToken)

	// Тело запроса в формате JSON
	jsonStr := []byte(fmt.Sprintf(`{"chat_id":"%s","text":"%s"}`, "664719998", notification.Message))

	// Отправка POST-запроса
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("не удалось отправить сообщение, статус: %s", resp.Status)
	}

	fmt.Println("Сообщение успешно отправлено!")
	return nil
}

// FIXME: move to handlers
func (s *NotificationService) HandlePushNotification(d amqp.Delivery) {
	// Пример декодирования сообщения (должно быть в формате JSON)
	var notificationData struct {
		Token   string `json:"token"`   // Токен устройства
		Title   string `json:"title"`   // Токен устройства
		Message string `json:"message"` // Текст уведомления
	}

	if err := json.Unmarshal(d.Body, &notificationData); err != nil {
		log.Printf("Failed to decode message: %v", err)
		return
	}

	log.Printf("NotificationData: %v", notificationData)

	// Отправка уведомления
	err := s.sendPush(notificationData.Token, notificationData.Title, notificationData.Message)
	if err != nil {
		log.Printf("Failed to send push notification: %v", err)
	}
}

func (s *NotificationService) sendPush(token, title, message string) error {
	// Инициализация Firebase приложения
	credentials := []byte("")

	client, err := fcm.NewClient(context.Background(), fcm.WithCredentialsJSON(credentials))

	if err != nil {
		log.Printf("Failed to init client: %v", err)
		return err
	}

	notification := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  message,
		},
	}

	// Отправка запроса в FCM
	response, err := client.Send(context.Background(), notification)
	if err != nil {
		log.Printf("Failed to send: %v", err)
		return err
	}

	log.Printf("Уведомление отправлено успешно: %v", response)

	return nil
}
