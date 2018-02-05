package main

import (
	"log"
	"os"

	"github.com/cborum/cbdb"
)

// Write
func main() {
	if len(os.Args) < 3 {
		log.Println("not enough args")
		return
	}

	db := cbdb.NewDB()
	err := db.Load()
	if err != nil {
		log.Println(err)
	}

	key, value := os.Args[1], os.Args[2]

	err = db.Write(key, value)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("written", os.Args[0], os.Args[1])
	log.Println(db.CurrentOffset)
	log.Println(db.HashIndex)
}
