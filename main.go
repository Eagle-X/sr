package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var (
	Ex01 = "logs"
	Ex02 = "logs-internal"
	Q01  = "fff"
)

func failOnError(err error, msg string) {
	if err != nil {
		//log.Printf("%s: %s, %t", msg, err, err == amqp.ErrClosed)
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func declare(ch *amqp.Channel) {
	err := ch.ExchangeDeclare(
		Ex01,     // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	err = ch.ExchangeDeclare(
		Ex02,     // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	err = ch.ExchangeBind(
		Ex02, // queue name
		"",   // routing key
		Ex01, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	args := make(amqp.Table)
	args["x-expires"] = int32(10000)
	//args["x-message-ttl"] = int32(-1)
	//args["x-max-priority"] = int32(64)

	q, err := ch.QueueDeclare(
		Q01,   // name
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		Ex02,   // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")
}

func main() {
	if lb {
		ListBrokers()
		return
	}

	if bid < 0 || bid >= len(brokers) {
		log.Fatalf("Invalid broker id: %d", bid)
	}
	bk := brokers[bid]
	if len(broker) > 0 {
		bk = broker
	}
	log.Printf("Connect to %s", bk)
	conn, err := amqp.Dial(bk)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	declare(ch)

	if isRecv {
		recv(ch)
	} else {
		send(ch)
	}
}
