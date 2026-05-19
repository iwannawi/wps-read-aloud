# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
Debian 包名：银河麒麟为 “wps-read-aloud-comate”，UOS 为 “cn.wps-read-aloud-comate”
版本：1.0.34
架构：x86/x64 Windows、Linux amd64、Linux arm64
适用操作系统：x86/x64 Windows、银河麒麟 x64/ARM64、UOS x64/ARM64，以及兼容 WPS JS 加载项和本地离线服务的同类系统
适用办公软件：Windows 平台要求 WPS Office 2019 或更高版本，推荐 WPS Office 最新稳定版；Linux 平台要求 WPS Office 2019 或更高版本，推荐最新版 WPS Office for Linux。
开发者：Zhang Jingyao
发布时间：20260519

## 本版本变更

- WPS 加载项授权显示名改为“文档朗读助手”，授权说明改为“WPS文档朗读助手加载项申请访问本机语音合成服务”，避免在用户提示中暴露内部包名。
- Windows 安装程序增加图形化安装界面，显示安装路径、安装进度、当前动作、安装结果、失败原因和安装完成后的操作建议。
- Windows 加载项注册改为“在线入口加本地入口”双注册：publish.xml 用于让 WPS 发现加载项，jsplugins.xml 作为本地文件兜底入口。
- Windows 安装阶段会启动并等待本机语音合成服务健康检查通过，同时清理本项目旧名称的 WPS 授权缓存和 JS 加载项阻止缓存，降低升级后看不到“文档朗读”选项卡的概率。
- Linux 注册脚本同步使用中文加载项显示名，并兼容清理旧版本使用的内部名称。
- 全部交付包版本更新为 1.0.34，并补充 Windows 图形安装器文件的发布产物校验。

## 交付文件

    dist/wps-read-aloud-comate_1.0.34_windows.exe
    dist/wps-read-aloud-comate_1.0.34_amd64.deb
    dist/wps-read-aloud-comate_1.0.34_arm64.deb
    dist/cn.wps-read-aloud-comate_1.0.34_amd64.deb
    dist/cn.wps-read-aloud-comate_1.0.34_arm64.deb

## 安装提示

x86/x64 Windows 环境：

    运行 dist/wps-read-aloud-comate_1.0.34_windows.exe

银河麒麟 x64 环境：

    sudo dpkg -i dist/wps-read-aloud-comate_1.0.34_amd64.deb

银河麒麟 ARM64 环境：

    sudo dpkg -i dist/wps-read-aloud-comate_1.0.34_arm64.deb

UOS x64 环境：

    sudo dpkg -i dist/cn.wps-read-aloud-comate_1.0.34_amd64.deb

UOS ARM64 环境：

    sudo dpkg -i dist/cn.wps-read-aloud-comate_1.0.34_arm64.deb

如果 WPS 已经打开，请安装完成后重启 WPS，再使用顶部“文档朗读”选项卡。

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对 Range、Selection、Page 等接口的支持；如果目标 WPS 版本接口行为不同，当页朗读会尽量回退到可读范围。
- 朗读过程中不能切换朗读方式和语速，需要停止朗读后再调整。
- 音量由目标操作系统、声卡设备或桌面环境声音设置控制，加载项内不提供音量调节。
