# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.1.10
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

- Windows 安装界面去掉安装路径展示区域，将原区域用于安装步骤和日志信息展示。
- Windows 安装界面继续保持无窗体滚动条，安装详情区域高度增加。
- Windows 朗读播放改为 WinMM 异步播放并按 WAV 时长等待，停止请求可直接中断当前声音。
- 前端切句不再把图片或对象占位符作为句子边界，避免图片位置产生额外停顿。

## 修复

- 修复点击“停止朗读”后 Windows 端仍需等待当前句播放结束的问题。
- 修复图片或嵌入对象附近可能生成轻音、停顿或选中图片对象的问题。
- 修复图片段落清理后只剩标点时仍可能生成朗读任务的问题。

## 交付文件

| 目标 | 文件 |
| --- | --- |
| x86/x64 Windows 10/11 | dist/wps-read-aloud-comate_1.1.10_windows.exe |
| x64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.10_amd64.deb |
| ARM64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.10_arm64.deb |
| x64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.10_amd64.deb |
| ARM64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.10_arm64.deb |

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对文档 Range、Selection、Page 等接口的支持。
- 朗读过程中不能切换朗读方式和语速，需要停止朗读后再调整。
- 音量由目标操作系统、声卡设备或桌面环境声音设置控制，加载项内不提供音量调节。
