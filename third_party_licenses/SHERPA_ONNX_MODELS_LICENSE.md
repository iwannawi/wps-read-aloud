# Sherpa-onnx 语音模型许可说明

本安装包内置以下 Sherpa-onnx 官方发布的离线 TTS 模型文件：

- `vits-zh-hf-fanchen-C`

来源：

- https://github.com/k2-fsa/sherpa-onnx/releases/tag/tts-models
- https://huggingface.co/csukuangfj/vits-zh-hf-fanchen-C

Sherpa-onnx 项目采用 Apache License 2.0。语音模型文件和训练数据集可能具有独立许可或使用限制，不能直接等同于 Sherpa-onnx 主项目许可证。

## 重要合规提示

- `vits-zh-hf-fanchen-C` 上游 Hugging Face 仓库目前未提供完整模型卡和明确许可证。
- Sherpa-onnx 官方文档说明该模型转换自 Hugging Face Space `lkz99/tts_model` 下的中文模型资源。
- 正式政企或商业交付前，应完成模型来源、训练数据、模型权重再分发和商用许可复核；如无法取得明确授权，应替换为许可明确的中文模型。

因此，包含 `vits-zh-hf-fanchen-C` 的当前安装包适合技术验证、内部功能测试和音质评估；若用于正式企业交付，应先完成模型授权确认或替换模型。

本项目未修改上述模型文件，仅作为离线运行资源随安装包一起分发到 `/opt/wps-read-aloud/voices/sherpa`。
