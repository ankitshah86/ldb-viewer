package main

import (
	"fmt"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func openDB(dbName string) {
	db, err = leveldb.OpenFile(dbName, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func closeDB() {
	err := db.Close()
	if err != nil {
		fmt.Println(err)
	}
}

//Insert is used to put key value pairs in database
func Insert(key []byte, value []byte) error {
	return db.Put(key, value, nil)
}

//ReadValue is used to retrieve the value of a given key from the database
func ReadValue(key []byte) ([]byte, error) {
	return db.Get(key, nil)
}

func initTestDB() {

	for i := 0; i < 1000; i++ {
		Insert(intToByteArray(i), GetByteArray("hello from "+strconv.Itoa(i), "string"))
	}
}
