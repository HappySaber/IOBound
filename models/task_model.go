package models

import "time"

type Task struct {
	ID                int           `json:"id"`
	Status            string        `json:"status"` // pending, in_progress, done
	CreatedAt         time.Time     `json:"created_at"`
	Text              string        `json:"text"`
	StartedAt         time.Time     `json:"started_at,omitempty"`
	CompletedAt       time.Time     `json:"completed_at,omitempty"`
	Duration          time.Duration `json:"duration"`
	DurationFormatted string        `json:"duration_formated"`
}

type TaskResponse struct {
	ID                int       `json:"id"`
	Status            string    `json:"status"` // pending, in_progress, done
	CreatedAt         time.Time `json:"created_at"`
	Text              string    `json:"text"`
	DurationFormatted string    `json:"duration_formated"`
}
