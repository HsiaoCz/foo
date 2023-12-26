package main

import (
	"log"

	"github.com/HsiaoCz/foo/Comment/server"
)

var (
	network = "tcp"
	addr    = "127.0.0.1:3006"
)

func main() {
	if err := server.RegisterGrpc(network, addr); err != nil {
		log.Fatal(err)
	}
}
