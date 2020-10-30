package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math"
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
	} else if keyType == "integer" {
		s0 := any.(string)
		s1, _ := strconv.Atoi(s0)
		s := intToByteArray(s1)
		return s
	}

	return []byte(fmt.Sprintf("%v", any.(interface{})))
}

func byteArrayToInt(b []byte) int {
	r, _ := strconv.Atoi(string(b))
	return r
}

func byteArrayToType(b []byte, bType DataType) interface{} {
	var r interface{}

	switch bType {
	case stringT:
		r = string(b)
	case hexT:
		r = hexutils.BytesToHex(b)
	case booleanT:
		r, _ = strconv.ParseBool(string(b))
	case byteArrayT:
		rb := make([]byte, len(b))
		copy(rb, b)
		return rb
	case integerT:
		r = binary.BigEndian.Uint64(b)
	case int32BigEndian:
		r = int32(binary.BigEndian.Uint32(b))
	case int32LittleEndian:
		r = int32(binary.LittleEndian.Uint32(b))
	case uint32BigEndian:
		r = binary.BigEndian.Uint32(b)
	case uint32LittleEndian:
		r = binary.LittleEndian.Uint32(b)
	case int64BigEndian:
		r = int64(binary.BigEndian.Uint64(b))
	case int64LittleEndian:
		r = int64(binary.LittleEndian.Uint64(b))
	case uint64BigEndian:
		r = binary.BigEndian.Uint64(b)
	case uint64LittleEndian:
		r = binary.LittleEndian.Uint64(b)
	case float32BigEndian:
		r = math.Float32frombits(binary.BigEndian.Uint32(b))
	case float32LittleEndian:
		r = math.Float32frombits(binary.LittleEndian.Uint32(b))
	case float64BigEndian:
		r = math.Float64frombits(binary.BigEndian.Uint64(b))
	case float64LittleEndian:
		r = math.Float64frombits(binary.LittleEndian.Uint64(b))

	}

	/*
		if bType == stringT {
			r = string(b)
		} else if bType == "integer" {
			r = binary.BigEndian.Uint64(b)
		} else if bType == hexT {
			r = hexutils.BytesToHex(b)
		} else if bType == booleanT {
			r, _ = strconv.ParseBool(string(b))
		} else if bType == byteArrayT {
			rb := make([]byte, len(b))
			//This needs to be done to ensure that actual values are appended as opposed to pointers
			copy(rb, b)
			return rb
		}
	*/
	return r
}

func intToByteArray(num int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(num))
	return b
}

func copyByteArray(b []byte) []byte {
	r := make([]byte, len(b))
	copy(r, b)
	return r
}
