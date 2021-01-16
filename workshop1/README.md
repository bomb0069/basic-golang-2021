# Basic Golang

## Test

```cmd
% go test ./...
ok  	demo	(cached)    ### จะรันใหม่ก็ต่อเมื่อ Production Code เปลี่ยน
?   	demo/cmd	[no test files]
```

go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

## Benchmask

```cmd
go test -bench .
```

## Alias TYPE

```go
    type MONTH int
```
