package routes

import (
	"fmt"
	// "github.com/gorilla/context"
	// "github.com/nlopes/slack"
	"net/http"
)

func RecieveWebhook(w http.ResponseWriter, r *http.Request) {
	// api := slack.New(fmt.Sprintf("%v", context.Get(r, "slack")))

	// groups, err := api.GetGroups(false)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, group := range groups {
	// 	fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	// }

	fmt.Printf("%v", r.Body)
}
