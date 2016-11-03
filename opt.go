package main

import (
	"flag"
)

var (
	lb        bool
	isRecv    bool
	verbose   bool
	noack     bool
	bid       int
	count     int
	qosPcount int
	qosGlobal bool
)

func init() {
	flag.BoolVar(&lb, "l", false, "List all brokers")
	flag.BoolVar(&isRecv, "r", false, "Receive messages")
	flag.BoolVar(&verbose, "v", false, "Verbose log info")
	flag.BoolVar(&noack, "noack", false, "Consume with noack")
	flag.IntVar(&bid, "b", 0, "Broker ID")
	flag.IntVar(&count, "c", 10, "Message count")
	flag.IntVar(&qosPcount, "qos-pc", 0, "Qos pretch count")
	flag.BoolVar(&qosGlobal, "qos-g", false, "Qos global")
	flag.Parse()
}
