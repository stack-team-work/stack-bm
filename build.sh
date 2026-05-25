#!/bin/bash

APP_NAME="stack-bm"
BUILD_DIR="build"
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")

mkdir -p ${BUILD_DIR}

echo "Building ${APP_NAME} version ${VERSION}"

PLATFORMS=(
    "linux/amd64"
    "linux/arm64"
    "windows/amd64"
    "darwin/amd64"
    "darwin/arm64"
)

for PLATFORM in "${PLATFORMS[@]}"; do
    IFS="/" read -r GOOS GOARCH <<< "$PLATFORM"
    OUTPUT="${BUILD_DIR}/${APP_NAME}_${GOOS}_${GOARCH}"

    if [ "$GOOS" = "windows" ]; then
        OUTPUT="${OUTPUT}.exe"
    fi

    echo "  Building for ${GOOS}/${GOARCH}..."

    CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="-s -w -X main.version=${VERSION}" -o ${OUTPUT} ./cmd/server

    if [ $? -eq 0 ]; then
        echo "    -> ${OUTPUT}"
    else
        echo "    -> Failed to build for ${GOOS}/${GOARCH}"
    fi
done

echo ""
echo "Build complete. Output in ${BUILD_DIR}/"
ls -lh ${BUILD_DIR}/
