package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

func recv(ch *amqp.Channel) {
	err := ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	err = ch.ExchangeDeclare(
		"logs-internal", // name
		"fanout",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare an exchange")
	err = ch.ExchangeDeclarePassive(
		"logs-internal", // name
		"fanout",        // type
		false,           // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare an exchange")
	err = ch.ExchangeBind(
		"logs-internal", // queue name
		"",              // routing key
		"logs",          // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	args := make(amqp.Table)
	args["x-expires"] = int32(10000)
	//args["x-message-ttl"] = int32(-1)
	//args["x-max-priority"] = int32(64)

	q, err := ch.QueueDeclare(
		"fff", // name
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,          // queue name
		"",              // routing key
		"logs-internal", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")
	err = ch.Qos(
		qosPcount, // prefetch count
		0,         // prefetch size
		qosGlobal, // global
	)
	failOnError(err, "Failed to Set Qos")
	//d, n, err := ch.Get(q.Name, true)
	//fmt.Println(d.Body, n, err)

	msgs, err := ch.Consume(
		q.Name, // queue
		"xxxx", // consumer
		noack,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to bind a queue")

	forever := make(chan bool)
	go func() {
		i := 0
		for d := range msgs {
			i++
			x, _ := strconv.ParseInt(string(d.Body), 10, 64)
			if verbose {
				fmt.Printf("=> %d, ==%d==, %t, %d, %v\n", len(d.Body), x, d.Redelivered, d.DeliveryTag, d.Priority)
			}
			if !noack {
				err = d.Ack(false)
				failOnError(err, "Failed to Ack")
			}
		}
		log.Printf("==========OK")
		close(forever)
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
