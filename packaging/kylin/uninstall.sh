#!/usr/bin/env bash
set -euo pipefail

systemctl --user disable --now wps-tts.service 2>/dev/null || true
rm -f "${HOME}/.config/systemd/user/wps-tts.service"
systemctl --user daemon-reload 2>/dev/null || true

echo "已移除用户服务。"
echo "如需删除程序文件：sudo rm -rf /opt/wps-read-aloud /etc/wps-read-aloud"
