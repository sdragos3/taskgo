package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/sdragos3/taskgo/models"
	db "github.com/sdragos3/taskgo/persistence"
)

func Create(task *models.Task) error {
	b, err := json.MarshalIndent(task, "", " ")
	if err != nil {
		return err
	}
	_, err = db.Write(b)
	if err != nil {
		return err
	}
	fmt.Println("Task created")
	return nil
}

func List() ([]models.Task, error) {
	tasks := make([]models.Task, 0)

	dbBuffer, err := db.ReadAll()
	if err != nil {
		return tasks, err
	}
	err = json.Unmarshal(dbBuffer, &tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}
