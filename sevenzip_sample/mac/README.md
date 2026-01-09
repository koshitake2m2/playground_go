# macos local execution

## Setup

```bash
cd sevenzip_sample/mac
mkdir bin
curl -L https://github.com/ip7z/7zip/releases/download/25.01/7z2501-mac.tar.xz -o bin/7z2501-mac.tar.xz
cd bin
tar -xf 7z2501-mac.tar.xz
```

## Run

```bash
go run main.go test.txt mypassword

7zz x test.txt.7z -pmypassword -ooutput
```
