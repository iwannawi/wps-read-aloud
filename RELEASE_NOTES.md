# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.1.12
发布时间：20260523
开发者：Zhang Jingyao

## 适用环境

| CPU 架构 + 操作系统 | WPS 要求 |
| --- | --- |
| x86/x64 Windows 10/11 | WPS Office 2019 或更高版本，推荐最新稳定版 |
| x64 银河麒麟 V10 及以上 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| ARM64 银河麒麟 V10 及以上 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| x64 UOS V20 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| ARM64 UOS V20 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |

## 变更

- README、验收测试、多平台打包说明、Debian 包说明、版本管理说明和发布脚本同步更新到当前 1.1.12 方案。
- 文档重新区分“共通架构”和“平台差异”，明确 Windows、银河麒麟、UOS 在安装目录、服务启动、播放层、加载项注册、许可弹窗和日志位置上的差异。
- README 的技术方案、朗读能力、安装验证和平台限制说明按现有实现重写，删除容易让人误解为右侧面板、音量调节、Piper/eSpeak 或双模型切换的旧描述。
- 验收说明补充 Windows 原生第三方加载项许可确认框的验证要求，并说明该弹窗由 Windows 版 WPS 客户端安全策略控制。

## 修复

- 修复说明文档中对技术方案、朗读能力和跨平台差异描述不完整的问题。
- 修复部分脚本和文档仍引用上一版本号的问题。
- 修复文档中未充分说明 Windows 和 Linux 平台能力差异的问题。

## 交付文件

| 目标 | 文件 |
| --- | --- |
| x86/x64 Windows 10/11 | dist/wps-read-aloud-comate_1.1.12_windows.exe |
| x64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.12_amd64.deb |
| ARM64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.12_arm64.deb |
| x64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.12_amd64.deb |
| ARM64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.12_arm64.deb |

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对文档 Range、Selection、Page 等接口的支持。
- 朗读过程中不能切换朗读方式和语速，需要停止朗读后再调整。
- 音量由目标操作系统、声卡设备或桌面环境声音设置控制，加载项内不提供音量调节。
- Windows WPS 首次信任第三方加载项时可能显示 WPS 原生安全确认框。安装脚本会保留已允许记录，避免升级时重复触发；首次确认是否出现由 WPS 客户端安全策略决定。
