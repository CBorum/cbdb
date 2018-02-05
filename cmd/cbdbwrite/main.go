package main

import (
	"fmt"
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

	fmt.Println("written", key, value)
}
