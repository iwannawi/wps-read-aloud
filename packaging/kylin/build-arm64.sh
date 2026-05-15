#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
mkdir -p "${ROOT_DIR}/dist"

if command -v python3 >/dev/null 2>&1; then
  PYTHON_BIN="python3"
elif command -v python >/dev/null 2>&1; then
  PYTHON_BIN="python"
else
  echo "missing required command: python3" >&2
  exit 1
fi

"${PYTHON_BIN}" "${ROOT_DIR}/packaging/sync_addin_web.py"

cd "${ROOT_DIR}/daemon"
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -buildvcs=false -trimpath -ldflags="-s -w" -o "${ROOT_DIR}/dist/wps-tts-daemon" ./cmd/wps-tts-daemon

echo "已生成 ${ROOT_DIR}/dist/wps-tts-daemon"
