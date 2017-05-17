package main

import (
	"flag"
)

var (
	broker      string
	lb          bool
	isRecv      bool
	verbose     bool
	noack       bool
	notAck      bool
	transient   bool
	md5Sum      bool
	bid         string
	count       int
	qosPcount   int
	qosGlobal   bool
	ndq         bool
	autoDeclare bool
	declareQ    bool
	deleteQ     bool
	bind        string
	sendLimit   int
)

func init() {
	flag.StringVar(&broker, "broker", "", "Broker connect string")
	flag.BoolVar(&lb, "l", false, "List all brokers")
	flag.BoolVar(&isRecv, "r", false, "Receive messages")
	flag.BoolVar(&verbose, "v", false, "Verbose log info")
	flag.BoolVar(&noack, "noack", false, "Consume with noack")
	flag.BoolVar(&notAck, "not-ack", false, "Not send ack")
	flag.BoolVar(&transient, "t", false, "Messsage is transient")
	flag.BoolVar(&md5Sum, "md5", false, "Get message md5 checksum")
	flag.StringVar(&bid, "b", "0", "Broker name")
	flag.IntVar(&count, "c", 0, "Message count")
	flag.IntVar(&qosPcount, "qos-pc", 0, "Qos pretch count")
	flag.BoolVar(&qosGlobal, "qos-g", false, "Qos global")
	flag.BoolVar(&ndq, "ndq", false, "None-durable queue")
	flag.StringVar(&Q01, "q", "fff", "Queue name")
	flag.StringVar(&Ex, "e", "logs-internal", "Exchange name")
	flag.StringVar(&bind, "bind", "", "Binding key")
	flag.BoolVar(&autoDeclare, "declare", false, "If auto-declare queue and exchange")
	flag.BoolVar(&declareQ, "declare_q", false, "Declare queue")
	flag.BoolVar(&deleteQ, "delete_q", false, "Delete queue")
	flag.IntVar(&sendLimit, "send_limit", 0, "Send limit")
	flag.Parse()
}
