package main

import (
	"fmt"
	"log"
	"net/http"
)

var config *Config

func main() {
	config = GetConfig()

	server := &http.Server{
		Addr:    config.Address,
		Handler: Router(),
	}

	fmt.Printf("API Server started at: %s\n", config.Address)
	log.Fatalln(server.ListenAndServe())
}
