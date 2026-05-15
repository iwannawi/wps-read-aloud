#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
mkdir -p "${ROOT_DIR}/dist"
rm -rf "${ROOT_DIR}/daemon/cmd/wps-tts-daemon/web"
mkdir -p "${ROOT_DIR}/daemon/cmd/wps-tts-daemon/web"
cp -a "${ROOT_DIR}/addin/." "${ROOT_DIR}/daemon/cmd/wps-tts-daemon/web/"

cd "${ROOT_DIR}/daemon"
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -trimpath -ldflags="-s -w" -o "${ROOT_DIR}/dist/wps-tts-daemon" ./cmd/wps-tts-daemon

echo "已生成 ${ROOT_DIR}/dist/wps-tts-daemon"
