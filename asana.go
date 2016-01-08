package main

import (
	"fmt"
	"github.com/tambet/go-asana/asana"
	"time"
)

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

type User struct {
	Id    int        `json:"id"`
	Name  string     `json:"name"`
	Photo *UserPhoto `json:"photo"`
}

type UserPhoto struct {
	Image string `json:"image_60x60"`
}

func (p *Payload) RelayTask(slackUser string) error {
	for _, event := range p.Events {
		fmt.Printf("## %s / %s", event.Type, event.Action)
		if !event.IsRelayable() {
			continue
		}

		task, err := GetAsanaTask(event.Resource)
		if err != nil {
			return err
		}

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

func GetAsanaTask(id int) (*asana.Task, error) {
	task := new(asana.Task)

	err := asanaAPI.Request(fmt.Sprintf("tasks/%v", id), nil, task)
	if err != nil {
		fmt.Println(fmt.Sprintf("issue with task fetch: %v", id))
		return nil, err
	}

	return task, err
}

func (pe *PayloadEvent) IsRelayable() bool {
	if pe.Type != "task" || pe.Action != "added" {
		return false
	}

	return true
}

func GetAsanaTaskProject(task *asana.Task) string {
	if len(task.Projects) <= 0 {
		return "Unassigned"
	}

	return task.Projects[0].Name
}
