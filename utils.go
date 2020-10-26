package main

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/status-im/keycard-go/hexutils"
)

//GetByteArray converts variable of any type into byte array
func GetByteArray(any interface{}) []byte {
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
		r = "0x" + hexutils.BytesToHex(b)
	} else if bType == "boolean" {
		r, _ = strconv.ParseBool(string(b))
	} else if bType == "bytearray" {
		return b
	}
	return r
}
