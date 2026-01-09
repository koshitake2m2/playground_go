# Build from Dockerfile

## Setup

Build the Docker image:

```bash
docker build -t sevenzip-docker .
```

## Run

```bash
docker run -v $(pwd):/data -w /data sevenzip-docker test.txt mypassword
7zz x test.txt.7z -pmypassword -ooutput
```
