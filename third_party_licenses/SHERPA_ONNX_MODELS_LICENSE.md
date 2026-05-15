# Sherpa-onnx 语音模型许可说明

本安装包内置以下 Sherpa-onnx 官方发布的离线 TTS 模型文件：

- `matcha-icefall-zh-baker`
- `matcha-icefall-en_US-ljspeech`
- `vocos-22khz-univ.onnx`

来源：

- https://github.com/k2-fsa/sherpa-onnx/releases/tag/tts-models
- https://github.com/k2-fsa/sherpa-onnx/releases/tag/vocoder-models

Sherpa-onnx 项目采用 Apache License 2.0。上述模型随 Sherpa-onnx 官方 release 分发，但模型文件和训练数据集可能具有独立许可或使用限制，不能直接等同于 Sherpa-onnx 主项目许可证。

## 重要合规提示

- `matcha-icefall-zh-baker` 上游模型说明中写明：训练数据集来自 Data Baker 免费数据集，且该数据集为非商业用途。正式政企或商业交付前，必须取得相应商业授权，或更换为许可允许商业分发和使用的中文模型。
- `matcha-icefall-en_US-ljspeech` 使用 LJSpeech 数据集训练，正式交付前仍建议按照单位合规流程复核模型、数据集和再分发许可。
- `vocos-22khz-univ.onnx` 为独立声码器模型，正式交付前也应保留上游来源与许可说明。

因此，包含 `matcha-icefall-zh-baker` 的当前安装包适合技术验证、内部功能测试和音质评估；若用于正式企业交付，应先完成模型授权确认或替换模型。

本项目未修改上述模型文件，仅作为离线运行资源随安装包一起分发到 `/opt/wps-read-aloud/voices/sherpa`。
