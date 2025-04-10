package database

import "go.etcd.io/bbolt"

const dbName = "taskgo.db"

type TaskgoDb struct {
	db *bbolt.DB
}

func Init() error {
	db, err := bbolt.Open(dbName, 0600, nil)
	if err != nil {
		panic(err)
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func Open() (*TaskgoDb, error) {
	db, err := bbolt.Open(dbName, 0600, nil)
	return &TaskgoDb{db: db}, err
}

func (d *TaskgoDb) Insert(key string, value []byte) error {
	return d.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		return b.Put([]byte(key), value)
	})
}

func (d *TaskgoDb) DumpBucket() (map[string][]byte, error) {
	data := make(map[string][]byte)
	err := d.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		if b == nil {
			return nil
		}
		return b.ForEach(func(k, v []byte) error {
			data[string(k)] = append([]byte(nil), v...)
			return nil
		})
	})
	return data, err
}
