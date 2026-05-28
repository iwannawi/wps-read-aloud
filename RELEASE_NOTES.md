# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.2.3
发布时间：2026/05/28
开发者：ZHANG JING YAO

## 适用环境

| CPU 架构 + 操作系统 | WPS 要求 |
| --- | --- |
| x86/x64 Windows 10/11 | WPS Office 2019 或更高版本，推荐最新稳定版 |
| x64 银河麒麟系统 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| ARM64 银河麒麟系统 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| x64 UOS系统 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| ARM64 UOS系统 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |

## 修复

- 加强本地朗读服务安全：CORS 策略从全开放改为校验 Origin，防止恶意网页调用本机朗读接口。
- 统一 Windows 与 Linux 朗读行为：语速下限、英文逐字符读法、标点停顿缩放、错误消息格式等现在两端一致。
- Windows 端补齐运行时语速切换（/read/settings）、暂停和恢复（/pause、/resume）功能。
- 修复启动弹窗 setInterval 未清理导致的资源泄漏。
- 修正 /health 和 /read/status 接口缺少请求方法校验的问题。
- Windows 端 HTTP 服务新增 ReadHeaderTimeout 超时设置。

## 交付文件

| 目标 | 文件 |
| --- | --- |
| x86/x64 Windows 10/11 | dist/wps-read-aloud-comate_1.2.3_windows.exe |
| x64 银河麒麟系统 | dist/wps-read-aloud-comate_1.2.3_amd64.deb |
| ARM64 银河麒麟系统 | dist/wps-read-aloud-comate_1.2.3_arm64.deb |
| x64 UOS系统 | dist/cn.wps-read-aloud-comate_1.2.3_amd64.deb |
| ARM64 UOS系统 | dist/cn.wps-read-aloud-comate_1.2.3_arm64.deb |

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对文档 Range、Selection、Page 等接口的支持。
- Windows 版 WPS 是否显示首次安全确认框由 WPS 客户端策略决定；本项目不伪造或关闭 WPS 原生安全确认。
