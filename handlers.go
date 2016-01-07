package main

import (
	// "github.com/bluele/slack"
	"encoding/json"
	"net/http"
)

func HandlerRecieveWebhook(w http.ResponseWriter, r *http.Request) {
	slackUser := r.URL.Query().Get("slack_user")
	payload := &Payload{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		panic(err)
	}

	// Relay task
	err = payload.RelayTask(slackUser)
	if err != nil {
		panic(err)
	}

	// Set the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
