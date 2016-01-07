package main

import (
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
