package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

func recv(ch *amqp.Channel) {
	err := ch.Qos(
		qosPcount, // prefetch count
		0,         // prefetch size
		qosGlobal, // global
	)
	failOnError(err, "Failed to Set Qos")

	msgs, err := ch.Consume(
		Q01,    // queue
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
			if !noack && !notAck {
				err = d.Ack(false)
				failOnError(err, "Failed to Ack")
			}
			if i == count {
				break
			}
		}
		log.Printf("==========OK")
		close(forever)
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
