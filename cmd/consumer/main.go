package main

import (
	"fmt"
	"github.com/Jhon-Henkel/go-lang-full-cycle-handling-events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitChannel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer rabbitChannel.Close()
	messagesChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(rabbitChannel, messagesChannel)
	for message := range messagesChannel {
		fmt.Println(string(message.Body))
		message.Ack(false)
	}
}
