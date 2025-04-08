package models

import "github.com/google/uuid"

type Task struct {
	Id          uuid.UUID
	Title       string
	Description string
	Completed   bool
	Priority    Priority
}

func TaskCreate(title string, description string, priority Priority) Task {
	return Task{
		Id:          uuid.New(),
		Title:       title,
		Description: description,
		Completed:   false,
		Priority:    priority,
	}
}
