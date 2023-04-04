package helpers

import (
	"errors"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

// db is the global leveldb database instance
var db *leveldb.DB

// OpenDB opens the LevelDB database with the specified name and returns any error encountered.
func OpenDB(dbName string) (err error) {
	if dbName == "" {
		return errors.New("empty dbName provided")
	}

	db, err = leveldb.OpenFile(dbName, nil)
	return err
}

// CloseDB closes the LevelDB database and returns any error encountered.
func CloseDB() error {
	if db == nil {
		return errors.New("database not initialized")
	}
	return db.Close()
}

// Insert puts a key-value pair into the database and returns any error encountered.
func Insert(key []byte, value []byte) error {
	if db == nil {
		return errors.New("database not initialized")
	}

	if len(key) == 0 {
		return errors.New("empty key provided")
	}

	return db.Put(key, value, nil)
}

// ReadValue retrieves the value of a given key from the database and returns the value along with any error encountered.
func ReadValue(key []byte) ([]byte, error) {
	if db == nil {
		return nil, errors.New("database not initialized")
	}

	if len(key) == 0 {
		return nil, errors.New("empty key provided")
	}

	return db.Get(key, nil)
}

// InitTestDB initializes the test database with 1000 key-value pairs.
func InitTestDB() error {
	if db == nil {
		return errors.New("database not initialized")
	}

	for i := 0; i < 1000; i++ {
		key := intToByteArray(i)
		value := GetByteArray("hello from "+strconv.Itoa(i), "string")
		err := Insert(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}
