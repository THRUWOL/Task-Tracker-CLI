package models

import (
	"time"
)

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func NewTask(description string) Task {
	now := time.Now().Format(time.RFC3339)
	return Task{
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func IsValidStatus(status string) bool {
	return status == StatusTodo || status == StatusInProgress || status == StatusDone
}

func FormatDateFriendly(dateStr string) string {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return dateStr
	}
	return t.Format("02 Jan 2006 15:04")
}
