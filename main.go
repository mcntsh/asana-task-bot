package main

import (
	"asana-task-bot/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ParseConfig()

	server := &http.Server{
		Addr:    config.Configuration.Address,
		Handler: Router(),
	}

	fmt.Printf("API Server started at: %s\n", config.Configuration.Address)
	log.Fatalln(server.ListenAndServe())
}
