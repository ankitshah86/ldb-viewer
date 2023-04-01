//go:build dev
// +build dev

package main

import (
	"flag"
	"log"
	"os"

	helpers "github.com/ankitshah86/ldb-viewer/internal/helpers"
)

var err error

func main() {

	dbPath := flag.String("dbpath", "testdb", "Absolute path to the database (default: 'testdb' in the current directory)")
	flag.Parse()

	if err := os.RemoveAll("testdb/"); err != nil {
		log.Fatalf("Failed to remove testdb: %v", err)
	}

	if err := helpers.OpenDB(*dbPath); err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer helpers.CloseDB()

	if *dbPath == "testdb" {
		log.Println("Using test database; starting initialization")
		go helpers.InitTestDB()
	}

	helpers.Serve()
}
