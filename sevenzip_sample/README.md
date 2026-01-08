# 7zip File Encryption Sample

This Go application encrypts files using the 7zip command line tool.

## Prerequisites

### Setup

1. Install 7zip using Homebrew:
   ```bash
   brew install p7zip
   ```

2. Install Docker Desktop for Mac from [https://www.docker.com/products/docker-desktop/](https://www.docker.com/products/docker-desktop/)

## Usage

### Local Run
```bash
cd sevenzip_sample
go run main.go <input_file> <password>
```

Example:
```bash
go run main.go test.txt mypassword
7z x test.txt.7z -ooutput -pmypassword
```

### Docker

```bash
docker build -t sevenzip_sample --platform=linux/arm64 .
docker run --rm -v $(pwd):/workspace sevenzip_sample /workspace/test.txt mypassword
```

## Files

- `main.go` - Main application code
- `go.mod` - Go module definition
- `Dockerfile` - Multi-stage Docker configuration with p7zip