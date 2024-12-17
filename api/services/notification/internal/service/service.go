package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"product-service/internal/model"

	"github.com/streadway/amqp"
)

type NotificationService struct {
	token string
}

func CreateNotificationService(token string) *NotificationService {
	return &NotificationService{token}
}

func (s *NotificationService) HandleNotification(d amqp.Delivery) error {
	var notification model.Notification

	err := json.Unmarshal(d.Body, &notification)
	// Формируем URL для API запроса
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", s.token)

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
