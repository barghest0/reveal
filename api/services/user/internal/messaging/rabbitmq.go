package messaging

import (
	"log"

	"github.com/streadway/amqp"
)

type PublisherManager struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func CreatePublisherManager(url string) (*PublisherManager, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &PublisherManager{
		conn:    conn,
		channel: ch,
	}, nil
}

func (m *PublisherManager) DeclareExchange(name, exchangeType string) error {
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

func (m *PublisherManager) Publish(exchange, routingKey string, body []byte) error {
	err := m.channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %s", err)
		return err
	}
	log.Printf("Message published to exchange '%s' with routing key '%s'", exchange, routingKey)
	return nil
}

func (m *PublisherManager) Close() {
	m.channel.Close()
	m.conn.Close()
}
