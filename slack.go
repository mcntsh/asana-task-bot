package main

import (
	"fmt"
	"github.com/bluele/slack"
	"github.com/tambet/go-asana/asana"
)

const (
	NEW_TASK_MESSAGE = "You have been assigned a new task:"

	ASANA_TASK_URL = "https://app.asana.com/0/0"
)

type MessageOpts struct {
	Options *slack.ChatPostMessageOpt
}

func (m *MessageOpts) GenerateTask(task *asana.Task) error {
	taskProject := GetAsanaTaskProject(task)
	taskName := fmt.Sprintf("[%s] %s", taskProject, task.Name)

	m.Options.Attachments = []*slack.Attachment{
		&slack.Attachment{
			Title:     taskName,
			TitleLink: fmt.Sprintf("%s/%s", ASANA_TASK_URL, task.ID),
			Text:      task.Notes,
			Fields: []*slack.AttachmentField{
				&slack.AttachmentField{
					Title: "Assignee",
					Value: task.Assignee.Name,
					Short: true,
				},
			},
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
