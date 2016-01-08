package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

func HandlerRecieveWebhook(w http.ResponseWriter, r *http.Request) {
	slackUser := r.URL.Query().Get("slack_user")
	payload := &Payload{}

	// Debug
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		// TODO: Handle this better
		fmt.Println(err)
	}

	// Relay task
	err = payload.RelayTask(slackUser)
	if err != nil {
		// TODO: Handle this better
		fmt.Println(err)
	}

	// Set the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
