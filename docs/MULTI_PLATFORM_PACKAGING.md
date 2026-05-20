# 多平台安装包方案

本项目采用同一套源码、同一套 WPS JS 加载项和同一套语音模型，根据 CPU 架构和操作系统生成不同安装包。所有面向用户的环境表述统一使用“CPU 架构 + 操作系统名”的顺序。

## 安装包分类

| 分类 | CPU 架构 | 操作系统 | 安装包格式 | 文件名示例 |
| --- | --- | --- | --- | --- |
| x86/x64 Windows 10/11 | x86/x64 | Windows 10/11 | exe 安装程序 | wps-read-aloud-comate_1.0.38_windows.exe |
| x64 银河麒麟 V10 及以上 | amd64 | 银河麒麟 V10 及以上 | deb | wps-read-aloud-comate_1.0.38_amd64.deb |
| ARM64 银河麒麟 V10 及以上 | arm64 | 银河麒麟 V10 及以上 | deb | wps-read-aloud-comate_1.0.38_arm64.deb |
| x64 UOS V20 | amd64 | UOS V20 | deb | cn.wps-read-aloud-comate_1.0.38_amd64.deb |
| ARM64 UOS V20 | arm64 | UOS V20 | deb | cn.wps-read-aloud-comate_1.0.38_arm64.deb |

## 共用部分

- addin：WPS JS 加载项前端。
- daemon：本地朗读服务源码。
- voices：语音模型文件。
- third_party_licenses：第三方组件声明。

## 平台差异

- x64/ARM64 银河麒麟 V10 及以上包名为 “wps-read-aloud-comate”，安装路径为 “/opt/wps-read-aloud-comate”。
- x64/ARM64 UOS V20 包名为 “cn.wps-read-aloud-comate”，安装路径为 “/opt/apps/cn.wps-read-aloud-comate/files”。
- x64/ARM64 银河麒麟 V10 及以上、x64/ARM64 UOS V20 使用 systemd 启动本地朗读服务。
- x86/x64 Windows 10/11 使用当前用户登录自启动项启动本地朗读服务。
- 当前 x86/x64 Windows 10/11 安装包内置本地朗读服务；由于加载项通过 127.0.0.1 调用独立服务，不注入 WPS 进程，因此可服务 32 位和 64 位 WPS。安装时仍会记录 WPS 位数，便于排查。
- x64/ARM64 银河麒麟 V10 及以上、x64/ARM64 UOS V20 deb 由 packaging/deb/build_deb.py 生成。
- x86/x64 Windows 10/11 exe 安装程序由 packaging/windows/build_windows_package.py 生成。安装程序内嵌 Windows 专用 payload，运行后解压到临时目录并执行安装逻辑。
- 原生语音引擎和动态库按 CPU 架构和操作系统放入 resources/runtime。

## 构建入口

列出支持目标：

    python packaging/build_all.py --list

构建单个目标示例：

    python packaging/build_all.py --target windows
    python packaging/build_all.py --target kylin-amd64
    python packaging/build_all.py --target kylin-arm64
    python packaging/build_all.py --target uos-amd64
    python packaging/build_all.py --target uos-arm64

构建全部目标：

    python packaging/build_all.py

如果目标 CPU 架构和操作系统缺少 daemon 二进制、sherpa-onnx 运行时或依赖库，脚本会停止并提示缺失路径，避免生成不可用安装包。

## 发布要求

正式出版本时必须同时生成五类安装包及其 SHA256 文件。任何一个目标缺少安装包、校验文件、daemon 二进制、Sherpa ONNX 运行时或模型文件，都不得创建 GitHub Release。

发布前检查入口：

    python packaging/verify_release_artifacts.py

该检查会确认：
- 五个安装包全部存在。
- 五个 SHA256 文件全部存在且内容正确。
- “CHECKSUMS.txt” 与五个安装包完全一致。
- x64/ARM64 银河麒麟 V10 及以上、x64/ARM64 UOS V20 包只包含 systemd 服务、对应架构 daemon、对应架构 Sherpa ONNX 运行时和共用语音模型。
- x86/x64 Windows 10/11 包只包含 Windows 安装脚本、Windows daemon、Windows Sherpa ONNX 运行时和共用语音模型。
- 安装包内不得混入 Piper、eSpeak NG 等已弃用资源。
- 安装包内不得混入其他 CPU 架构或操作系统专用二进制和服务文件。

旧版手工安装脚本已删除。正式交付只能使用按目标环境生成的安装包，不能直接复制本地 “engines” 目录。
