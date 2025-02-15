# grpc_sample

## Tips

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    greet/greet.proto

go run ./server
go run ./client
go run ./sample

grpcurl -plaintext -d '{"name": "Alice"}' localhost:50051 greet.GreetService/SayHello
```

## Setup

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Memo

```bash
go mod init example.com/aaa
go mod tidy
go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get google.golang.org/grpc
```
