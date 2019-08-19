package db

import (
	"fmt"
	"sort"

	"github.com/mitchellh/go-homedir"
	bolt "go.etcd.io/bbolt"
)

type DB struct {
	conn *bolt.DB
}

func NewConnection() (DB, error) {
	hd, err := homedir.Dir()
	if err != nil {
		return DB{}, err
	}
	db, err := bolt.Open(fmt.Sprintf("%s/db.bolt", hd), 0600, nil)
	if err != nil {
		return DB{}, err
	}
	return DB{db}, nil
}

func (db *DB) AddTask(task string) error {
	err := db.conn.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("tasks"))

		if err != nil {
			return err
		}

		// Set the value "bar" for the key "foo".
		if err = b.Put([]byte(task), []byte("true")); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (db *DB) ListTasks() ([]string, error) {
	var ret []string

	err := db.conn.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("tasks"))

		if err != nil {
			return err
		}

		return b.ForEach(func(k, v []byte) error {
			ret = append(ret, string(k))
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	sort.Strings(ret)

	return ret, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

func (db *DB) DeleteTask(taskId int) error {
	tasks, err := db.ListTasks()

	if err != nil {
		panic(err)
	}

	if len(tasks) < taskId {
		return nil
	}

	taskToDelete := tasks[taskId]

	return db.conn.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("tasks"))

		if err != nil {
			return err
		}

		return b.Delete([]byte(taskToDelete))
	})
}
