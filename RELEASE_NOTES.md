# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
Debian 包名：x64/ARM64 银河麒麟 V10 及以上为 “wps-read-aloud-comate”，x64/ARM64 UOS V20 为 “cn.wps-read-aloud-comate”
版本：1.0.39
架构和系统：x86/x64 Windows 10/11、x64/ARM64 银河麒麟 V10 及以上、x64/ARM64 UOS V20
适用办公软件：x86/x64 Windows 10/11 要求 WPS Office 2019 或更高版本，推荐 WPS Office 最新稳定版；x64/ARM64 银河麒麟 V10 及以上、x64/ARM64 UOS V20 要求 WPS Office 2019 或更高版本，推荐最新版 WPS Office for Linux。
开发者：Zhang Jingyao
发布时间：20260520

## 本版本变更

- 修复 x86/x64 Windows 10/11 点击“开始朗读”后 Sherpa-onnx 语音引擎启动失败的问题，Windows daemon 改用当前 Sherpa-onnx 支持的 “--vits-tokens” 参数。
- 修复部分 WPS 内置浏览器不支持 “URLSearchParams” 时弹窗空白的问题，弹窗参数解析改为兼容实现。
- 状态检查增加语音引擎自检，并在 x86/x64 Windows 10/11 显示 Windows SoundPlayer 探测时间和结果。
- WPS 加载项注册改为正式启用模式，去掉选项卡右侧“打开JS调试器”入口。
- 优化关于朗读弹窗排版，避免长平台信息重叠。
- 优化 x86/x64 Windows 10/11 安装界面，增加专属图标、头图、任务栏图标和更友好的安装提示。
- 内部经验文档不再进入安装包，并在发布校验中增加拦截规则。
- 全部交付包版本更新为 1.0.39。

## 交付文件

    dist/wps-read-aloud-comate_1.0.39_windows.exe
    dist/wps-read-aloud-comate_1.0.39_amd64.deb
    dist/wps-read-aloud-comate_1.0.39_arm64.deb
    dist/cn.wps-read-aloud-comate_1.0.39_amd64.deb
    dist/cn.wps-read-aloud-comate_1.0.39_arm64.deb

## 安装提示

x86/x64 Windows 10/11：

    运行 dist/wps-read-aloud-comate_1.0.39_windows.exe

x64 银河麒麟 V10 及以上：

    sudo dpkg -i dist/wps-read-aloud-comate_1.0.39_amd64.deb

ARM64 银河麒麟 V10 及以上：

    sudo dpkg -i dist/wps-read-aloud-comate_1.0.39_arm64.deb

x64 UOS V20：

    sudo dpkg -i dist/cn.wps-read-aloud-comate_1.0.39_amd64.deb

ARM64 UOS V20：

    sudo dpkg -i dist/cn.wps-read-aloud-comate_1.0.39_arm64.deb

如果 WPS 已经打开，请安装完成后重启 WPS，再使用顶部“文档朗读”选项卡。

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对 Range、Selection、Page 等接口的支持；如果目标 WPS 版本接口行为不同，当页朗读会尽量回退到可读范围。
- 朗读过程中不能切换朗读方式和语速，需要停止朗读后再调整。
- 音量由目标操作系统、声卡设备或桌面环境声音设置控制，加载项内不提供音量调节。
