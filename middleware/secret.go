package middleware

import (
	"github.com/gorilla/context"
	"net/http"
)

func SlackSecret(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "slack-secret", r.Header.Get("X-Hook-Secret"))
		h.ServeHTTP(w, r)
	})
}
