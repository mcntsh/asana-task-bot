package middleware

import (
	"asana-task-bot/config"
	"github.com/gorilla/context"
	"net/http"
)

func SlackAPI(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "slack", config.Configuration.APIKey)
		h.ServeHTTP(w, r)
	})
}
