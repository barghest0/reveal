package messaging

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewRabbitMQ создает новое подключение и канал RabbitMQ
func CreateRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{conn, ch}, nil
}

// Consume подключается к указанной очереди и возвращает канал для чтения сообщений
func (r *RabbitMQ) Consume(queueName string) (<-chan amqp.Delivery, error) {
	// Декларируем очередь, если она еще не существует
	q, err := r.channel.QueueDeclare(
		queueName, // имя очереди
		false,     // не сохраняем очередь при рестарте сервера
		false,     // не удаляем очередь, если нет потребителей
		false,     // очередь доступна только для текущего соединения
		false,     // без ожидания
		nil,       // дополнительные аргументы
	)
	if err != nil {
		return nil, err
	}

	// Подключаемся к очереди для потребления сообщений
	messages, err := r.channel.Consume(
		q.Name, // имя очереди
		"",     // потребитель
		true,   // авто-подтверждение
		false,  // эксклюзивное потребление
		false,  // отсутствие локального режима
		false,  // без ожидания
		nil,    // дополнительные аргументы
	)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// Close закрывает соединение и канал RabbitMQ
func (r *RabbitMQ) Close() {
	if err := r.channel.Close(); err != nil {
		log.Printf("Error closing channel: %v", err)
	}
	if err := r.conn.Close(); err != nil {
		log.Printf("Error closing connection: %v", err)
	}
}
