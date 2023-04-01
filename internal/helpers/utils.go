package helpers

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
	logDir := "Request_logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			log.Fatalf("Error creating '%s' directory: %v", logDir, err)
		}
	}

	logFileName := logDir + "/" + strconv.FormatInt(time.Now().Unix(), 10) + ".log"
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error creating log file '%s': %v", logFileName, err)
	}

	log.SetOutput(logFile)
	log.Println("LDB-Viewer server started")
}

// GetByteArray converts variable of any type into byte array
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

	return []byte(fmt.Sprintf("%v", any))
}

// byteArrayToType converts byte array to the required type
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

	return r
}

// intToByteArray converts integer to byte array
func intToByteArray(num int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(num))
	return b
}

// copyByteArray copies byte array
func copyByteArray(b []byte) []byte {
	r := make([]byte, len(b))
	copy(r, b)
	return r
}
