FROM alpine:latest

RUN apk add --no-cache p7zip

COPY sevenzip_sample /app/sevenzip_sample

WORKDIR /app

ENTRYPOINT ["/app/sevenzip_sample"]