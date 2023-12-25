package main

import (
	"log"

	"github.com/HsiaoCz/foo/Admin/server"
)

var (
	network = "tcp"
	addr    = "127.0.0.1:3003"
)

func main() {
	if err := server.ResGrpcServer(network, addr); err != nil {
		log.Fatal(err)
	}
}
