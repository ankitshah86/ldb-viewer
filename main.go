package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB
var err error

func main() {
	fmt.Println("Hello World")
	db, err = leveldb.OpenFile("testdb", nil)
	if err != nil {
		fmt.Println(err)
	}

	Insert(GetByteArray("key"), GetByteArray("value"))

	for i := 0; i < 1000; i++ {
		Insert(GetByteArray(i), GetByteArray("heelo"))
	}

	for i := 0; i < 1000; i++ {
		v, e := ReadValue(GetByteArray(i))

		if e != nil {
			fmt.Println("Error ", e)
		} else {
			fmt.Println(i, string(v))
		}
	}

	defer db.Close()
}

//Insert is used to put key value pairs in database
func Insert(key []byte, value []byte) error {
	return db.Put(key, value, nil)
}

//ReadValue is used to retrieve the value of a given key from the database
func ReadValue(key []byte) ([]byte, error) {
	return db.Get(key, nil)
}

//GetByteArray converts any type into byte array
func GetByteArray(any interface{}) []byte {
	return []byte(fmt.Sprintf("%v", any.(interface{})))
}
