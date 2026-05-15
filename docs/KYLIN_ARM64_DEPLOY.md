# 银河麒麟 V10 ARM64 部署清单

## 1. 准备系统依赖

```bash
sudo apt update
sudo apt install -y ca-certificates curl alsa-utils espeak-ng
```

如果目标机没有 Go，可以在其他 Linux 机器交叉编译，把 `dist/wps-tts-daemon` 拷贝到目标机。

## 2. 准备 Piper

目录约定：

```text
engines/piper/piper
voices/zh_CN.onnx
voices/zh_CN.onnx.json
```

把 ARM64 可执行的 Piper 放到 `engines/piper/piper`，并确认：

```bash
chmod +x engines/piper/piper
./engines/piper/piper --help
```

如果没有 Piper，服务会回退到 `espeak-ng`。

## 3. 编译服务

在项目根目录：

```bash
chmod +x packaging/kylin/build-arm64.sh
./packaging/kylin/build-arm64.sh
```

## 4. 安装服务和加载项文件

```bash
chmod +x packaging/kylin/install.sh packaging/kylin/uninstall.sh
./packaging/kylin/install.sh
```

检查服务：

```bash
systemctl --user status wps-tts.service
curl http://127.0.0.1:19860/health
```

测试朗读：

```bash
curl -X POST http://127.0.0.1:19860/speak \
  -H 'Content-Type: application/json' \
  -d '{"text":"你好，这是一段离线朗读测试。","rate":1.0,"volume":80}'
```

## 5. 安装 WPS 加载项

加载项文件安装位置：

```text
/opt/wps-read-aloud/addin
```

在 WPS 2023 for Linux 的加载项管理或调试入口中加载：

```text
/opt/wps-read-aloud/addin/manifest.xml
```

如果目标 WPS 环境要求使用指定加载项目录，请把整个 `addin` 目录复制到对应目录，并保持 `manifest.xml`、`ribbon.xml`、`index.html` 的相对位置不变。

## 6. 验收点

- 打开 WPS 文字后能看到“离线朗读”入口或任务窗格
- 点击“朗读选区”能读取当前选中内容
- 点击“朗读全文”能读取当前文档正文
- 服务接口只监听 `127.0.0.1:19860`
- 断网状态下 Piper 或 eSpeak NG 能正常出声
- “停止”能中断当前播放
- “暂停/继续”能控制正在播放的音频进程
