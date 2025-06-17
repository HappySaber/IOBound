package models

import "time"

type Task struct {
	ID        int           `json:"id"`
	Status    string        `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	Duration  time.Duration `json:"duration"`
}
