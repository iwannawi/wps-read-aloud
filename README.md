# WPS 离线朗读加载项

目标环境：

- 银河麒麟 V10 ARM64
- WPS 2023 for Linux 12.1.x
- 允许安装 WPS JS 加载项
- WPS JS API 可读取文字选区和全文
- 本机允许访问 `127.0.0.1`

本项目采用：

- `addin/`：WPS JS 加载项，提供任务窗格、读取选区/全文、朗读控制
- `daemon/`：Go 本地服务，监听 `127.0.0.1:19860`
- `packaging/kylin/`：银河麒麟部署脚本和 systemd 用户服务模板

## 工作方式

```text
WPS 文字 -> JS 加载项 -> http://127.0.0.1:19860 -> Go 服务 -> Piper/eSpeak NG -> 系统音频
```

Piper 是首选引擎，eSpeak NG 是兜底引擎。所有文本只发送到本机回环地址，不访问外网。
朗读时按句合成并播放，加载项会在 WPS 文档中选中当前朗读语句；朗读进入下一句时同步选中下一句，WPS 通常会自动滚动到当前选区。
低配置机器上建议优先使用“朗读选区”；加载项会限制单次句子数量和单句长度，避免超长文档造成长时间等待或资源占用过高。

## 目录

```text
addin/
  manifest.xml
  ribbon.xml
  index.html
  assets/
daemon/
  cmd/wps-tts-daemon/main.go
  config.example.yaml
packaging/kylin/
  install.sh
  uninstall.sh
  wps-tts.service
voices/
  .gitkeep
```

## 在银河麒麟上编译本地服务

安装 Go 1.21+ 后：

```bash
cd daemon
go build -o ../dist/wps-tts-daemon ./cmd/wps-tts-daemon
```

如果在 x86_64 机器上交叉编译 ARM64：

```bash
cd daemon
GOOS=linux GOARCH=arm64 go build -o ../dist/wps-tts-daemon ./cmd/wps-tts-daemon
```

## 安装运行时依赖

推荐：

```bash
sudo apt install -y alsa-utils espeak-ng
```

Piper 请使用目标系统可运行的 ARM64 版本，把可执行文件放到：

```text
/opt/wps-read-aloud/engines/piper/piper
```

把中文模型放到：

```text
/opt/wps-read-aloud/voices/zh_CN.onnx
/opt/wps-read-aloud/voices/zh_CN.onnx.json
```

如果 Piper 不存在或模型缺失，服务会自动尝试 eSpeak NG。

## 安装

```bash
cd packaging/kylin
chmod +x install.sh uninstall.sh
./install.sh
```

安装后检查：

```bash
curl http://127.0.0.1:19860/health
```

## WPS 加载项安装

把 `addin/` 注册到 WPS JS 加载项目录，或使用 WPS 提供的加载项管理/调试工具加载 `manifest.xml`。不同政企版 WPS 的加载项目录可能被策略定制，建议以目标机实际 WPS 管理入口为准。

加载项页面默认访问：

```text
http://127.0.0.1:19860
```

所以请先启动 `wps-tts-daemon`。

## API

```http
GET /health
POST /speak
POST /stop
POST /pause
POST /resume
GET /voices
```

`POST /speak` 示例：

```json
{
  "text": "需要朗读的内容",
  "voice": "zh_CN",
  "rate": 1.0,
  "volume": 80
}
```

## 生成最终 DEB 安装包

最终交付目标是：

```text
dist/wps-read-aloud-zhangjingyao_1.0.1_arm64.deb
```

打包前必须放入：

```text
engines/piper/piper
engines/piper/lib/
engines/espeak-ng/espeak-ng
engines/espeak-ng/espeak-ng-data/
engines/espeak-ng/lib/
voices/zh_CN.onnx
voices/zh_CN.onnx.json
```

这些文件会被打入 `.deb`，安装到 `/opt/wps-read-aloud/engines` 和 `/opt/wps-read-aloud/voices`。打包脚本会强制校验，缺任意一项都会失败。

构建：

```bash
chmod +x packaging/kylin/build-arm64.sh packaging/deb/build-deb.sh
./packaging/kylin/build-arm64.sh
./packaging/deb/build-deb.sh
```

安装：

```bash
sudo dpkg -i dist/wps-read-aloud-zhangjingyao_1.0.1_arm64.deb
```

安装包会安装并启动 `wps-tts.service`，同时为已有普通用户注册 WPS 加载项。若 WPS 已打开，需要重启 WPS。
