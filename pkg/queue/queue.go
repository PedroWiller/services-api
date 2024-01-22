package queue

import (
	"encoding/json"

	"github.com/streadway/amqp"

	"face-id-detection/config/env"
)

var (
	connectRabbitMQ *amqp.Connection
	channelRabbitMQ *amqp.Channel
)

func Init() error {
	var err error

	connectRabbitMQ, err = amqp.Dial(env.AmqpServerURL)
	if err != nil {
		return err
	}

	channelRabbitMQ, err = connectRabbitMQ.Channel()
	if err != nil {
		return err
	}
	_, err = channelRabbitMQ.QueueDeclare(
		env.AmqpQueueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

func Publish(body interface{}) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        jsonBody,
	}

	if err := channelRabbitMQ.Publish(
		"",
		env.AmqpQueueName,
		false,
		false,
		message,
	); err != nil {
		return err
	}

	return nil
}

func Close() {
	connectRabbitMQ.Close()
	channelRabbitMQ.Close()
}
