# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.2.4
发布时间：2026/05/29
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

- 下拉菜单项改为独立 WPS Ribbon 回调，避免 Windows WPS 在菜单项点击时不传控件标识导致动作识别失败。
- Linux 打包脚本中的发布日期改为随版本矩阵注入，避免模板默认值和实际发布版本不一致。

## 修复

- 修复 Windows 环境点击“朗读方式”和“朗读语速”下拉项时可能提示“未识别的文档朗读按钮：未知按钮”的问题。
- 修复关于朗读弹窗内说明文件在部分弹窗来源下可能提示“说明文件加载失败：Failed to fetch”的问题。
- 修复部分说明文档、打包模板和安装脚本残留旧版本号、旧发布日期的问题。

## 交付文件

| 目标 | 文件 |
| --- | --- |
| x86/x64 Windows 10/11 | dist/wps-read-aloud-comate_1.2.4_windows.exe |
| x64 银河麒麟系统 | dist/wps-read-aloud-comate_1.2.4_amd64.deb |
| ARM64 银河麒麟系统 | dist/wps-read-aloud-comate_1.2.4_arm64.deb |
| x64 UOS系统 | dist/cn.wps-read-aloud-comate_1.2.4_amd64.deb |
| ARM64 UOS系统 | dist/cn.wps-read-aloud-comate_1.2.4_arm64.deb |

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对文档 Range、Selection、Page 等接口的支持。
- Windows 版 WPS 是否显示首次安全确认框由 WPS 客户端策略决定；本项目不伪造或关闭 WPS 原生安全确认。
