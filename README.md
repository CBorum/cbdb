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
