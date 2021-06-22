package db

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)

type Handler struct {
	db *bolt.DB
}

func (h *Handler) OpenDB() error {
	db, err := bolt.Open("task.db", 0600, nil)
	if err != nil {
		return err
	}
	h.db = db
	return nil
}

func (h *Handler) CloseDB() {
	h.db.Close()
}

type P struct {
	Name string
}

func (h *Handler) Write(bucketName string, s string) error {
	return h.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		id, _ := b.NextSequence()
		err := b.Put(itob(int(id)), []byte(s))
		if err != nil {
			return err
		}
		return nil
	})
}

func (h *Handler) List(bucketName string) ([]Task, error) {
	var ret []Task
	err := h.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			ret = append(ret, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (h *Handler) Delete(bucketName string, key int) error {
	err := h.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if err := b.Delete(itob(key)); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) InitBucket(bucketName string) error {
	return h.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	ret := binary.BigEndian.Uint64(b)
	return int(ret)
}

type Task struct {
	Key   int
	Value string
}
