package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func serve() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/data", handleReq)
	fmt.Println("Serving on port 8080, To see the database, kindly go to http://localhost:8080 on the browser of your choice.")
	createLogFile()
	http.ListenAndServe(":8080", nil)
}

var firstElement []byte
var lastElement []byte

var keysCache [][]byte
var valsCache [][]byte

//Direction can be used for pagination or lack thereof
type Direction string

type response struct {
	Keys   []interface{} `json:"keys"`
	Values []interface{} `json:"values"`
}

const (
	previous Direction = "previous"
	next     Direction = "next"
	none     Direction = "none"
)

/*
Parameters to be accepted from front end
limit
keytype
valuetype
direction : prev, next, none
*/

func handleReq(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	params := req.PostForm
	limit, _ := strconv.Atoi(params["limit"][0])
	direction := Direction(params["direction"][0])
	keyType := params["keyType"][0]
	valueType := params["valueType"][0]

	//log the incoming request
	log.Println("New Request with folowing Params :", "limit =", limit, "| direction =", direction, "| keyType =", keyType, "| valueType =", valueType)

	k, v := getKVSlice(limit, direction, keyType, valueType)

	s := response{Keys: k, Values: v}

	s1, _ := json.Marshal(s)
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprint(res, string(s1))
}

func getKVSlice(limit int, direction Direction, keyType string, valueType string) ([]interface{}, []interface{}) {
	var keys []interface{}
	var values []interface{}

	if direction != none || len(keysCache) != limit {
		rebuildCache(limit, direction)
	}

	keys = make([]interface{}, limit)
	values = make([]interface{}, limit)

	for i := 0; i < limit; i++ {
		keys[i] = byteArrayToType(keysCache[i], keyType)
		values[i] = byteArrayToType(valsCache[i], valueType)
	}

	return keys, values
}

func rebuildCache(limit int, direction Direction) {

	iter := db.NewIterator(nil, nil)
	if direction == none && len(keysCache) != 0 {
		//this means that only limit was changed
		iter.Seek(firstElement)
		iter.Prev()
	}

	keysCache = make([][]byte, limit)
	valsCache = make([][]byte, limit)
	i := 0

	if direction != previous {

		if direction == next {
			iter.Seek(lastElement) //go to the last key from previous slice
		}

		for iter.Next() {

			if i >= limit {
				break
			}

			keysCache[i] = copyByteArray(iter.Key())
			valsCache[i] = copyByteArray(iter.Value())
			i++
		}
	} else {
		iter.Seek(firstElement)
		for iter.Prev() {
			if i >= limit {
				break
			}
			keysCache[i] = copyByteArray(iter.Key())
			valsCache[i] = copyByteArray(iter.Value())
			i++
		}

		//the array needs to be fliped
		for i, j := 0, len(keysCache)-1; i < j; i, j = i+1, j-1 {
			keysCache[i], keysCache[j] = keysCache[j], keysCache[i]
			valsCache[i], valsCache[j] = valsCache[j], valsCache[i]
		}
	}

	firstElement = keysCache[0]
	lastElement = keysCache[limit-1]
}
