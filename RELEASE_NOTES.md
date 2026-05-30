# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.2.5
发布时间：2026/05/30
开发者：ZHANG JING YAO

## 适用环境

| CPU 架构 + 操作系统 | WPS 要求 |
| --- | --- |
| x86/x64 Windows 10/11 | WPS Office 2019 或更高版本，推荐最新稳定版 |
| x64 银河麒麟系统 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| ARM64 银河麒麟系统 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| x64 UOS系统 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| ARM64 UOS系统 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |

## 变更

- 将文档朗读选项卡 6 个功能图标更新为深浅主题通用样式。
- 图标保留黑色主体并增加浅色外描边，提升 WPS 黑色主题下的可见性。

## 修复

- 修复 WPS 切换为黑色主题后，纯黑 PNG 图标与深色功能区背景对比不足的问题。

## 交付文件

| 目标 | 文件 |
| --- | --- |
| x86/x64 Windows 10/11 | dist/wps-read-aloud-comate_1.2.5_windows.exe |
| x64 银河麒麟系统 | dist/wps-read-aloud-comate_1.2.5_amd64.deb |
| ARM64 银河麒麟系统 | dist/wps-read-aloud-comate_1.2.5_arm64.deb |
| x64 UOS系统 | dist/cn.wps-read-aloud-comate_1.2.5_amd64.deb |
| ARM64 UOS系统 | dist/cn.wps-read-aloud-comate_1.2.5_arm64.deb |

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对文档 Range、Selection、Page 等接口的支持。
- Windows 版 WPS 是否显示首次安全确认框由 WPS 客户端策略决定；本项目不伪造或关闭 WPS 原生安全确认。
