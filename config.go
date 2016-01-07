package main

import (
	"flag"
	"fmt"
)

type Config struct {
	Address     string
	SlackAPIKey string
	AsanaAPIKey string
}

var (
	portNumber  = flag.Int("port", 3000, "Port number")
	slackAPIKey = flag.String("slack-key", "", "Slack API authentication token")
	asanaAPIKey = flag.String("asana-key", "", "Asana API authentication token")
)

func GetConfig() *Config {
	flag.Parse()

	config := &Config{}

	config.Address = fmt.Sprintf(":%v", *portNumber)
	config.SlackAPIKey = *slackAPIKey
	config.AsanaAPIKey = *asanaAPIKey

	return config
}
