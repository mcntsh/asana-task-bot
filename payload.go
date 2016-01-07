package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	ASANA_TASK_ENDPOINT = "https://app.asana.com/api/1.0/tasks"
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

func (p *Payload) RelayToSlack() []*Task {
	var tasks []*Task

	for _, event := range p.Events {
		if !event.IsRelayable() {
			continue
		}

		task, err := event.GetTaskInfo()
		if err != nil {
			panic(err)
		}

		tasks = append(tasks, task)
	}

	return tasks
}

func (pe *PayloadEvent) IsRelayable() bool {
	if pe.Type != "task" || pe.Action != "added" {
		return false
	}

	return true
}

func (pe *PayloadEvent) GetTaskInfo() (*Task, error) {
	fmt.Println(pe.Resource)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%b", ASANA_TASK_ENDPOINT, pe.Resource), nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer: %v", config.AsanaAPIKey))

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	task := &Task{}

	err = json.NewDecoder(res.Body).Decode(&task)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", task)

	return task, nil
}
