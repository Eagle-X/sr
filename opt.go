package main

import (
	"flag"
)

var (
	lb     bool
	isRecv bool
	bid    int
	count  int
)

func init() {
	flag.BoolVar(&lb, "l", false, "List all brokers")
	flag.BoolVar(&isRecv, "r", false, "Receive messages")
	flag.IntVar(&bid, "b", 1, "Broker ID")
	flag.IntVar(&count, "c", 10, "Message count")
	flag.Parse()
}
