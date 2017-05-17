package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

func send(ch *amqp.Channel) {
	returnCh := ch.NotifyReturn(make(chan amqp.Return))
	go func() {
		for r := range returnCh {
			log.Println(r)
			log.Println(r.ReplyCode, r.ReplyText)
		}

	}()

	//ch.QueuePurge(q.Name, false)
	//body := bodyFrom(os.Args)
	//f := amqp.Transient
	f := amqp.Persistent
	if transient {
		f = amqp.Transient
	}
	body := make([]byte, 512)
	log.Printf("Send %d messages", count)
	if count == 0 {
		count = 10
	}
	for i := 0; i < count; i++ {
		copy(body, []byte(fmt.Sprint(i)))
		//log.Printf("===> %d", i)
		/*if f == amqp.Persistent {
			f = amqp.Transient
		} else {
			f = amqp.Persistent
		}*/
		if verbose {
			log.Printf("===> %d", i)
		}
		err := ch.Publish(
			Ex,
			//"logs-internal", // exchange
			//"amq.topic", // exchange
			"a001", // routing key
			false,  // mandatory
			false,  // immediate
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
		time.Sleep(time.Duration(sendLimit) * time.Millisecond)
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
