// +build dev
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB
var err error

func main() {

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/data", handleReq)

	dbArg := flag.String("dbpath", "testdb", "Absolute Path to the database")
	flag.Parse()

	fmt.Println(*dbArg)

	db, err = leveldb.OpenFile(*dbArg, nil)
	if err != nil {
		fmt.Println(err)
	}

	Insert(GetByteArray("key"), GetByteArray("value"))

	for i := 0; i < 1000; i++ {
		Insert(GetByteArray(i), GetByteArray("hello from "+strconv.Itoa(i)))
	}

	for i := 0; i < 1000; i++ {
		_, e := ReadValue(GetByteArray(i))

		if e != nil {
			fmt.Println("Error ", e)
		} else {
			//fmt.Println(i, string(v))
		}
	}

	defer db.Close()
	http.ListenAndServe(":8080", nil)
}

func handleReq(res http.ResponseWriter, req *http.Request) {

	//res.Write([]byte("Hello world"))
	req.ParseForm()
	fmt.Println(req.PostForm)
	params := req.PostForm

	startPoint := params["startPoint"][0]

	k := []int{}
	val := []string{}

	iter := db.NewIterator(nil, nil)
	limit, _ := strconv.Atoi(params["limit"][0])

	i := 0
	if startPoint == "null" {
		for iter.Next() {
			if i >= limit {
				break
			}
			key := iter.Key()
			value := iter.Value()

			k = append(k, byteArrayToInt(key))
			val = append(val, string(value))

			i++
		}
	} else {
		iter := db.NewIterator(nil, nil)
		for ok := iter.Seek(GetByteArray(startPoint)); ok; ok = iter.Next() {
			// Use key/value.
			if i == 0 {
				i++
				//need to skip first one as it is the last value from previous page
				continue
			}

			if i > limit {
				break
			}
			key := iter.Key()
			value := iter.Value()

			k = append(k, byteArrayToInt(key))
			val = append(val, string(value))
			i++

		}
	}

	/*
		for i := 0; i < n; i++ {
			v, e := ReadValue(GetByteArray(i))

			if e != nil {
				fmt.Println("Error ", e)
			} else {
				//fmt.Println(i, string(v))
				k = append(k, i)
				val = append(val, string(v))
			}
		}
	*/

	type response struct {
		Keys   []int    `json:"keys"`
		Values []string `json:"values"`
	}

	s := response{Keys: k, Values: val}

	s1, _ := json.Marshal(s)
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprint(res, string(s1))
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

func byteArrayToInt(b []byte) int {
	//add better handling here
	r, _ := strconv.Atoi(string(b))
	return r
}
