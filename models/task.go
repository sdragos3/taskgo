package models

import (
	"fmt"
	"github.com/google/uuid"
)

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

func (task *Task) String() string {
	return fmt.Sprintf("id: %s | title: \"%s\" | description: \"%s\" | completed: %t | priority: %s", task.Id, task.Title, task.Description, task.Completed, task.Priority)
}

func (task *Task) Update(title *string, description *string, priority *Priority) {
	if priority != nil {
		task.Priority = *priority
	}
	if title != nil {
		task.Title = *title
	}
	if description != nil {
		task.Description = *description
	}
}
