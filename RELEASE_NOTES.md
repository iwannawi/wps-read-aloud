# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.1.13
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

- Windows 安装包改为按需启动本地朗读服务，不再写入开机或登录自启动项，安装阶段也不再启动 wps-tts-daemon.exe。
- Windows 加载项改为本地注册入口，点击“开始朗读”或“状态检查”时再启动本地服务；停止朗读、朗读结束或状态检查完成后会主动请求服务退出。
- Windows 安装包新增开始菜单卸载入口，并写入控制面板“应用和功能”卸载信息。
- Windows 卸载脚本重写为完整清理流程，覆盖朗读进程、旧版自启动项、旧计划任务、WPS 加载项注册、授权缓存、本项目开始菜单入口、卸载注册表和安装目录。

## 修复

- 修复 Windows 包安装后长期驻留本地朗读服务的问题，避免未使用朗读功能时占用 CPU、内存和后台进程资源。
- 修复旧版安装包可能遗留当前用户 Run 自启动项的问题。
- 修复卸载后可能残留 WPS 加载项配置、开始菜单入口或卸载注册表项的问题。
- 修复文档中仍把 Windows 服务描述为登录自启动的问题。

## 交付文件

| 目标 | 文件 |
| --- | --- |
| x86/x64 Windows 10/11 | dist/wps-read-aloud-comate_1.1.13_windows.exe |
| x64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.13_amd64.deb |
| ARM64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.13_arm64.deb |
| x64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.13_amd64.deb |
| ARM64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.13_arm64.deb |

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对文档 Range、Selection、Page 等接口的支持。
- 朗读过程中不能切换朗读方式和语速，需要停止朗读后再调整。
- 音量由目标操作系统、声卡设备或桌面环境声音设置控制，加载项内不提供音量调节。
- Windows WPS 首次信任第三方加载项时可能显示 WPS 原生安全确认框。安装脚本会保留已允许记录，避免升级时重复触发；首次确认是否出现由 WPS 客户端安全策略决定。
