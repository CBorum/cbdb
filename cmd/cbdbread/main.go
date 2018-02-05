package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cborum/cbdb"
)

// Read
func main() {
	if len(os.Args) < 2 {
		log.Println("not enough args")
		return
	}

	db := cbdb.NewDB()
	err := db.Load()
	if err != nil {
		log.Println(err)
	}

	key := os.Args[1]

	value, err := db.Read(key)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("value:", value)
}
