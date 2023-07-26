package main

import (
	"log"

	api "github.com/benmorehouse/example-service/handlers"
)

func main() {
	log.Println("example_app_started_up")
	api.Start()
}
