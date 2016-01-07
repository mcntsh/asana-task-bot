package main

import (
	"fmt"
	"github.com/bluele/slack"
	"github.com/tambet/go-asana/asana"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

var config *Config
var asanaAPI *asana.Client
var slackAPI *slack.Slack

func main() {
	config = GetConfig()

	tok := &oauth2.Token{
		AccessToken: config.AsanaAPIKey,
	}

	tokSrc := oauth2.StaticTokenSource(tok)
	tokenAuth := oauth2.NewClient(context.TODO(), tokSrc)

	asanaAPI = asana.NewClient(tokenAuth)
	slackAPI = slack.New(config.SlackAPIKey)

	server := &http.Server{
		Addr:    config.Address,
		Handler: Router(),
	}

	fmt.Printf("API Server started at: %s\n", config.Address)
	log.Fatalln(server.ListenAndServe())
}
