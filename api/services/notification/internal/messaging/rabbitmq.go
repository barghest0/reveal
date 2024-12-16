package messaging

import (
	"log"
	"product-service/internal/service"

	"github.com/streadway/amqp"
)

type ConsumerManager struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func CreateConsumerManager(url string) (*ConsumerManager, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &ConsumerManager{
		conn:    conn,
		channel: ch,
	}, nil
}

func (m *ConsumerManager) DeclareExchange(name, exchangeType string) error {
	return m.channel.ExchangeDeclare(
		name,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
}

func (m *ConsumerManager) DeclareAndBindQueue(queueName, routingKey, exchangeName string) error {
	_, err := m.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return m.channel.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false,
		nil,
	)
}

func (m *ConsumerManager) Consume(queueName string) (<-chan amqp.Delivery, error) {
	return m.channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
}

func (m *ConsumerManager) Close() {
	m.channel.Close()
	m.conn.Close()
}

func StartConsumer(m *ConsumerManager, queueName string, service *service.NotificationService) {
	msgs, err := m.Consume(
		queueName,
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	for msg := range msgs {
		service.HandleNotification(msg.Body)
	}
}
