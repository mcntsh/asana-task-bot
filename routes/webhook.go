package routes

import (
	"fmt"
	"github.com/bluele/slack"
	"github.com/gorilla/context"
	"net/http"
)

const (
	USER_JON = "U039YMRF8"
)

func RecieveWebhook(w http.ResponseWriter, r *http.Request) {
	api := slack.New(fmt.Sprintf("%v", context.Get(r, "slack")))

	api.ChatPostMessage(USER_JON, "Something happened on Asana!", nil)

	w.Header().Set("X-Hook-Secret", fmt.Sprintf("%v", context.Get(r, "slack-secret")))
	Respond(http.StatusOK, "Done.", w)
}
