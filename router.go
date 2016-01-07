package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

var (
	appChain = alice.New(
		MiddlewareLogging,
		MiddlewareJSON,
		MiddlewareCors,
	)
)

func Router() http.Handler {
	r := mux.NewRouter()

	// REST Handlers

	r.Methods("POST").Path("/webhooks/recieve").Handler(appChain.ThenFunc(HandlerRecieveWebhook))

	// Catch-all Handler

	r.PathPrefix("/").Handler(http.DefaultServeMux)

	return r
}
