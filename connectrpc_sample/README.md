# connectrpc_sample

## Tips

```bash
buf lint
buf generate

go run ./server

grpcurl \
    -protoset <(buf build -o -) -plaintext \
    -d '{"name": "Alice"}' \
    localhost:50051 greet.GreetService/SayHello

grpcurl \
    -protoset <(buf build -o -) -plaintext \
    -d '{"filename": "hello.pdf"}' \
    localhost:50051 pdf.PdfService/GetPdf | jq -r '.content' | base64 -d > ~/Downloads/output.pdf
```

## Setup

```bash
go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
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
