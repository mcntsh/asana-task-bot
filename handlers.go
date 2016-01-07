package main

import (
	"fmt"
	"github.com/bluele/slack"
	"net/http"
	"net/http/httputil"
)

func HandlerRecieveWebhook(w http.ResponseWriter, r *http.Request) {
	slackUser := r.URL.Query().Get("slack_user")

	api := slack.New(config.APIKey)
	err := api.ChatPostMessage(slackUser, "Something happe1ned on Asana!", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("NOT OK"))
	}

	// TEST
	postData, err := httputil.DumpRequest(r, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(postData))

	// Set the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
