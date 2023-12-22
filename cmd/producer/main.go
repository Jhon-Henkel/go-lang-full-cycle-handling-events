package main

import "github.com/Jhon-Henkel/go-lang-full-cycle-handling-events/pkg/rabbitmq"

func main() {
	rabbitChannel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer rabbitChannel.Close()
	rabbitmq.Publish(rabbitChannel, "Hello World from producer package!", "amq.direct")
}
