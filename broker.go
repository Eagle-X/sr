package main

import (
	"fmt"
)

var brokers = []string{
	"amqp://guest:guest@192.168.10.11:35672/test",
	"amqp://guest:guest@192.168.10.12:35672/test",
	"amqp://guest:guest@192.168.10.12:35673/test",
	"amqp://test:test@192.168.10.11:5672",
	"amqp://test:test@192.168.10.12:5672",
	"amqp://annatar:m3n23R3O@vpcal-arch-q-1.vm.elenet.me:5672/arch.q_annatar",
	"amqp://annatar:m3n23R3O@vpca-arch-q-1.vm.elenet.me:5672/arch.q_annatar",
}

func ListBrokers() {
	for i, b := range brokers {
		fmt.Printf("%d: %s\n", i, b)
	}
}
