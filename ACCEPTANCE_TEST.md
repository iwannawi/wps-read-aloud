# 验收测试说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-comate
版本：1.1.6

## 环境矩阵

| CPU 架构 + 操作系统 | WPS 要求 | 安装包 |
| --- | --- | --- |
| x86/x64 Windows 10/11 | WPS Office 2019 或更高版本 | wps-read-aloud-comate_1.1.6_windows.exe |
| x64 银河麒麟 V10 及以上 | WPS Office 2019 for Linux 或更高版本 | wps-read-aloud-comate_1.1.6_amd64.deb |
| ARM64 银河麒麟 V10 及以上 | WPS Office 2019 for Linux 或更高版本 | wps-read-aloud-comate_1.1.6_arm64.deb |
| x64 UOS V20 | WPS Office 2019 for Linux 或更高版本 | cn.wps-read-aloud-comate_1.1.6_amd64.deb |
| ARM64 UOS V20 | WPS Office 2019 for Linux 或更高版本 | cn.wps-read-aloud-comate_1.1.6_arm64.deb |

## 安装验收

| CPU 架构 + 操作系统 | 操作 |
| --- | --- |
| x86/x64 Windows 10/11 | 运行 wps-read-aloud-comate_1.1.6_windows.exe |
| x64 银河麒麟 V10 及以上 | sudo dpkg -i wps-read-aloud-comate_1.1.6_amd64.deb |
| ARM64 银河麒麟 V10 及以上 | sudo dpkg -i wps-read-aloud-comate_1.1.6_arm64.deb |
| x64 UOS V20 | sudo dpkg -i cn.wps-read-aloud-comate_1.1.6_amd64.deb |
| ARM64 UOS V20 | sudo dpkg -i cn.wps-read-aloud-comate_1.1.6_arm64.deb |

预期结果：

- 安装过程无错误退出。
- Windows 安装器能显示安装进度、安装路径、当前动作和最终结果。
- Windows 安装日志写入 %LOCALAPPDATA%\WPSReadAloudComate\Logs\install.log。
- Windows 当前用户 Run 自启动项写入 WPSReadAloudComate。
- Linux systemd 服务 wps-read-aloud-comate.service 启动并保持运行。
- Linux 安装日志写入 /var/log/wps-read-aloud-install.log。
- 覆盖安装同版本或升级安装时不发生文件覆盖冲突。
- 安装后重启 WPS，可看到“文档朗读”选项卡。

## 选项卡验收

1. 打开 WPS 文字。
2. 确认顶部出现“文档朗读”选项卡。
3. 确认按钮为：开始朗读、停止朗读、朗读方式、朗读语速、状态检查、关于朗读。
4. 未朗读时，“停止朗读”置灰。
5. 朗读中，“开始朗读”“朗读方式”“朗读语速”“状态检查”“关于朗读”置灰，“停止朗读”可用。
6. “朗读方式”包含“连页朗读”和“当页朗读”，默认“连页朗读”。
7. “朗读语速”包含 0.75x、1x、1.2x、1.5x，默认 1.2x。
8. 任何按钮不应弹出“未知按钮”。

## 朗读验收

- 连页朗读：有光标时从光标处读到文档末尾；无可识别光标时从文档开头读到文档末尾。
- 当页朗读：有光标时从光标处读到当前页末尾；无可识别光标时从当前页开头读到当前页末尾。
- 朗读过程中，当前语句应在 WPS 文档中同步选中。
- 停止朗读后，不继续播放后续句子。
- 启动提示显示“朗读服务正在启动，请耐心等待...”。
- 启动提示不固定倒计时关闭，进入实际播放后自动关闭。
- 低性能机器上仍可点击“停止朗读”结束任务。

## 中英文数字验收

示例文本：

    这是 WPS 2026 read aloud test，版本是 v1.1.6。

预期结果：

- 中文正常朗读。
- 英文和数字不被跳过。
- 英文和数字按单字符中文读法朗读。
- 逗号、顿号、分号、冒号等语义标点处有自然停顿。
- 双引号、单引号、书名号、括号等成对符号不额外增加停顿。

## 弹窗验收

- 弹窗不出现黄色三角叹号图标。
- 状态检查显示服务版本、语音引擎、自检结果、播放器和探测时间。
- 关于朗读标题为“WPS 文档朗读助手”。
- 关于朗读信息不重叠，默认大小下尽量不出现滚动条。
- 关于朗读中的说明文件可在同一弹窗内打开，并可返回关于页。
- 关闭弹窗后 WPS 不应异常最小化。

## 异常验收

- 连续快速点击“开始朗读”不会产生多段音频同时播放。
- 朗读中关闭 WPS 后，服务端不应残留长期运行的合成或播放进程。
- 网络不可用时仍可离线朗读。
- 服务只访问 127.0.0.1:19860。
