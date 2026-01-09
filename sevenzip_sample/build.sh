#!/bin/bash

# Build script for sevenzip_sample
# Usage: BUILD_METHOD=ko ./build.sh
#        BUILD_METHOD=docker ./build.sh

BUILD_METHOD=${BUILD_METHOD:-ko}

if [ "$BUILD_METHOD" = "ko" ]; then
    echo "Building with ko..."
    IMAGE=$(ko build --local --platform=linux/arm64 .)
    echo "Built image: $IMAGE"
    echo "To run: docker run --rm -v \$(pwd):/workspace $IMAGE /workspace/test.txt mypassword123"
    echo ""
    echo "Export IMAGE for easy use:"
    echo "export IMAGE=$IMAGE"
elif [ "$BUILD_METHOD" = "docker" ]; then
    echo "Building with Docker..."
    docker build -t sevenzip_sample --platform=linux/arm64 .
    echo "Built image: sevenzip_sample"
    echo "To run: docker run --rm -v \$(pwd):/workspace sevenzip_sample /workspace/test.txt mypassword123"
    echo ""
    echo "Export IMAGE for easy use:"
    echo "export IMAGE=sevenzip_sample"
else
    echo "Invalid BUILD_METHOD: $BUILD_METHOD"
    echo "Use BUILD_METHOD=ko or BUILD_METHOD=docker"
    exit 1
fi