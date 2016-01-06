package routes

import (
	"net/http"
)

func Respond(status int, response string, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Write([]byte(response))
}
