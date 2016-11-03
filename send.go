package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func send(ch *amqp.Channel) {
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

	args := make(amqp.Table)
	args["x-expires"] = int32(10000)
	//args["x-message-ttl"] = int32(100000)

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

	returnCh := ch.NotifyReturn(make(chan amqp.Return))
	go func() {
		for r := range returnCh {
			//log.Println(r)
			log.Println(r.ReplyCode, r.ReplyText)
		}

	}()

	//ch.QueuePurge(q.Name, false)
	//body := bodyFrom(os.Args)
	//f := amqp.Transient
	f := amqp.Persistent
	body := make([]byte, 512)
	log.Printf("Send %d messages", count)
	for i := 0; i < count; i++ {
		copy(body, []byte(fmt.Sprint(i)))
		//log.Printf("===> %d", i)
		for {
			/*if f == amqp.Persistent {
				f = amqp.Transient
			} else {
				f = amqp.Persistent
			}*/
			err = ch.Publish(
				"logs-internal", // exchange
				"",              // routing key
				false,           // mandatory
				false,           // immediate
				amqp.Publishing{
					DeliveryMode: f,
					//DeliveryMode: amqp.Persistent,
					//DeliveryMode: amqp.Transient,
					ContentType: "text/plain",
					//Body:        body,
					//Body: []byte(body),
					Body: []byte(fmt.Sprint(i)),
				})
			if err != nil {
				failOnError(err, "Failed to publish a message")
			}
		}
		//time.Sleep(time.Microsecond * 20)
		//time.Sleep(time.Second)

	}
	//c := make(chan bool)
	//<-c

	//log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
