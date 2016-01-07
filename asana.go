package main

import (
	"fmt"
	"github.com/tambet/go-asana/asana"
	"time"
)

const (
	ASANA_TASK_ENDPOINT = "https://app.asana.com/api/1.0/tasks"
	ASANA_USER_ENDPOINT = "https://app.asana.com/api/1.0/users"
	ASANA_TASK_URL      = "https://app.asana.com/0/0"
)

// -----
// Types
// -----

type Payload struct {
	Events []*PayloadEvent `json:"events"`
}

type PayloadEvent struct {
	Resource  int       `json:"resource"`
	User      int       `json:"user"`
	Type      string    `json:"type"`
	Action    string    `json:"action"`
	CreatedAt time.Time `json:"created_at"`
	Parent    int       `json:"parent"`
}

func (p *Payload) RelayTask(slackUser string) error {
	for _, event := range p.Events {
		if !event.IsRelayable() {
			continue
		}

		task, err := GetAsanaTask(event.Resource)
		if err != nil {
			return err
		}

		// Get the assignee info
		// if task.Assignee != nil {
		// 	user, err := GetAsanaUser(task.Assignee.ID)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	task.Assignee = user
		// }

		msgOpts := NewMessageOptions()

		err = msgOpts.GenerateTask(task)
		if err != nil {
			return err
		}

		err = SendSlackMessage(slackUser, msgOpts)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetAsanaUser(id int64) (*asana.User, error) {
	user := new(asana.User)

	err := asanaAPI.Request(fmt.Sprintf("users/%v", id), nil, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetAsanaTask(id int) (*asana.Task, error) {
	task := new(asana.Task)

	err := asanaAPI.Request(fmt.Sprintf("tasks/%v", id), nil, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (pe *PayloadEvent) IsRelayable() bool {
	if pe.Type != "task" || pe.Action != "added" {
		return false
	}

	return true
}
