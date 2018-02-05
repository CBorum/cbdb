package cbdb

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	dbFileName      = "db"
	dbIndexFileName = "dbi"
)

// Database ...
type Database struct {
	HashIndex     map[string]*Index
	CurrentOffset int64
}

// Index with offset and byte length of the stored value
type Index struct {
	Offset int64
	Length int
}

// NewDB returns a pointer to a new database struct
func NewDB() *Database {
	return &Database{
		// HashIndex:     map[string]int64{}, //old
		HashIndex: map[string]*Index{},
	}
}

// Write puts the value in the db with the specified key
func (db *Database) Write(key, value string) error {
	bytesWritten, err := db.LogWrite(fmt.Sprintf("%s,%s", key, value))
	if err != nil {
		return err
	}

	db.HashIndex[key] = &Index{
		Offset: db.CurrentOffset,
		Length: bytesWritten,
	}
	db.CurrentOffset += int64(bytesWritten)

	// persists the hash index on every write...
	err = db.Persist()

	return nil
}

// LogWrite writes the value in binary to the database file
func (db *Database) LogWrite(value string) (n int, err error) {
	file, err := os.OpenFile(dbFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err = enc.Encode(value)
	if err != nil {
		return
	}

	n, err = file.WriteAt(buf.Bytes(), db.CurrentOffset)
	if err != nil {
		return
	}
	return
}

// Read returns the value of the key
func (db *Database) Read(key string) (value string, err error) {
	//todo key should not contain ","
	if index, ok := db.HashIndex[key]; ok {
		var read string
		read, err = db.LogRead(index.Offset, index.Length)
		log.Println(read)
		if !strings.Contains(read, ",") {
			err = errors.New("no value found")
			return
		}
		value = strings.SplitN(read, ",", 2)[1]
		value = strings.TrimSuffix(value, "\n")
		return
	}
	err = errors.New("key not found")
	return
}

// LogRead returns the value of the key
func (db *Database) LogRead(offset int64, len int) (val string, err error) {
	file, err := os.OpenFile(dbFileName, os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		return
	}

	readBytes := make([]byte, len)
	n, err := file.ReadAt(readBytes, offset)
	if err != nil {
		return
	}
	log.Println("n", n)

	decVal := bytes.NewBuffer(readBytes)
	dec := gob.NewDecoder(decVal)

	err = dec.Decode(&val)
	if err != nil {
		return
	}

	return
}

// Persist indexes
func (db *Database) Persist() error {
	file, err := os.OpenFile(dbIndexFileName, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(db)
	if err != nil {
		return err
	}

	_, err = file.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// Load indexes
func (db *Database) Load() error {
	file, err := os.OpenFile(dbIndexFileName, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	dec := gob.NewDecoder(file)
	return dec.Decode(&db)
}

func removeDB() {
	err := os.Remove(dbFileName)
	if err != nil {
		log.Println(err)
	}

	err = os.Remove(dbIndexFileName)
	if err != nil {
		log.Println(err)
	}
}
