consolidated


## Init
```bash
go test
```

## Run
```bash
gomon . or gomon ro go run main.go
```

### 1 - Download dependencies via vendor folder

- run `go mod vendor`

- the vendor folder should have new folders in it representing dependencies

### 2 - Back to offline

- copy `go.mod`, `go.sum` and `vendor`
- paste them on your offline PC, edit `go.mod` module name if necessary
- run your go commands with the flag `-mod=vendor` like `go run -mod=vendor main.go`

## Using git