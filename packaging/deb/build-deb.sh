#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
PKG_NAME="wps-read-aloud-zhangjingyao"
VERSION="${VERSION:-1.0.1}"
ARCH="${ARCH:-arm64}"
BUILD_DIR="${ROOT_DIR}/build/deb/${PKG_NAME}_${VERSION}_${ARCH}"
OUT_DIR="${ROOT_DIR}/dist"

require_file() {
  if [[ ! -e "$1" ]]; then
    echo "missing required file: $1" >&2
    exit 1
  fi
}

require_file "${ROOT_DIR}/dist/wps-tts-daemon"
require_file "${ROOT_DIR}/engines/piper/piper"
require_file "${ROOT_DIR}/engines/piper/lib"
require_file "${ROOT_DIR}/engines/espeak-ng/espeak-ng"
require_file "${ROOT_DIR}/engines/espeak-ng/espeak-ng-data"
require_file "${ROOT_DIR}/engines/espeak-ng/lib"
require_file "${ROOT_DIR}/voices/zh_CN.onnx"
require_file "${ROOT_DIR}/voices/zh_CN.onnx.json"
require_file "${ROOT_DIR}/third_party_licenses/THIRD_PARTY_NOTICES.md"
require_file "${ROOT_DIR}/third_party_licenses/PIPER_LICENSE.md"
require_file "${ROOT_DIR}/third_party_licenses/PIPER_VOICES_LICENSE.md"
require_file "${ROOT_DIR}/third_party_licenses/ONNXRUNTIME_LICENSE.txt"
require_file "${ROOT_DIR}/third_party_licenses/ESPEAK_NG_COPYING.txt"
require_file "${ROOT_DIR}/RELEASE_NOTES.md"
require_file "${ROOT_DIR}/ACCEPTANCE_TEST.md"
require_file "${ROOT_DIR}/SOURCE_OFFER.md"
require_file "${ROOT_DIR}/CHECKSUMS.txt"

rm -rf "${BUILD_DIR}"
mkdir -p \
  "${BUILD_DIR}/DEBIAN" \
  "${BUILD_DIR}/opt/wps-read-aloud/daemon" \
  "${BUILD_DIR}/opt/wps-read-aloud/addin" \
  "${BUILD_DIR}/opt/wps-read-aloud/engines" \
  "${BUILD_DIR}/opt/wps-read-aloud/voices" \
  "${BUILD_DIR}/etc/wps-read-aloud" \
  "${BUILD_DIR}/usr/bin" \
  "${BUILD_DIR}/usr/share/doc/wps-read-aloud-zhangjingyao" \
  "${BUILD_DIR}/lib/systemd/system" \
  "${OUT_DIR}"

cp "${ROOT_DIR}/packaging/deb/control" "${BUILD_DIR}/DEBIAN/control"
sed -i "s/^Version:.*/Version: ${VERSION}/" "${BUILD_DIR}/DEBIAN/control"
sed -i "s/^Architecture:.*/Architecture: ${ARCH}/" "${BUILD_DIR}/DEBIAN/control"
cp "${ROOT_DIR}/packaging/deb/postinst" "${BUILD_DIR}/DEBIAN/postinst"
cp "${ROOT_DIR}/packaging/deb/preinst" "${BUILD_DIR}/DEBIAN/preinst"
cp "${ROOT_DIR}/packaging/deb/prerm" "${BUILD_DIR}/DEBIAN/prerm"
cp "${ROOT_DIR}/packaging/deb/postrm" "${BUILD_DIR}/DEBIAN/postrm"
chmod 0755 "${BUILD_DIR}/DEBIAN/preinst" "${BUILD_DIR}/DEBIAN/postinst" "${BUILD_DIR}/DEBIAN/prerm" "${BUILD_DIR}/DEBIAN/postrm"

cp "${ROOT_DIR}/dist/wps-tts-daemon" "${BUILD_DIR}/opt/wps-read-aloud/daemon/wps-tts-daemon"
chmod 0755 "${BUILD_DIR}/opt/wps-read-aloud/daemon/wps-tts-daemon"
cp -a "${ROOT_DIR}/addin/." "${BUILD_DIR}/opt/wps-read-aloud/addin/"
cp "${ROOT_DIR}/daemon/config.example.yaml" "${BUILD_DIR}/etc/wps-read-aloud/config.yaml"
cp "${ROOT_DIR}/packaging/deb/wps-tts.service" "${BUILD_DIR}/lib/systemd/system/wps-tts.service"
cp "${ROOT_DIR}/packaging/deb/wps-read-aloud-register" "${BUILD_DIR}/usr/bin/wps-read-aloud-register"
chmod 0755 "${BUILD_DIR}/usr/bin/wps-read-aloud-register"
cp "${ROOT_DIR}/third_party_licenses/THIRD_PARTY_NOTICES.md" "${BUILD_DIR}/usr/share/doc/wps-read-aloud-zhangjingyao/THIRD_PARTY_NOTICES.md"
cp "${ROOT_DIR}/third_party_licenses/PIPER_LICENSE.md" "${BUILD_DIR}/usr/share/doc/wps-read-aloud-zhangjingyao/PIPER_LICENSE.md"
cp "${ROOT_DIR}/third_party_licenses/PIPER_VOICES_LICENSE.md" "${BUILD_DIR}/usr/share/doc/wps-read-aloud-zhangjingyao/PIPER_VOICES_LICENSE.md"
cp "${ROOT_DIR}/third_party_licenses/ONNXRUNTIME_LICENSE.txt" "${BUILD_DIR}/usr/share/doc/wps-read-aloud-zhangjingyao/ONNXRUNTIME_LICENSE.txt"
cp "${ROOT_DIR}/third_party_licenses/ESPEAK_NG_COPYING.txt" "${BUILD_DIR}/usr/share/doc/wps-read-aloud-zhangjingyao/ESPEAK_NG_COPYING.txt"
cp "${ROOT_DIR}/RELEASE_NOTES.md" "${BUILD_DIR}/usr/share/doc/wps-read-aloud-zhangjingyao/RELEASE_NOTES.md"
cp "${ROOT_DIR}/ACCEPTANCE_TEST.md" "${BUILD_DIR}/usr/share/doc/wps-read-aloud-zhangjingyao/ACCEPTANCE_TEST.md"
cp "${ROOT_DIR}/SOURCE_OFFER.md" "${BUILD_DIR}/usr/share/doc/wps-read-aloud-zhangjingyao/SOURCE_OFFER.md"
cp "${ROOT_DIR}/CHECKSUMS.txt" "${BUILD_DIR}/usr/share/doc/wps-read-aloud-zhangjingyao/CHECKSUMS.txt"

if [[ -d "${ROOT_DIR}/engines" ]]; then
  cp -a "${ROOT_DIR}/engines/." "${BUILD_DIR}/opt/wps-read-aloud/engines/"
fi
if [[ -d "${ROOT_DIR}/voices" ]]; then
  cp -a "${ROOT_DIR}/voices/." "${BUILD_DIR}/opt/wps-read-aloud/voices/"
fi
find "${BUILD_DIR}/opt/wps-read-aloud/engines" -type f -name "piper" -exec chmod 0755 {} \;
find "${BUILD_DIR}/opt/wps-read-aloud/engines" -type f -name "espeak-ng" -exec chmod 0755 {} \;

dpkg-deb --build --root-owner-group "${BUILD_DIR}" "${OUT_DIR}/${PKG_NAME}_${VERSION}_${ARCH}.deb"
echo "created ${OUT_DIR}/${PKG_NAME}_${VERSION}_${ARCH}.deb"
