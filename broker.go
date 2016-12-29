package main

import (
	"fmt"
	"sort"
)

var brokers = map[string]string{
	"0":   "amqp://guest:guest@192.168.10.11:35672/test",
	"1":   "amqp://guest:guest@192.168.10.12:35672/test",
	"2":   "amqp://guest:guest@192.168.10.12:35673/test",
	"3":   "amqp://test:test@192.168.10.11:5672",
	"4":   "amqp://test:test@192.168.10.12:5672",
	"5":   "amqp://annatar:m3n23R3O@vpcal-arch-q-1.vm.elenet.me:5672/arch.q_annatar",
	"6":   "amqp://annatar:m3n23R3O@vpca-arch-q-1.vm.elenet.me:5672/arch.q_annatar",
	"q5":  "amqp://tester:tester@vpcl-mq-test-5.vm.elenet.me:5674/mirror_test",
	"q8":  "amqp://tester:tester@vpcl-mq-test-8.vm.elenet.me:5674/mirror_test",
	"q9":  "amqp://tester:tester@vpcl-mq-test-9.vm.elenet.me:5674/mirror_test",
	"q10": "amqp://tester:tester@vpcl-mq-test-10.vm.elenet.me:5674/mirror_test",
}

func ListBrokers() {
	var bs []string
	for n := range brokers {
		bs = append(bs, n)
	}
	sort.Strings(bs)
	for _, n := range bs {
		fmt.Printf("%s: %s\n", n, brokers[n])
	}
}
