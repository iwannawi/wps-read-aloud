# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.1.6
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

- Windows 播放链路改为直接调用系统 WinMM 接口播放 WAV，不再为每句朗读启动 PowerShell 播放进程。
- Windows 启动朗读时只等待第一句合成完成，后续句子在后台继续预合成，避免长文档被启动预热阶段拖慢或提前失败。
- Windows 语音合成增加并发闸门，预合成窗口最多 4 句，同时运行的 Sherpa-onnx 合成进程最多 2 个，单个合成进程最多使用 4 个 CPU 线程。
- 中英文与符号预处理增加控制字符、私有区字符、表情符号等清洗逻辑，降低长文档中个别特殊字符导致合成失败的概率。
- Windows 朗读失败提示改为包含具体句号和底层错误摘要，不再把短文档可用场景误报为安装包完整性问题。

## 修复

- 修复 Windows 长文档朗读时，因启动阶段等待多句预合成和逐句 PowerShell 播放带来的进程开销过高，导致朗读一开始就失败的问题。
- 修复 Windows 环境启动朗读等待时间明显长于 Linux 的问题，减少每句播放前后的额外进程启动成本。
- 修复长文档中包含未映射特殊符号时，语音合成错误提示不准确的问题。

## 交付文件

| 目标 | 文件 |
| --- | --- |
| x86/x64 Windows 10/11 | dist/wps-read-aloud-comate_1.1.6_windows.exe |
| x64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.6_amd64.deb |
| ARM64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.6_arm64.deb |
| x64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.6_amd64.deb |
| ARM64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.6_arm64.deb |

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对文档 Range、Selection、Page 等接口的支持。
- 朗读过程中不能切换朗读方式和语速，需要停止朗读后再调整。
- 音量由目标操作系统、声卡设备或桌面环境声音设置控制，加载项内不提供音量调节。
