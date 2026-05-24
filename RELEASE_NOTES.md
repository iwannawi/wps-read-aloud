# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.1.14
发布时间：20260524
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

- Windows 加载项入口改为实机验证通过的 WPS HTTP 根地址模式：publish.xml 和 jsplugins.xml 均使用 http://127.0.0.1:19860/addin/。
- Windows 安装阶段会清理旧的 wps-read-aloud_版本号目录和旧的文档朗读助手本地目录，避免多版本残留影响 WPS 扫描。
- Windows 安装完成后会启动一次本地朗读服务，供 WPS 重新打开时读取 ribbon 并显示“文档朗读”选项卡；不写入开机自启动项，空闲 30 分钟后自动退出。
- Windows 开始菜单继续使用“WPS文档朗读助手”文件夹，并在文件夹内提供“卸载 WPS 文档朗读助手”入口。

## 修复

- 修复 Windows WPS 顶部不显示“文档朗读”选项卡的问题。
- 修复 Windows 入口 URL 指向 index.html 导致 WPS 不能稳定加载 ribbon 的问题。

## 交付文件

| 目标 | 文件 |
| --- | --- |
| x86/x64 Windows 10/11 | dist/wps-read-aloud-comate_1.1.14_windows.exe |
| x64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.14_amd64.deb |
| ARM64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.14_arm64.deb |
| x64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.14_amd64.deb |
| ARM64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.14_arm64.deb |

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对文档 Range、Selection、Page 等接口的支持。
- Windows 版 WPS 是否显示首次安全确认框由 WPS 客户端策略决定；本项目不伪造或关闭 WPS 原生安全确认。
