package repositories

import (
	"encoding/json"
	"github.com/sdragos3/taskgo/models"
	"github.com/sdragos3/taskgo/persistence/database"
)

func Create(task *models.Task) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(task)
	if err != nil {
		return err
	}
	return db.Insert(task.Id.String(), bytes)
}

func List() ([]models.Task, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	bucket, err := db.DumpBucket()
	if err != nil {
		return nil, err
	}
	var tasks []models.Task

	for _, elem := range bucket {
		var task models.Task
		if err := json.Unmarshal(elem, &task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
