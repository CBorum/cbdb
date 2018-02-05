package main

import (
	"github.com/cborum/cbdb"
	"log"
	"os"
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

	log.Println("value:", value)
	log.Println(db.CurrentOffset)
	log.Println(db.HashIndex)
}
