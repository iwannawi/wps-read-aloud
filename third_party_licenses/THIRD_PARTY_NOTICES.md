# 第三方组件声明

软件包：wps-read-aloud-zhangjingyao
开发者：zhangjingyao
发布时间：20260515

本软件包内置第三方运行组件，以便在纯净银河麒麟 V10 ARM64 系统上离线安装和使用。

## 内置组件

| 组件 | 包内路径 | 许可证 | 来源 |
| --- | --- | --- | --- |
| Piper | `/opt/wps-read-aloud/engines/piper/piper` | MIT | https://github.com/rhasspy/piper |
| Piper phonemize/运行库 | `/opt/wps-read-aloud/engines/piper/lib/*` | 见对应组件许可证 | https://github.com/rhasspy/piper-phonemize |
| ONNX Runtime | `/opt/wps-read-aloud/engines/piper/lib/libonnxruntime*` | MIT | https://github.com/microsoft/onnxruntime |
| eSpeak NG | `/opt/wps-read-aloud/engines/espeak-ng/*` | GPL-3.0-or-later | https://github.com/espeak-ng/espeak-ng |
| Piper 中文语音模型 | `/opt/wps-read-aloud/voices/zh_CN.onnx` 和 `.json` | 上游模型卡声明 MIT | https://huggingface.co/rhasspy/piper-voices |

## 隔离说明

第三方动态库不会安装到 `/usr/lib`、`/lib` 等系统公共库目录。服务进程只在启动内置语音引擎时，为子进程设置局部环境变量：

- `LD_LIBRARY_PATH`
- `ESPEAK_DATA_PATH`

这样可以避免覆盖或影响 WPS 以及其他软件使用的系统库。
