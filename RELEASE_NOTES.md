# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.1.0
发布时间：20260520
开发者：Zhang Jingyao

## 适用环境

| CPU 架构 + 操作系统 | WPS 要求 |
| --- | --- |
| x86/x64 Windows 10/11 | WPS Office 2019 或更高版本，推荐最新稳定版 |
| x64 银河麒麟 V10 及以上 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| ARM64 银河麒麟 V10 及以上 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| x64 UOS V20 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| ARM64 UOS V20 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |

## 新增

- 新增 Windows 安装器专属图标、头图和任务栏图标。
- 新增 Windows 状态检查自检结果，显示语音引擎、播放器和探测时间。

## 变更

- 项目发布版本升级为 1.1.0，后续按语义化版本规则管理。
- Release 名称改为“wps-read-aloud-comate 版本号 发布时间”。
- README、验收说明、多平台打包说明和版本管理说明改为表格化、短句化表达。
- WPS 加载项注册改为正式启用模式，不再显示“打开JS调试器”入口。
- 关于朗读弹窗压缩长平台描述，降低文字重叠风险。

## 修复

- 修复 Windows 点击“开始朗读”后 Sherpa-onnx 启动失败的问题。
- 修复 Windows 语音引擎参数不兼容问题，改用当前 Sherpa-onnx 支持的参数组合。
- 修复部分 WPS 内置浏览器不支持 URLSearchParams 时弹窗空白的问题。
- 修复状态检查显示“尚未探测”的问题。
- 修复内部经验文档误进入安装包的风险，发布校验会拦截此类文件。

## 交付文件

| 目标 | 文件 |
| --- | --- |
| x86/x64 Windows 10/11 | dist/wps-read-aloud-comate_1.1.0_windows.exe |
| x64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.0_amd64.deb |
| ARM64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.0_arm64.deb |
| x64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.0_amd64.deb |
| ARM64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.0_arm64.deb |

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对文档 Range、Selection、Page 等接口的支持。
- 朗读过程中不能切换朗读方式和语速，需要停止朗读后再调整。
- 音量由目标操作系统、声卡设备或桌面环境声音设置控制，加载项内不提供音量调节。
