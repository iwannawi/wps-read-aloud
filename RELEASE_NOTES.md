# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
Debian 包名：银河麒麟为 “wps-read-aloud-comate”，UOS 为 “cn.wps-read-aloud-comate”
版本：1.0.37
架构：x86/x64 Windows、Linux amd64、Linux arm64
适用操作系统：x86/x64 Windows、银河麒麟 x64/ARM64、UOS x64/ARM64，以及兼容 WPS JS 加载项和本地离线服务的同类系统
适用办公软件：Windows 平台要求 WPS Office 2019 或更高版本，推荐 WPS Office 最新稳定版；Linux 平台要求 WPS Office 2019 或更高版本，推荐最新版 WPS Office for Linux。
开发者：Zhang Jingyao
发布时间：20260520

## 本版本变更

- 修复 Windows WPS 中点击“开始朗读”“状态检查”“关于朗读”后弹出空白窗体的问题。
- 修复 Windows WPS 中功能按钮额外调起外部浏览器打开 “dialog.html” 的问题。
- 弹窗页面地址改为固定使用本机服务绝对地址，避免 WPS 内置弹窗加载相对路径时出现空白页。
- WPS 环境下如果内置弹窗接口异常，不再回退到浏览器窗口，而是使用本地提示兜底，避免重复弹窗。
- 保留上一版本中的 Linux 首次授权提示缓解、启动小窗滚动条修复和同版本重装修复。
- 全部交付包版本更新为 1.0.37。

## 交付文件

    dist/wps-read-aloud-comate_1.0.37_windows.exe
    dist/wps-read-aloud-comate_1.0.37_amd64.deb
    dist/wps-read-aloud-comate_1.0.37_arm64.deb
    dist/cn.wps-read-aloud-comate_1.0.37_amd64.deb
    dist/cn.wps-read-aloud-comate_1.0.37_arm64.deb

## 安装提示

x86/x64 Windows 环境：

    运行 dist/wps-read-aloud-comate_1.0.37_windows.exe

银河麒麟 x64 环境：

    sudo dpkg -i dist/wps-read-aloud-comate_1.0.37_amd64.deb

银河麒麟 ARM64 环境：

    sudo dpkg -i dist/wps-read-aloud-comate_1.0.37_arm64.deb

UOS x64 环境：

    sudo dpkg -i dist/cn.wps-read-aloud-comate_1.0.37_amd64.deb

UOS ARM64 环境：

    sudo dpkg -i dist/cn.wps-read-aloud-comate_1.0.37_arm64.deb

如果 WPS 已经打开，请安装完成后重启 WPS，再使用顶部“文档朗读”选项卡。

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对 Range、Selection、Page 等接口的支持；如果目标 WPS 版本接口行为不同，当页朗读会尽量回退到可读范围。
- 朗读过程中不能切换朗读方式和语速，需要停止朗读后再调整。
- 音量由目标操作系统、声卡设备或桌面环境声音设置控制，加载项内不提供音量调节。
