package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	rabbitChannel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return rabbitChannel, nil
}

func Consume(rabbitChannel *amqp.Channel, output chan<- amqp.Delivery, queue string) error {
	messages, err := rabbitChannel.Consume(
		queue,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for message := range messages {
		output <- message
	}
	return nil
}

func Publish(rabbitChannel *amqp.Channel, body string, exchangeName string) error {
	err := rabbitChannel.Publish(
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
