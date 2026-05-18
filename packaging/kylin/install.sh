#!/usr/bin/env bash
set -euo pipefail

cat >&2 <<'MSG'
该脚本已停用。

原因：
旧脚本会按目录复制本地 engines 内容，容易把目标环境不需要的库或已弃用语音引擎带入安装目录，不符合五套环境分别交付、按平台最小化打包的要求。

请使用对应目标环境的正式安装包：
  Windows x86：wps-read-aloud-xc_<版本>_windows_x86.exe
  银河麒麟 x64：wps-read-aloud-xc_<版本>_kylin_amd64.deb
  银河麒麟 ARM64：wps-read-aloud-xc_<版本>_kylin_arm64.deb
  UOS x64：wps-read-aloud-xc_<版本>_uos_amd64.deb
  UOS ARM64：wps-read-aloud-xc_<版本>_uos_arm64.deb
MSG

exit 1
