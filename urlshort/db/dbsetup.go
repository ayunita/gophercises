package boltdb

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

func SetupDB() (*bolt.DB, error) {
	// Open the example.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("example.db", 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		return nil, err
	}

	// Read-write transaction
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("PATHURLS"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("DB setup done...")

	if err = addPathUrl(db, "/bolt", "https://pkg.go.dev/github.com/boltdb/bolt"); err != nil {
		return nil, err
	}

	if err = addPathUrl(db, "/go-dev", "https://pkg.go.dev/"); err != nil {
		return nil, err
	}

	return db, nil
}

func addPathUrl(db *bolt.DB, path string, url string) error {
	pathUrl := PathUrl{Path: path, URL: url}
	val, err := json.Marshal(pathUrl)
	if err != nil {
		return err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		if err := tx.Bucket([]byte("PATHURLS")).Put([]byte(path), val); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	fmt.Println("Entry added...")
	return nil
}

type PathUrl struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}
