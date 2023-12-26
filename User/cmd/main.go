package main

import (
	"log"

	"github.com/HsiaoCz/foo/User/server"
)

var (
	network = "tcp"
	addr    = "127.0.0.1:3005"
)

func main() {
	if err := server.ResGrpcServer(network, addr); err != nil {
		log.Fatal(err)
	}
}
