// +build dev
package main

import (
	"flag"
)

var err error

func main() {

	dbArg := flag.String("dbpath", "testdb", "Absolute Path to the database")
	flag.Parse()

	openDB(*dbArg)
	defer closeDB()

	if *dbArg == "testdb" {
		//this means that the argument was not supplied for the database name.
		//For testing purposes, initTestDB would create a database in the same directory from where this program is being called.
		go initTestDB()
	}
	serve()
}
