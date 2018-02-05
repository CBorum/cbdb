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
$ docker run -it --rm -v $(pwd):/go/src/github.com/cborum/cbdb -w /go/src/github.com/cborum/cbdb golang:alpine sh -c "go run cmd/cbdbwrite/main.go abc 123; go run cmd/cbdbread/main.go abc"
```