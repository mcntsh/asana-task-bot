package main

import (
	"asana-task-bot/middleware"
	"asana-task-bot/routes"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

var (
	appChain = alice.New(
		middleware.Logging,
		middleware.Cors,
		middleware.SlackSecret,
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
