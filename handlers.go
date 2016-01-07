package main

import (
	// "github.com/bluele/slack"
	"encoding/json"
	"net/http"
)

func HandlerRecieveWebhook(w http.ResponseWriter, r *http.Request) {
	// slackUser := r.URL.Query().Get("slack_user")

	// api := slack.New(config.SlackAPIKey)
	// err := api.ChatPostMessage(slackUser, "Something happe1ned on Asana!", nil)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("NOT OK"))
	// }

	payload := &Payload{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		panic(err)
	}

	payload.RelayToSlack()

	// Set the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
