package persistence

import (
	"fmt"
	"os"
)

const dbPath = "./taskgo_database.json"

func Initialize() error {
	if databaseFileExists() {
		fmt.Println("Database file already exists")
		return nil
	}
	return createDatabaseFile()
}

func Delete() error {
	if databaseFileExists() {
		return os.Remove(dbPath)
	}
	return nil
}

func Write(bytes []byte) (int, error) {
	f, err := os.OpenFile(dbPath, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return 0, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("%s", err.Error())
		}
	}(f)
	return f.Write(bytes)
}

func databaseFileExists() bool {
	_, err := os.Stat(dbPath)
	return err == nil || !os.IsNotExist(err)
}

func createDatabaseFile() error {
	_, err := os.Create(dbPath)
	return err
}
