default: build

build:
	go build -o sr broker.go  main.go  opt.go  recv.go  send.go

.PHONY: sr
