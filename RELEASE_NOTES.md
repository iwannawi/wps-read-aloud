# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.1.1
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

- 新增 Windows 朗读预合成队列，播放当前句时提前合成后续句。
- 新增数学符号中文读法，覆盖加、减、乘、除、等于、大于、小于、百分号等常见符号。

## 变更

- Windows 安装界面改为绘制式头部区域，使用透明背景 logo，避免白底遮挡和文字不可见。
- 首次授权提示文案改短，降低用户理解成本。
- 超长文档超过 1000 句时不再额外弹窗，只在 WPS 状态栏提示本次朗读范围。
- GitHub Release 发布流程改为使用 UTF-8 安全方式生成说明文件，避免中文乱码。

## 修复

- 修复 Windows 朗读上一句和下一句之间等待时间过长的问题。
- 修复 Windows 安装器 logo 白底和头部文字显示异常的问题。
- 修复旧 WPS 加载项注册项残留可能导致“打开JS调试器”入口继续显示的问题。
- 修复 Release 草稿发布超时后可能留下乱码说明和不完整附件的问题。

## 交付文件

| 目标 | 文件 |
| --- | --- |
| x86/x64 Windows 10/11 | dist/wps-read-aloud-comate_1.1.1_windows.exe |
| x64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.1_amd64.deb |
| ARM64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.1_arm64.deb |
| x64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.1_amd64.deb |
| ARM64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.1_arm64.deb |

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对文档 Range、Selection、Page 等接口的支持。
- 朗读过程中不能切换朗读方式和语速，需要停止朗读后再调整。
- 音量由目标操作系统、声卡设备或桌面环境声音设置控制，加载项内不提供音量调节。
