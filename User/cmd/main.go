package main

import (
	"log"

	"github.com/HsiaoCz/foo/User/pkg"
	"github.com/HsiaoCz/foo/User/server"
)

var (
	network   = "tcp"
	addr      = "127.0.0.1:3005"
	startTime = "2023-12-31"
	machineID = 1
)

func main() {
	if err := server.ResGrpcServer(network, addr); err != nil {
		log.Fatal(err)
	}
	if err := pkg.Init(startTime, int64(machineID)); err != nil {
		log.Fatal(err)
	}
}
