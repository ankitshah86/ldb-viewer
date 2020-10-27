package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/status-im/keycard-go/hexutils"
)

func createLogFile() {

	if _, er := os.Stat("Request_logs"); os.IsNotExist(er) {
		os.Mkdir("Request_logs", 0777)
	}

	f, err := os.Create("Request_logs/" + strconv.FormatInt(time.Now().Unix(), 10) + ".log")

	if err != nil {
		log.Fatalf("Error in creating log file %s\n", err)
	}

	logfile, err := os.OpenFile(f.Name(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(logfile)
	log.Println("LDB-Viewer server started")
	if err != nil {
		log.Fatalf("Error in opening Log File %s\n", err)
	}
}

//GetByteArray converts variable of any type into byte array
func GetByteArray(any interface{}, keyType string) []byte {

	if keyType == "hexadecimal" {
		b, e := hex.DecodeString(any.(string))
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println("bytes", b)
		return b
	}

	return []byte(fmt.Sprintf("%v", any.(interface{})))
}

func byteArrayToInt(b []byte) int {
	r, _ := strconv.Atoi(string(b))
	return r
}

func byteArrayToType(b []byte, bType string) interface{} {
	var r interface{}
	if bType == "string" {
		r = string(b)
	} else if bType == "integer" {
		r = binary.BigEndian.Uint32(b)
	} else if bType == "hexadecimal" {
		r = hexutils.BytesToHex(b)
	} else if bType == "boolean" {
		r, _ = strconv.ParseBool(string(b))
	} else if bType == "bytearray" {
		rb := make([]byte, len(b))
		//This needs to be done to ensure that actual values are appended as opposed to pointers
		copy(rb, b)
		return rb
	}
	return r
}

func intToByteArray(num int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(num))
	return b
}
