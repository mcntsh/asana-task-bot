package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

type ResourceData struct {
	Data Resource `json:"data"`
}

type Task struct {
	Id        int        `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	Name      string     `json:"name"`
	Notes     string     `json:"notes"`
	Projects  []*Project `json:"projects"`
}

type Project struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id    int        `json:"id"`
	Name  string     `json:"name"`
	Photo *UserPhoto `json:"photo"`
}

type UserPhoto struct {
	Image string `json:"image_60x60"`
}

type Resource interface {
	GetId() int
}

// -----
// Methods
// -----

func (p *Payload) RelayToSlack(slackUser string) error {
	for _, event := range p.Events {
		if !event.IsRelayable() {
			continue
		}

		task := &ResourceData{Data: &Task{Id: event.Resource}}

		err := task.GetResourceData(ASANA_TASK_ENDPOINT)
		if err != nil {
			panic(err)
		}

		msgOpts := NewMessageOptions()

		err = msgOpts.GenerateTask(task.Data)
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

func (t *Task) GetId() int { return t.Id }
func (u *User) GetId() int { return u.Id }

func (rd *ResourceData) GetResourceData(endpoint string) error {
	res, err := sendRequest("GET", fmt.Sprintf("%s/%v", endpoint, rd.Data.GetId()))
	if err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&rd)
	if err != nil {
		return err
	}

	return nil
}

func (pe *PayloadEvent) IsRelayable() bool {
	if pe.Type != "task" || pe.Action != "added" {
		return false
	}

	return true
}

// -----
// Helpers
// -----

func sendRequest(method, endpoint string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, endpoint, nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", config.AsanaAPIKey))

	res, err := client.Do(req)

	return res, err
}
