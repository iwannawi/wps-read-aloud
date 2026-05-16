#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"

required=(
  "dist/wps-tts-daemon"
  "engines/sherpa-onnx/sherpa-onnx-offline-tts"
  "engines/sherpa-onnx/lib"
  "voices/sherpa/vits-zh-hf-fanchen-C/vits-zh-hf-fanchen-C.onnx"
  "voices/sherpa/vits-zh-hf-fanchen-C/lexicon.txt"
  "voices/sherpa/vits-zh-hf-fanchen-C/tokens.txt"
  "voices/sherpa/vits-zh-hf-fanchen-C/phone.fst"
  "voices/sherpa/vits-zh-hf-fanchen-C/date.fst"
  "voices/sherpa/vits-zh-hf-fanchen-C/number.fst"
  "voices/sherpa/vits-zh-hf-fanchen-C/new_heteronym.fst"
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
