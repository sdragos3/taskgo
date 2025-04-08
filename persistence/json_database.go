package persistence

import (
	"fmt"
	"os"
)

const filePath = "./taskgo_database.json"

func Initialize() error {
	if databaseFileExists() {
		fmt.Println("Database file already exists")
		return nil
	}
	return createDatabaseFile()
}

func databaseFileExists() bool {
	_, err := os.Stat(filePath)
	return err == nil || !os.IsNotExist(err)
}

func createDatabaseFile() error {
	_, err := os.Create(filePath)
	return err
}
