# 7zip File Encryption Sample

This Go application encrypts files using the 7zip command line tool.

## Prerequisites

### Mac Setup

1. Install 7zip using Homebrew:
   ```bash
   brew install p7zip
   ```

2. Install ko.build:
   ```bash
   go install github.com/google/ko@latest
   ```

3. Install Docker Desktop for Mac from [https://www.docker.com/products/docker-desktop/](https://www.docker.com/products/docker-desktop/)

## Usage

### Local Run
```bash
cd sevenzip_sample
go run main.go <input_file> <password>
```

Example:
```bash
echo "test content" > test.txt
go run main.go test.txt mypassword
7z x test.txt.7z -ooutput -pmypassword
```

### Docker with ko.build

1. Build and run with ko:
   ```bash
   cd sevenzip_sample
   ko build --local --platform=linux/amd64 .
   ```

2. Run the container:
   ```bash
   docker run --rm -v $(pwd):/workspace sevenzip_sample:latest /workspace/test.txt mypassword
   ```

## Files

- `main.go` - Main application code
- `go.mod` - Go module definition
- `Dockerfile` - Docker configuration
- `.ko.yaml` - ko.build configuration