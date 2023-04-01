package helpers

import (
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func OpenDB(dbName string) (err error) {
	db, err = leveldb.OpenFile(dbName, nil)
	return err
}

func CloseDB() error {
	return db.Close()
}

// Insert is used to put key value pairs in database
func Insert(key []byte, value []byte) error {
	return db.Put(key, value, nil)
}

// ReadValue is used to retrieve the value of a given key from the database
func ReadValue(key []byte) ([]byte, error) {
	return db.Get(key, nil)
}

func InitTestDB() {

	for i := 0; i < 1000; i++ {
		Insert(intToByteArray(i), GetByteArray("hello from "+strconv.Itoa(i), "string"))
	}
}
