package main

import (
	"github.com/bluele/slack"
	"reflect"
)

type MessageOpts struct {
	Options *slack.ChatPostMessageOpt
}

func (m *MessageOpts) GenerateTask(t Resource) error {
	reflected := reflect.ValueOf(t).Elem()

	m.Options.Attachments = []*slack.Attachment{
		&slack.Attachment{
			Title:     reflected.FieldByName("Name").String(),
			TitleLink: "https://asana.com",
			Text:      reflected.FieldByName("Notes").String(),
		},
	}

	return nil
}

func SendSlackMessage(slackUser string, options *MessageOpts) error {
	api := slack.New(config.SlackAPIKey)
	err := api.ChatPostMessage(slackUser, "", options.Options)

	if err != nil {
		return err
	}

	return nil
}

func NewMessageOptions() *MessageOpts {
	return &MessageOpts{
		Options: &slack.ChatPostMessageOpt{
			Username: "Asana Task Monitor",
			IconUrl:  "https://luna1.co/01f74a.png",
		},
	}
}
