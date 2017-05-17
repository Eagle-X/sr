package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var (
	Ex   = "logs-internal"
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

func declareBindQueueOnly(ch *amqp.Channel) {
	args := make(amqp.Table)
	args["x-expires"] = int32(10000)

	for i := 0; i < count; i++ {
		q, e := ch.QueueDeclare(
			fmt.Sprintf("%s_%03d", Q01, i), // name
			true,  // durable
			false, // delete when usused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(e, "Failed to declare a queue")
		err := ch.QueueBind(
			q.Name, // queue name
			"",     // routing key
			Ex02,   // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}

}

func deleteQueueOnly(ch *amqp.Channel) {
	for i := 0; i < count; i++ {
		ch.QueueDelete(fmt.Sprintf("%s_%03d", Q01, i), false, false, true)
	}
}

func declare(ch *amqp.Channel) {
	err := ch.ExchangeDeclare(
		Ex,       // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	if !isRecv {
		return
	}

	/*err = ch.ExchangeDeclare(
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
	failOnError(err, "Failed to bind a queue")*/

	args := make(amqp.Table)
	//args["x-expires"] = int32(10000)
	args["x-queue-master-locator"] = "client-local"
	//args["x-message-ttl"] = int32(-1)
	//args["x-max-priority"] = int32(64)

	q, err := ch.QueueDeclare(
		Q01,   // name
		!ndq,  // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name, // queue name
		bind,   // routing key
		Ex,     // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")
}

func main() {
	if lb {
		ListBrokers()
		return
	}

	bk := brokers[bid]
	if bk == "" {
		log.Fatalf("Invalid broker name: %s", bid)
	}
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

	if declareQ {
		declareBindQueueOnly(ch)
		return
	}

	if deleteQ {
		deleteQueueOnly(ch)
		return
	}

	if autoDeclare {
		declare(ch)
	}

	if isRecv {
		recv(ch)
	} else {
		send(ch)
	}
}
