# Description
My implementation has two binaries that respectively write to and read from the db. 
On writes the key and value is binary encoded and appended to a database file, afterwards recognizing how many bytes was written and saving the offset and length of the content in a hashmap on the key value.
When reading, the key provided is searched for in the in-memory hashmap, returning the offset and how many bytes that should be read.
On every write the Database struct is persisted to an index file, which also contains the currect byte offset.

# Usage
## In Go environment
if you have Go and paths setup correctly, else skip to Docker

```bash
go get github.com/cborum/cbdb/...
```
(yes, with the three dots)
This should install two binaries `cbdbwrite` and `cbdbread` in your `$GOPATH/bin`, that should be accessible in your terminal, and used as followed:
```
$ cbdbwrite abc 123
written abc 123
$ cbdbread abc
value: 123
```

## Docker
```bash
$ git clone https://github.com/CBorum/cbdb.git
$ cd cbdb
```

Write:
```bash
docker run -it --rm -v $(pwd):/go/src/github.com/cborum/cbdb -w /go/src/github.com/cborum/cbdb golang:alpine sh -c "go run cmd/cbdbwrite/main.go abc 123"
```

Read:
```bash
docker run -it --rm -v $(pwd):/go/src/github.com/cborum/cbdb -w /go/src/github.com/cborum/cbdb golang:alpine sh -c "go run cmd/cbdbread/main.go abc"
```
