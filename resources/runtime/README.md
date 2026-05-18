# 多平台运行时资源目录

本目录用于存放不同系统和架构的原生语音引擎文件。源码、WPS 加载项前端和语音模型共用，只有原生二进制和动态库按平台区分。

## 目录约定

    resources/runtime/windows-x86/sherpa-onnx/
    resources/runtime/linux-amd64/sherpa-onnx/
    resources/runtime/linux-arm64/sherpa-onnx/

每个 “sherpa-onnx” 目录至少需要包含：

- sherpa-onnx-offline-tts 或 Windows 下的 sherpa-onnx-offline-tts.exe
- lib 目录或同级依赖库

正式构建只从本目录读取平台运行时资源。旧版 “engines” 目录不作为正式安装包输入，避免把目标环境不需要的库和废弃语音引擎带入安装包。

## 打包规则

- 银河麒麟和 UOS 的 amd64 安装包共用 “linux-amd64” 运行时资源。
- 银河麒麟和 UOS 的 arm64 安装包共用 “linux-arm64” 运行时资源。
- Windows x86 exe 安装程序使用 “windows-x86” 运行时资源。
- 语音模型目录 “voices/sherpa” 不按系统区分。

运行时资源体积较大，默认不进入普通 Git 提交。正式发布时应通过受控制品库或 GitHub Release 附件保存。
