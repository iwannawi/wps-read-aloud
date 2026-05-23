# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.1.11
发布时间：20260523
开发者：Zhang Jingyao

## 适用环境

| CPU 架构 + 操作系统 | WPS 要求 |
| --- | --- |
| x86/x64 Windows 10/11 | WPS Office 2019 或更高版本，推荐最新稳定版 |
| x64 银河麒麟 V10 及以上 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| ARM64 银河麒麟 V10 及以上 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| x64 UOS V20 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |
| ARM64 UOS V20 | WPS Office 2019 for Linux 或更高版本，推荐最新稳定版 |

## 变更

- 朗读文本预处理增加数学读法规则，百分数按“百分之”朗读，常见运算符、比较符、集合符号和数学符号会转换为中文读法。
- 英文专有词读法调整：“WPS”读作“达不溜屁挨思”，“Office”和“office”读作“凹斐思”。
- Windows 安装脚本升级时保留并刷新 WPS 已有授权缓存，避免每次升级都主动清掉用户已允许的加载项记录。
- README、验收测试、发布说明和构建经验文档同步更新为当前 1.1.11 方案。

## 修复

- 修复 “10%” 被读成逐字符或“百分号”的问题，现在读作“百分之十”。
- 修复 “WPS Office” 在中文模型下逐字母读法不符合预期的问题。
- 修复 Windows 升级安装后由于授权缓存被清理，可能再次出现重复加载项许可确认的问题。

## 交付文件

| 目标 | 文件 |
| --- | --- |
| x86/x64 Windows 10/11 | dist/wps-read-aloud-comate_1.1.11_windows.exe |
| x64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.11_amd64.deb |
| ARM64 银河麒麟 V10 及以上 | dist/wps-read-aloud-comate_1.1.11_arm64.deb |
| x64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.11_amd64.deb |
| ARM64 UOS V20 | dist/cn.wps-read-aloud-comate_1.1.11_arm64.deb |

## 已知限制

- 当前句选中和翻页依赖 WPS Office 对文档 Range、Selection、Page 等接口的支持。
- 朗读过程中不能切换朗读方式和语速，需要停止朗读后再调整。
- 音量由目标操作系统、声卡设备或桌面环境声音设置控制，加载项内不提供音量调节。
- Windows WPS 首次信任第三方加载项时可能显示 WPS 原生安全确认框。安装脚本会保留已允许记录，避免升级时重复触发；首次确认是否出现由 WPS 客户端安全策略决定。
