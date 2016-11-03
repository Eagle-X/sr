package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		//log.Printf("%s: %s, %t", msg, err, err == amqp.ErrClosed)
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	if lb {
		ListBrokers()
		return
	}

	if bid < 0 || bid >= len(brokers) {
		log.Fatalf("Invalid broker id: %d", bid)
	}
	conn, err := amqp.Dial(brokers[bid])
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	if isRecv {
		recv(ch)
	} else {
		send(ch)
	}
}
