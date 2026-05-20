# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
Debian 包名：x64/ARM64 银河麒麟 V10 及以上为 “wps-read-aloud-comate”，x64/ARM64 UOS V20 为 “cn.wps-read-aloud-comate”
版本：1.0.38
架构和系统：x86/x64 Windows 10/11、x64/ARM64 银河麒麟 V10 及以上、x64/ARM64 UOS V20
适用办公软件：x86/x64 Windows 10/11 要求 WPS Office 2019 或更高版本，推荐 WPS Office 最新稳定版；x64/ARM64 银河麒麟 V10 及以上、x64/ARM64 UOS V20 要求 WPS Office 2019 或更高版本，推荐最新版 WPS Office for Linux。
开发者：Zhang Jingyao
发布时间：20260520

## 本版本变更

- 统一所有说明、弹窗和打包元数据中的平台表述，固定使用“CPU 架构 + 操作系统名”的顺序。
- 明确支持范围为 x86/x64 Windows 10/11、x64/ARM64 银河麒麟 V10 及以上、x64/ARM64 UOS V20。
- 更新关于朗读弹窗、验收说明、源码获取说明、第三方组件声明和多平台打包说明中的适用环境描述。
- 更新 Debian 控制信息，使安装包描述按 CPU 架构和操作系统写入。
- 全部交付包版本更新为 1.0.38。

## 交付文件

    dist/wps-read-aloud-comate_1.0.38_windows.exe
    dist/wps-read-aloud-comate_1.0.38_amd64.deb
    dist/wps-read-aloud-comate_1.0.38_arm64.deb
    dist/cn.wps-read-aloud-comate_1.0.38_amd64.deb
    dist/cn.wps-read-aloud-comate_1.0.38_arm64.deb

## 安装提示

x86/x64 Windows 10/11：

    运行 dist/wps-read-aloud-comate_1.0.38_windows.exe

x64 银河麒麟 V10 及以上：

    sudo dpkg -i dist/wps-read-aloud-comate_1.0.38_amd64.deb

ARM64 银河麒麟 V10 及以上：

    sudo dpkg -i dist/wps-read-aloud-comate_1.0.38_arm64.deb

x64 UOS V20：

    sudo dpkg -i dist/cn.wps-read-aloud-comate_1.0.38_amd64.deb

ARM64 UOS V20：

    sudo dpkg -i dist/cn.wps-read-aloud-comate_1.0.38_arm64.deb

如果 WPS 已经打开，请安装完成后重启 WPS，再使用顶部“文档朗读”选项卡。

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对 Range、Selection、Page 等接口的支持；如果目标 WPS 版本接口行为不同，当页朗读会尽量回退到可读范围。
- 朗读过程中不能切换朗读方式和语速，需要停止朗读后再调整。
- 音量由目标操作系统、声卡设备或桌面环境声音设置控制，加载项内不提供音量调节。
