package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb/iterator"
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

//DataType can be used to parse the datatype required by the frontend
type DataType string

type response struct {
	Keys   []interface{} `json:"keys"`
	Values []interface{} `json:"values"`
}

const (
	previous Direction = "previous"
	next     Direction = "next"
	none     Direction = "none"
)

const (
	integerT           DataType = "integer"
	int32LittleEndian  DataType = "int32LE"
	int32BigEndian     DataType = "int32BE"
	uint32LittleEndian DataType = "uint32LE"
	uint32BigEndian    DataType = "uint32BE"
	int64LittleEndian  DataType = "int64LE"
	int64BigEndian     DataType = "int64BE"
	uint64LittleEndian DataType = "uint64LE"
	uint64BigEndian    DataType = "uint64BE"

	float32BigEndian    DataType = "float32BE"
	float32LittleEndian DataType = "float32LE"
	float64LittleEndian DataType = "float64LE"
	float64BigEndian    DataType = "float64BE"

	hexT       DataType = "hexadecimal"
	stringT    DataType = "string"
	byteArrayT DataType = "bytearray"
	booleanT   DataType = "boolean"
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
	keyType := DataType(params["keyType"][0])
	valueType := DataType(params["valueType"][0])

	//log the incoming request
	log.Println("New Request with folowing Params :", "limit =", limit, "| direction =", direction, "| keyType =", keyType, "| valueType =", valueType)

	k, v := getKVSlice(limit, direction, keyType, valueType)

	s := response{Keys: k, Values: v}

	s1, _ := json.Marshal(s)
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprint(res, string(s1))
}

func getKVSlice(limit int, direction Direction, keyType DataType, valueType DataType) ([]interface{}, []interface{}) {
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

			readKVfromIter(iter, i)
			i++
		}
	} else {
		iter.Seek(firstElement)
		for iter.Prev() {
			if i >= limit {
				break
			}
			readKVfromIter(iter, i)
			i++
		}

		//the array needs to be fliped
		flipSlice(&keysCache)
		flipSlice(&valsCache)
	}

	firstElement = keysCache[0]
	lastElement = keysCache[limit-1]
}

func flipSlice(inSlice *[][]byte) {
	for i, j := 0, len(*inSlice)-1; i < j; i, j = i+1, j-1 {
		(*inSlice)[i], (*inSlice)[j] = (*inSlice)[j], (*inSlice)[i]
	}
}

func readKVfromIter(iter iterator.Iterator, i int) {

	keysCache[i] = copyByteArray(iter.Key())
	valsCache[i] = copyByteArray(iter.Value())
}
