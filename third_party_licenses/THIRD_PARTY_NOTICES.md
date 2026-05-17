# 第三方组件声明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-xc
开发者：Zhang Jingyao
发布时间：20260518

本软件包内置第三方离线语音运行组件，便于在纯净 ARM64 麒麟操作系统上安装后直接使用。

## 内置组件

| 组件 | 包内路径 | 许可证 | 来源 |
| --- | --- | --- | --- |
| Sherpa-onnx | `/opt/wps-read-aloud/engines/sherpa-onnx/*` | Apache License 2.0 | https://github.com/k2-fsa/sherpa-onnx |
| ONNX Runtime | `/opt/wps-read-aloud/engines/sherpa-onnx/lib/libonnxruntime*` | MIT | https://github.com/microsoft/onnxruntime |
| 中文 VITS 模型 | `/opt/wps-read-aloud/voices/sherpa/vits-zh-hf-fanchen-C/*` | 上游模型许可未明确，正式商用需授权确认或替换模型 | https://github.com/k2-fsa/sherpa-onnx/releases/tag/tts-models |

## 隔离说明

第三方动态库不会安装到 `/usr/lib`、`/lib` 等系统公共库目录。服务进程只在启动 Sherpa-onnx 子进程时，为该子进程设置局部环境变量：

- `LD_LIBRARY_PATH`

这样可以避免覆盖或影响 WPS 以及其他软件使用的系统库。

## 企业交付合规提示

当前中文模型 `vits-zh-hf-fanchen-C` 的上游 Hugging Face 仓库未提供完整模型卡和明确许可证。若安装包用于正式政企或商业交付，应在交付前完成商业授权确认，或替换为许可允许商业分发和使用的中文模型。
