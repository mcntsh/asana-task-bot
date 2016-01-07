package main

import (
	"github.com/bluele/slack"
	"github.com/tambet/go-asana/asana"
)

type MessageOpts struct {
	Options *slack.ChatPostMessageOpt
}

func (m *MessageOpts) GenerateTask(task *asana.Task) error {
	m.Options.Attachments = []*slack.Attachment{
		&slack.Attachment{
			Title:     task.Name,
			TitleLink: "https://asana.com",
			Text:      task.Notes,
		},
	}

	return nil
}

func SendSlackMessage(slackUser string, options *MessageOpts) error {
	err := slackAPI.ChatPostMessage(slackUser, "", options.Options)

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
