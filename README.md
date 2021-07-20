consolidated

## 
```bash
gomon . || gomon || go run main.go

go mod download 

go mod tidy

go mod vendor

go build -mod=vendor || GOFLAGS="-mod=vendor" go build

```
