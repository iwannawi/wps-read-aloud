# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.1.5
发布时间：20260522
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

- 长文档朗读上限从 1000 句提升到 20000 句，并将朗读请求体上限提升到 64 MB。
- 长文档总文本保护上限调整为 200 万字符，单句仍保留 1000 字保护，避免异常长句拖垮语音合成。
- 预合成策略保留累计约 100 字的启动缓冲，同时增加每轮最多 6 句的并发上限。
- Windows 安装对话框调整为 900×660 客户区，在 1024×768 屏幕下可完整显示；安装路径、进度和执行信息区域同步加宽加高。

## 修复

- 修复短句较多的长文档启动朗读时，预合成并发过高导致 Sherpa-onnx 启动失败的问题。
- 修复 1000 句以上文档被前端或后端截断后无法按预期朗读的问题。
- 修复 Windows 安装对话框过小，安装路径和执行信息显示不完整的问题。
- 修复构建目录中旧 GitHub Release 日志残留时，全量发布包校验失败的问题。

## 交付文件

| 目标 | 文件 |
| --- | --- |
| x86/x64 Windows 10/11 | dist/wps-read-aloud-comate_1.1.5_windows.exe |
| x64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.5_amd64.deb |
| ARM64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.5_arm64.deb |
| x64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.5_amd64.deb |
| ARM64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.5_arm64.deb |

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对文档 Range、Selection、Page 等接口的支持。
- 朗读过程中不能切换朗读方式和语速，需要停止朗读后再调整。
- 音量由目标操作系统、声卡设备或桌面环境声音设置控制，加载项内不提供音量调节。
