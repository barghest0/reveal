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

const TelegramToken string = "7782679125:AAG33kQ2c7DX0AfUAyvgMFRjWI2mIuzKy7c"

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
	credentials := []byte(`{
		"type": "service_account",
		"project_id": "reveal-ec60d",
		"private_key_id": "f72bdc7ce18406853e2296567a11a1cbc870644e",
		"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDrzGur3AeUszlN\nAsHunDvdbcESl4t/RyUxS8C1cnB7Oho9Ple2wG8ETC+PXA6S1RHiboEBj98W3tv4\nEWrJCfSvsQUFAJEvQhKiAkmaIAI8nQWif8uOMG0h68nSdE66txr6SSvkrrlioYD1\nDQPW2N4uH/htmCxQ3wgce8MCm3e/8AHk1e4t6tjor8XF1PkcEY+ydDwkuixzzju8\niFrEE+COTfk7abrqKt5WYgaIWjEZwiXdJ2TArM/BwRbZYHmAZN6kkiM7ss8pPZDK\nuwb9Sq9vnXnZU+xqz3ep6XgGAtWP+QbJ3PZpoJl8AKbbKtLFQW4j/bCrxbTXb1D7\nuLOsbVShAgMBAAECggEAA1WGIvyHdLxCqUIsgKtJL3hwauzWQCe+8rEJPKD4zpE/\nJcNGgZx4GhedveSkq5SQiXtnIntMfc6xGSnlS036zoIoEjRJKE7rulHp3z6R02jJ\nYusw/oAnd9JkVxWl1YLTq8+/b1O9qfz10xXo3S3CoLd7GOAUj0Y8DLMr78PW55zS\n16Xwte2eekvhsloWnS4HsIBqzmd+qvuO15YmEzGILW1PExktKpz9KYIz8lUJoWma\ngy97oS1GqFResxS56YiKoo7msjWOd+6hdb4f6/La+nJg5AwzGbpvQBJa0rpZarO7\nBP/zU3wnm+UKuEnJA6v2AML14pJondzraplRHSv6XQKBgQD66a4hGPdhz2/unt4Q\n0CBxMlpDCzBEbcVl6p+z7/QgBLYeC57JKjhjwdX7gQPaeFRqJL5jE101wajFg/KS\n37zIiB8Yx8jAQ1RsltCO5Ci0h4LSRkQ0gAytl2cVvCVOJV84FQkhw6uLCOlPf5fr\nA6NnHuFvkIi6kJnTLy6PyrKJ7QKBgQDwlErQwysJK5pI/xJdNv3Z3o6f0LzzLHUn\nQrgSDqQyoM2HCLz5DCwmOVhh73UP92QDA2EuHnD5J7XhM6gatlLDxzzHYWE2IErF\nvY43veHSkgYHVuSUs5nGUAeM2r+ykm1K7GFeaQ4kSffLq2+jngC1s4vjVusvWWsw\nIdbzadLPBQKBgQCZ7XO6sENLSOPlIqce71+Hshk3mgyRXjpETMyOqb09S0aiZgSS\naolPgaGRelNv2nQG7eWyiDWdfeHY0hqlSgbZitBw9ldMw8FuIJNvpb7nCV8TOrIq\n9wdQE1rVNag5KaohkJ4ajZyWXfArqh37ui6rW4F3XNbUhVMpAH/zDKm8/QKBgQDM\nszDxbo7or5wIblhezhKy1YZ/fym2s2Y/cKqUWrbliSM3uXypX+0U3QhR1GDDpBkJ\n409F4Yr4xQwhwUVQ7T+A677QWxBO8K0OkZo33a4O8EaG6i12RVcbgS7bb6latPKf\npH+4yCTLRdC9EHLeQeuSr2Hgny1fzXLBspll5e4ETQKBgQD3Pyi+kos3t0vtTp/B\nCDUp4gSWPJFFDs8+wOT2ipBinDuWNvvM0JCKuVtBljmlKAcprimANP7600P7VPzZ\nHRgyWQ0sI99Cn078x/fDOtvvrg+J2H9bUc3Aa/v2a18+UbUd/V5SBg9EWWgpwp89\nNT3v2h0+GpzwTRZQc5mczO3jcA==\n-----END PRIVATE KEY-----\n",
		"client_email": "firebase-adminsdk-htff0@reveal-ec60d.iam.gserviceaccount.com",
		"client_id": "109601878392440352685",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-htff0%40reveal-ec60d.iam.gserviceaccount.com",
		"universe_domain": "googleapis.com"
	}`)

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
