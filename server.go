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

func handleReq(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	//fmt.Println(req.PostForm)
	params := req.PostForm

	limit, _ := strconv.Atoi(params["limit"][0])
	startPoint := params["startPoint"][0]
	directon := params["direction"][0]
	keyType := params["keyType"][0]
	valueType := params["valueType"][0]

	//log the incoming request
	log.Println("New Request with folowing Params :", "limit =", limit, "| startPoint =", startPoint, "| direction =", directon, "| keyType =", keyType, "| valueType =", valueType)

	var k []interface{}
	var val []interface{}

	iter := db.NewIterator(nil, nil)

	i := 0
	if startPoint == "null" {
		for iter.Next() {
			if i >= limit {
				break
			}
			appendKeyValueToResp(&k, &val, iter, keyType, valueType)
			i++
		}
	} else {

		if directon == "previous" {

			iter := db.NewIterator(nil, nil)
			for ok := iter.Seek(GetByteArray(startPoint)); ok; ok = iter.Prev() {
				// Use key/value.
				if i == 0 {
					i++
					//need to skip first one as it is the last value from previous page
					continue
				}

				if i > limit {
					break
				}

				appendKeyValueToResp(&k, &val, iter, keyType, valueType)
				i++

			}
			//the array needs to be fliped
			for i, j := 0, len(k)-1; i < j; i, j = i+1, j-1 {
				k[i], k[j] = k[j], k[i]
				val[i], val[j] = val[j], val[i]
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
				appendKeyValueToResp(&k, &val, iter, keyType, valueType)
				i++

			}
		}

	}

	type response struct {
		Keys   []interface{} `json:"keys"`
		Values []interface{} `json:"values"`
	}

	s := response{Keys: k, Values: val}

	s1, _ := json.Marshal(s)
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprint(res, string(s1))
}

func appendKeyValueToResp(keys *[]interface{}, values *[]interface{}, iter iterator.Iterator, keyType string, valueType string) {
	key := iter.Key()
	value := iter.Value()

	*keys = append(*keys, byteArrayToType(key, keyType))
	*values = append(*values, byteArrayToType(value, valueType))
}
