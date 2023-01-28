package main

import (
	"fmt"

	"github.com/mayankfawkes/httptoy/pkg/server"
)

var print = fmt.Println

func main() {
	print("Starting server hehe")
	srv := server.Server()
	srv.Run("127.0.0.1:8000")
}
