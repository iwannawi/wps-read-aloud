# 第三方源码获取说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.0.39
开发者：Zhang Jingyao
发布时间：20260520

本软件包内置第三方开源组件，以便在 x86/x64 Windows 10/11、x64/ARM64 银河麒麟 V10 及以上、x64/ARM64 UOS V20 环境中离线运行。x64/ARM64 银河麒麟 V10 及以上安装包安装在 “/opt/wps-read-aloud-comate”，x64/ARM64 UOS V20 安装包安装在 “/opt/apps/cn.wps-read-aloud-comate/files”，x86/x64 Windows 10/11 安装包安装在用户选择的程序目录。

## Sherpa-onnx

- 许可证：Apache License 2.0
- 源码地址：https://github.com/k2-fsa/sherpa-onnx
- x64/ARM64 银河麒麟 V10 及以上、x64/ARM64 UOS V20 包内路径：“engines/sherpa-onnx”

## ONNX Runtime

- 许可证：MIT
- 源码地址：https://github.com/microsoft/onnxruntime
- x64/ARM64 银河麒麟 V10 及以上、x64/ARM64 UOS V20 包内路径：“engines/sherpa-onnx/lib/libonnxruntime*”

## Sherpa-onnx VITS 语音模型

- 中文模型：“vits-zh-hf-fanchen-C”
- 上游模型发布页：https://github.com/k2-fsa/sherpa-onnx/releases/tag/tts-models
- Hugging Face 镜像页：https://huggingface.co/csukuangfj/vits-zh-hf-fanchen-C
- 包内路径：“voices/sherpa/vits-zh-hf-fanchen-C/”

神经语音模型的再分发可能涉及单位内部合规要求。正式对外分发前，建议完成模型来源、训练数据和再分发许可复核，必要时取得授权或替换为许可明确的中文模型。
