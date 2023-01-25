package main

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	_, err = ch.QueueDeclare("test-queue", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind("test-queue", "", "test-exchange", false, nil)
	if err != nil {
		log.Fatal(err)
	}

	d, err := ch.Consume("test-queue", "test", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	for v := range d {
		fmt.Println(string(v.Body))
	}
}
