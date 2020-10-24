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
		Insert(GetByteArray(i), GetByteArray("heelo"))
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
	k := []int{}
	val := []string{}
	for i := 0; i < 20; i++ {
		v, e := ReadValue(GetByteArray(i))

		if e != nil {
			fmt.Println("Error ", e)
		} else {
			//fmt.Println(i, string(v))
			k = append(k, i)
			val = append(val, string(v))
		}
	}

	type response struct {
		Keys   []int    `json:"keys"`
		Values []string `json:"values"`
	}

	s := response{Keys: k, Values: val}

	s1, _ := json.Marshal(s)
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprint(res, string(s1))
}

func serveIndex(res http.ResponseWriter, req *http.Request) {

	s := `<style> h1 {text-align:center;color:green;}
	table {
		font-family: arial, sans-serif;
		border-collapse: collapse;
		width: 100%;
	  }
	  
	  td, th {
		border: 1px solid #dddddd;
		text-align: left;
		padding: 8px;
	  }
	  
	  tr:nth-child(even) {
		background-color: #99ff99;
	  }</style> <h1>Hello World</h1>

	  <body>
	  `

	//add a table with all the values
	s += "<table><tr><th>Key</th><th>Value</th></tr>"
	for i := 0; i < 1000; i++ {
		v, e := ReadValue(GetByteArray(i))

		if e != nil {
			fmt.Println("Error ", e)
		} else {
			//fmt.Println(i, string(v))
			s = s + "<tr><td>" + strconv.Itoa(i) + "</td><td>" + string(v) + "</td></tr>"
		}
	}
	s += "</table>"
	res.Write([]byte(s))

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
