# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.2.1
发布时间：2026/05/26
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

- 修复“关于朗读”弹窗中“发布说明”和“第三方声明”无法打开，提示“说明文件加载失败：Failed to fetch”的问题。
- 保持原有说明文件打开方式不变，仅将带说明文件链接的弹窗恢复为本机服务端弹窗页面，避免 WPS 内置浏览器从本地弹窗跨源读取 127.0.0.1 文档时被拦截。

## 交付文件

| 目标 | 文件 |
| --- | --- |
| x86/x64 Windows 10/11 | dist/wps-read-aloud-comate_1.2.1_windows.exe |
| x64 银河麒麟系统 | dist/wps-read-aloud-comate_1.2.1_amd64.deb |
| ARM64 银河麒麟系统 | dist/wps-read-aloud-comate_1.2.1_arm64.deb |
| x64 UOS系统 | dist/cn.wps-read-aloud-comate_1.2.1_amd64.deb |
| ARM64 UOS系统 | dist/cn.wps-read-aloud-comate_1.2.1_arm64.deb |

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对文档 Range、Selection、Page 等接口的支持。
- Windows 版 WPS 是否显示首次安全确认框由 WPS 客户端策略决定；本项目不伪造或关闭 WPS 原生安全确认。
