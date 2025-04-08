package persistence

import (
	"fmt"
	"io"
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
	f, err := os.OpenFile(dbPath, os.O_WRONLY, 0644)

	if err != nil {
		return 0, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("%s", err.Error())
		}
	}(f)
	dbEmpty, err := databaseEmpty(f)
	if err != nil {
		return 0, err
	}
	if dbEmpty {
		bytes = append([]byte("[\n"), bytes...)
		bytes = append(bytes, []byte("\n]")...)
	} else {
		_, err := f.Seek(-2, io.SeekEnd)
		if err != nil {
			return 0, err
		}
		bytes = append([]byte(","), bytes...)
		bytes = append(bytes, []byte("\n]")...)
	}
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

func databaseEmpty(file *os.File) (bool, error) {
	info, err := file.Stat()
	if err != nil {
		return false, err
	}
	if info.Size() == 0 {
		return true, nil
	}
	return false, nil
}
