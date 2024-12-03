package messaging

import (
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

func (m *ConsumerManager) DeclareQueue(queueName string) (amqp.Queue, error) {
	return m.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
}

func (m *ConsumerManager) BindQueue(queueName, routingKey, exchangeName string) error {
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
