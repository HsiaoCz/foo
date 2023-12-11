package main

import (
	"log"

	"github.com/HsiaoCz/foo/router"
)

// foo i have no idea what i want to do
// so i named foo
// jsut foo

var (
	addr = "127.0.0.1:3001"
)

func main() {
	if err := router.Router(addr); err != nil {
		log.Fatal(err)
	}
}
