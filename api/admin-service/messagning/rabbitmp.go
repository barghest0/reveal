package messaging

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

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

func (r *RabbitMQ) Publish(queue string, message interface{}) error {
	q, err := r.channel.QueueDeclare(queue, false, false, false, false, nil)
	if err != nil {
		return err
	}

	return r.channel.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(fmt.Sprintf("%v", message)),
	})
}

func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.conn.Close()
}
