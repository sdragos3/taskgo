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
