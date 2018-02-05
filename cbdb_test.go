package cbdb

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	defer removeDB()
	m.Run()
}

func TestWrite(t *testing.T) {
	db := NewDB()
	db.Write("abc", "123")
	log.Println(db.HashIndex)
	log.Println(db.HashIndex["abc"])
	log.Println(db.CurrentOffset)
}

func TestWriteTwice(t *testing.T) {
	db := NewDB()
	err := db.Write("abc", "123")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	err = db.Write("mango", "abcdefghijklmnpqrstuvwxyzæøå")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	log.Println(db.HashIndex)
	log.Println(db.HashIndex["abc"])
	log.Println(db.HashIndex["mango"])
	log.Println(db.CurrentOffset)
}

func TestRead(t *testing.T) {
	db := NewDB()
	db.Write("abc", "123")
	val, err := db.Read("abc")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	log.Println(val)
	if val != "123" {
		log.Println("value doesn't match original value.", val)
		t.Fail()
	}

	log.Println(db.HashIndex)
	log.Println(db.CurrentOffset)
}

func TestReadTwice(t *testing.T) {
	db := NewDB()
	err := db.Write("abc", "123")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	err = db.Write("mango", "abcdefghijklmnpqrstuvwxyzæøå")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	val, err := db.Read("abc")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	log.Println(val)
	if val != "123" {
		log.Println("value doesn't match original value.", val)
		t.Fail()
	}

	val, err = db.Read("mango")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	log.Println(val)
	if val != "abcdefghijklmnpqrstuvwxyzæøå" {
		log.Println("value doesn't match original value.", val)
		t.Fail()
	}

	log.Println(db.HashIndex)
	log.Println(db.CurrentOffset)
}

func TestReadNoKey(t *testing.T) {
	db := NewDB()
	db.Write("abc", "123")
	val, err := db.Read("def")
	if err == nil {
		log.Println(err)
		t.FailNow()
	}
	log.Println(val)
	log.Println(db.HashIndex)
	log.Println(db.CurrentOffset)
}

func TestPersist(t *testing.T) {
	db := NewDB()
	db.Write("test", "1234")
	log.Println(db.CurrentOffset, db.HashIndex)

	err := db.Persist()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	log.Println("👍")
}

func TestLoad(t *testing.T) {
	db := NewDB()
	db.Write("test", "1234")
	db.Write("mango", "xddd")
	log.Println(db.CurrentOffset, db.HashIndex)

	err := db.Persist()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	db2 := NewDB()
	err = db2.Load()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	log.Println(db.CurrentOffset, db.HashIndex)
}

func TestLogWrite(t *testing.T) {
	db := NewDB()
	n, err := db.LogWrite("123")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	log.Println("bytes written", n)
	log.Println(db.CurrentOffset)
}

func TestLogRead(t *testing.T) {
	db := NewDB()
	n, err := db.LogWrite("123")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	log.Println("bytes written", n)
	val, err := db.LogRead(0, n)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	log.Println("value", string(val))
	log.Println(db.CurrentOffset)
}

func BenchmarkWrite(b *testing.B) {
	db := NewDB()
	for i := 0; i < b.N; i++ {
		value := RandStringBytesMaskImprSrc(32)
		err := db.Write(string(i), value)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkRead(b *testing.B) {
	db := NewDB()
	for i := 0; i < b.N; i++ {
		value := RandStringBytesMaskImprSrc(32)
		key := RandStringBytesMaskImprSrc(32)
		err := db.Write(key, value)
		if err != nil {
			panic(err)
		}

		readVal, err := db.Read(key)
		if err != nil {
			panic(err)
		}

		if readVal != value {
			panic("readVal != value")
		}
	}
}

// source: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// end
