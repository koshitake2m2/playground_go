# 7zip File Encryption Sample

This Go application encrypts files using the 7zip command line tool.

## Setup

The 7zip binaries are already included in the repository:
- `bin/7zz` - macOS binary
- `bin/linux/7zz` - Linux ARM64 binary

## Usage

### Local Run (macOS/Linux)

```bash
go run main.go <input_file> <password>
```

The application automatically detects the operating system and uses the appropriate binary.

### Environment Variable Override

You can override the OS detection using the `TARGET_OS` environment variable:

```bash
# Force use of Linux binary (useful for testing)
TARGET_OS=linux go run main.go test.txt mypassword

# Force use of macOS binary
TARGET_OS=darwin go run main.go test.txt mypassword
```

### Examples

```bash
# Encrypt a file
go run main.go test.txt mypassword

# Decrypt the encrypted file
./bin/7zz x test.txt.7z -pmypassword
```

### Container Build Options

You can choose between two build methods using the `BUILD_METHOD` environment variable:

#### Method 1: ko build (default) - Uses bundled binaries

```bash
# Install ko
go install github.com/google/ko@latest

# Build with ko (uses bundled 7z binaries)
BUILD_METHOD=ko ./build.sh

# Or directly:
IMAGE=$(ko build --local --platform=linux/arm64 .)
docker run --rm -v $(pwd):/workspace $IMAGE /workspace/test.txt mypassword123

# Clean, build and run in one command
./clean.sh && IMAGE=$(ko build --local --platform=linux/arm64 .) && docker run --rm -v $(pwd):/workspace $IMAGE /workspace/test.txt mypassword123
```

#### Method 2: Docker build - Uses system 7z command

```bash
# Build with Dockerfile (installs p7zip package)
BUILD_METHOD=docker ./build.sh

# Or directly:
docker build -t sevenzip_sample --platform=linux/arm64 .
docker run --rm -v $(pwd):/workspace sevenzip_sample /workspace/test.txt mypassword123
```

### Environment Variables

- `BUILD_METHOD`: Choose build method (`ko` or `docker`)
- `TARGET_OS`: Override OS detection (`darwin` or `linux`)
- `USE_SYSTEM_7Z`: Use system 7z command instead of bundled binaries (set to `1`)

### Cleanup

To remove generated 7z files:

```bash
./clean.sh
```

## Files

- `main.go` - Main application code
- `go.mod` - Go module definition
- `bin/7zz` - macOS 7zip binary
- `bin/linux/7zz` - Linux ARM64 7zip binary
- `clean.sh` - Cleanup script for generated files