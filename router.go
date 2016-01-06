package main

import (
	"asana-task-bot/Godeps/_workspace/src/github.com/gorilla/mux"
	"asana-task-bot/Godeps/_workspace/src/github.com/justinas/alice"
	"asana-task-bot/middleware"
	"asana-task-bot/routes"
	"net/http"
)

var (
	appChain = alice.New(
		middleware.Logging,
		middleware.Cors,
		middleware.SlackAPI,
	)
)

func Router() http.Handler {
	r := mux.NewRouter()

	// REST Handlers

	r.Methods("POST").Path("/webhooks/recieve").Handler(appChain.ThenFunc(routes.RecieveWebhook))

	// Catch-all Handler

	r.PathPrefix("/").Handler(http.DefaultServeMux)

	return r
}
