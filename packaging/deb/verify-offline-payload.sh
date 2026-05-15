#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"

required=(
  "dist/wps-tts-daemon"
  "engines/piper/piper"
  "engines/piper/lib"
  "engines/espeak-ng/espeak-ng"
  "engines/espeak-ng/espeak-ng-data"
  "engines/espeak-ng/lib"
  "voices/zh_CN.onnx"
  "voices/zh_CN.onnx.json"
)

missing=0
for item in "${required[@]}"; do
  if [[ ! -e "${ROOT_DIR}/${item}" ]]; then
    echo "MISSING ${item}"
    missing=1
  else
    echo "OK      ${item}"
  fi
done

if [[ "${missing}" -ne 0 ]]; then
  echo "offline payload is incomplete; refusing to build a clean-system package." >&2
  exit 1
fi

echo "offline payload check passed."
