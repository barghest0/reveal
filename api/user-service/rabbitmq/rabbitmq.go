package rabbitmq

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Channel *amqp.Channel
}

func CreateRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{Channel: channel}, nil
}

func (r *RabbitMQ) Publish(queueName string, body []byte) error {
	_, err := r.Channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = r.Channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	return err
}
