package repositories

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sdragos3/taskgo/models"
	"github.com/sdragos3/taskgo/persistence/database"
)

func Insert(task *models.Task) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(task)
	if err != nil {
		return err
	}
	err = db.Insert(task.Id.String(), bytes)
	if err != nil {
		return err
	}
	err = db.Close()
	return err
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
	err = db.Close()
	return tasks, nil
}

func GetById(id uuid.UUID) (*models.Task, error) {
	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	var task models.Task

	b, err := db.GetByKey(id.String())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &task)
	if err != nil {
		return nil, err
	}
	err = db.Close()
	return &task, err
}

func DeleteById(id uuid.UUID) error {
	db, err := database.Open()
	if err != nil {
		return err
	}
	err = db.Delete(id.String())
	if err != nil {
		return err
	}
	return db.Close()
}
