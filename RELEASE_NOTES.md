# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
Debian 包名：银河麒麟为 “wps-read-aloud-comate”，UOS 为 “cn.wps-read-aloud-comate”
版本：1.0.36
架构：x86/x64 Windows、Linux amd64、Linux arm64
适用操作系统：x86/x64 Windows、银河麒麟 x64/ARM64、UOS x64/ARM64，以及兼容 WPS JS 加载项和本地离线服务的同类系统
适用办公软件：Windows 平台要求 WPS Office 2019 或更高版本，推荐 WPS Office 最新稳定版；Linux 平台要求 WPS Office 2019 或更高版本，推荐最新版 WPS Office for Linux。
开发者：Zhang Jingyao
发布时间：20260520

## 本版本变更

- 修复 Windows WPS 客户端已弹出“申请访问本机语音合成服务”授权提示，但顶部没有出现“文档朗读”选项卡的问题。
- Windows 在线加载项入口从 “/addin/index.html” 调整为 “/addin/” 目录根，便于 Windows WPS 正确定位 “ribbon.xml”等加载项资源。
- Windows 升级安装时同时清理旧内部名称和当前中文名称的 WPS 授权缓存，避免 Windows WPS 沿用旧的 “/addin/index.html” 地址。
- 保留 Windows 图形安装器可见性修复：安装界面使用 STA 模式启动，隐藏控制台但不隐藏安装窗体；启动失败时会弹出错误提示。
- Linux 加载项改为通过本机服务地址同源运行，减少首次点击朗读时触发 WPS “允许访问 127.0.0.1”跨源授权提示的概率。
- 修复“朗读服务正在启动”小窗在部分 WPS 内置浏览器中出现滚动条的问题。
- Linux 安装脚本允许同版本安装包覆盖重装本项目自己的服务，避免端口已由本项目服务占用时误判为其他程序占用。
- 全部交付包版本更新为 1.0.36。

## 交付文件

    dist/wps-read-aloud-comate_1.0.36_windows.exe
    dist/wps-read-aloud-comate_1.0.36_amd64.deb
    dist/wps-read-aloud-comate_1.0.36_arm64.deb
    dist/cn.wps-read-aloud-comate_1.0.36_amd64.deb
    dist/cn.wps-read-aloud-comate_1.0.36_arm64.deb

## 安装提示

x86/x64 Windows 环境：

    运行 dist/wps-read-aloud-comate_1.0.36_windows.exe

银河麒麟 x64 环境：

    sudo dpkg -i dist/wps-read-aloud-comate_1.0.36_amd64.deb

银河麒麟 ARM64 环境：

    sudo dpkg -i dist/wps-read-aloud-comate_1.0.36_arm64.deb

UOS x64 环境：

    sudo dpkg -i dist/cn.wps-read-aloud-comate_1.0.36_amd64.deb

UOS ARM64 环境：

    sudo dpkg -i dist/cn.wps-read-aloud-comate_1.0.36_arm64.deb

如果 WPS 已经打开，请安装完成后重启 WPS，再使用顶部“文档朗读”选项卡。

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对 Range、Selection、Page 等接口的支持；如果目标 WPS 版本接口行为不同，当页朗读会尽量回退到可读范围。
- 朗读过程中不能切换朗读方式和语速，需要停止朗读后再调整。
- 音量由目标操作系统、声卡设备或桌面环境声音设置控制，加载项内不提供音量调节。
