# Build stage
FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o sevenzip_sample .

# Final stage
FROM alpine:latest

RUN apk add --no-cache p7zip

COPY --from=builder /app/sevenzip_sample /app/sevenzip_sample

WORKDIR /app

# Set environment variable to use system 7z command
ENV USE_SYSTEM_7Z=1

ENTRYPOINT ["/app/sevenzip_sample"]