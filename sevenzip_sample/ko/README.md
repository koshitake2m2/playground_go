# Ko build

## Setup

```bash
cd sevenzip_sample/ko
mkdir -p kodata/bin
curl -L https://github.com/ip7z/7zip/releases/download/25.01/7z2501-linux-arm64.tar.xz -o kodata/bin/7z2501-linux-arm64.tar.xz
cd kodata/bin
tar -xf 7z2501-linux-arm64.tar.xz
```

`kodata` directory is static assets for ko.

## Run

```bash
IMAGE=$(ko build --local --platform=linux/arm64 .) && docker run --rm -v $(pwd):/workspace $IMAGE /workspace/test.txt mypassword

7zz x test.txt.7z -pmypassword -ooutput
```
