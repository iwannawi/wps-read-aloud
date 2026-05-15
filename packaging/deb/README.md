# DEB 打包说明

目标交付文件：

```text
dist/wps-read-aloud-zhangjingyao_1.0.1_arm64.deb
```

## 必需输入

构建 `.deb` 前需要准备：

```text
dist/wps-tts-daemon
engines/piper/piper
engines/piper/lib/
engines/espeak-ng/espeak-ng
engines/espeak-ng/espeak-ng-data/
engines/espeak-ng/lib/
voices/zh_CN.onnx
voices/zh_CN.onnx.json
```

`engines/` 和 `voices/` 会被原样打进安装包。当前 `build-deb.sh` 会强制校验 Piper、eSpeak NG、动态库目录和中文模型是否存在，避免生成“装得上但不能朗读”的安装包。

包内服务运行时会给 Piper、eSpeak NG 分别设置 `LD_LIBRARY_PATH`，优先加载 `/opt/wps-read-aloud/engines/*/lib` 下的库。音频播放由 WPS 加载项页面完成，不依赖系统 `aplay`。

## 构建

在银河麒麟 ARM64 或其他 Linux 打包机执行：

```bash
chmod +x packaging/kylin/build-arm64.sh packaging/deb/build-deb.sh
./packaging/kylin/build-arm64.sh
./packaging/deb/build-deb.sh
```

## 安装

```bash
sudo dpkg -i dist/wps-read-aloud-zhangjingyao_1.0.1_arm64.deb
```

安装后会：

- 安装 `/opt/wps-read-aloud`
- 安装并启动系统服务 `wps-tts.service`
- 为已有普通用户写入 WPS 加载项注册文件
- 提供 `wps-read-aloud-register` 用于后续用户或注册修复
- 写入安装/卸载日志：`/var/log/wps-read-aloud-install.log`
- 安装第三方组件许可证文件：`/usr/share/doc/wps-read-aloud-zhangjingyao/`
- 安装发布说明、验收测试说明、源码获取说明、校验说明到同一文档目录

## 验证

```bash
systemctl status wps-tts.service
curl http://127.0.0.1:19860/health
wps-read-aloud-register
```

重启 WPS 后查看“离线朗读”加载项。
