# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.1.16
发布时间：20260525
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

- Windows 加载项改为 OEM 指向 jsplugins.xml 的 publish 离线模式，安装时写入本地“文档朗读助手_版本号”目录，不再依赖 19860 服务加载选项卡。
- Windows 安装器会请求管理员权限，用于修改 WPS 安装目录 office6/cfgs/oem.ini，并提示安装完成后重启电脑。
- Windows 开始菜单继续使用“WPS文档朗读助手”文件夹，并在文件夹内提供“卸载 WPS文档朗读助手”入口。

## 修复

- 修复 WPS 打开时本地语音服务未启动导致加载项选项卡无法显示的问题。
- 修复开始菜单卸载入口名称与指定菜单层级不完全一致的问题。

## 交付文件

| 目标 | 文件 |
| --- | --- |
| x86/x64 Windows 10/11 | dist/wps-read-aloud-comate_1.1.16_windows.exe |
| x64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.16_amd64.deb |
| ARM64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.16_arm64.deb |
| x64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.16_amd64.deb |
| ARM64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.16_arm64.deb |

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对文档 Range、Selection、Page 等接口的支持。
- Windows 版 WPS 是否显示首次安全确认框由 WPS 客户端策略决定；本项目不伪造或关闭 WPS 原生安全确认。

