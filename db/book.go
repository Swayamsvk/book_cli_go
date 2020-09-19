package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var bookBucket = []byte("books")
var db *bolt.DB

type Book struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bookBucket)
		return err
	})
}

func CreateBook(book string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bookBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		//keep data in db
		key := itob(int(id64))
		return b.Put(key, []byte(book))
	})
	if err != nil {
		return -1, err
	}
	return id, nil

}

func AllBooks() ([]Book, error) {
	var books []Book
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bookBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			books = append(books, Book{
				Key:   btoi(k),
				Value: string(v),
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return books, nil
}

func DeleteBook(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bookBucket)
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
