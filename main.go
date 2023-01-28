package main

import (
	"log"

	"github.com/mayankfawkes/httptoy/pkg/server"
)

func main() {
	log.Println("Server listening on http://0.0.0.0:8000 and http://127.0.0.1:8000")
	srv := server.Server()
	srv.Run(":8000")
}
