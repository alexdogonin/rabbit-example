package main

import (
	"fmt"
	"log"
	"time"

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

	err = ch.ExchangeDeclare("test-exchange", amqp091.ExchangeTopic, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; ; i++ {
		time.Sleep(time.Second)

		err = ch.Publish("test-exchange", "", false, false, amqp091.Publishing{
			Body: []byte(fmt.Sprintf("hello %d", i)),
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}
