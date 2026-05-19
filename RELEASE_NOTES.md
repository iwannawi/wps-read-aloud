# 发布说明

软件名称：WPS 文档朗读助手
软件包：wps-read-aloud-xc
Debian 包名：银河麒麟为 “wps-read-aloud-xc”，UOS 为 “cn.wps-read-aloud-xc”
版本：1.0.29
架构：Windows x86、Linux amd64、Linux arm64
适用操作系统：Windows x86、银河麒麟 x64/ARM64、UOS x64/ARM64
适用办公软件：WPS Office 2023 for Linux / WPS Office 2019 for Linux
开发者：Zhang Jingyao
发布时间：20260519

## 本次变更

- 项目改造为同一套源码、多目标安装包结构，新增 Windows x86 exe 安装程序、银河麒麟 amd64、银河麒麟 arm64、UOS amd64、UOS arm64 五类安装包矩阵。
- Linux deb 打包脚本按目标系统和架构生成不同包；Debian 文件名统一为 “包名_版本_架构.deb”，UOS 包名使用 “cn.wps-read-aloud-xc”，银河麒麟包名使用 “wps-read-aloud-xc”。
- UOS 安装路径调整为 “/opt/apps/cn.wps-read-aloud-xc/files”，符合 UOS 应用目录习惯；银河麒麟继续使用 “/opt/wps-read-aloud” 和 “/etc/wps-read-aloud/config.yaml”。
- 守护进程新增运行时根目录环境变量支持，systemd 服务会按目标系统传入对应安装目录，避免 UOS 包安装后仍访问麒麟路径。
- Windows 安装包默认改为当前用户安装，路径为 “%LOCALAPPDATA%\Programs\WPS Read Aloud XC”，安装日志写入 “%LOCALAPPDATA%\WPSReadAloudXC\Logs”，避免无管理员权限时写入 “C:\Program Files (x86)” 失败。
- 新增 “resources/runtime” 运行时资源目录规范，按系统和架构放置 sherpa-onnx、ONNX Runtime 等原生二进制依赖。
- 正式构建不再从旧版 “engines” 目录复制运行时资源，避免把目标环境不需要的库和 Piper、eSpeak NG 等已弃用资源带入安装包。
- 新增 Windows 安装包骨架，包含安装、卸载、WPS 加载项注册和登录任务启动本地朗读服务的脚本。实际生成 Windows 安装包前，需要补齐 Windows x86 版 daemon 和 sherpa-onnx 运行时资源。
- 关于朗读弹窗尺寸调整为 960×720，避免默认窗口下出现滚动条或内容拥挤。
- 全部中文说明文件清理英文反引号符号，路径、命令、组件名等改用中文引号或缩进排版，提高中文阅读一致性。
- 服务端 “/addin/assets/icons/*.png” 改为优先读取 “/opt/wps-read-aloud/addin/assets/icons/” 中的真实图标文件，读取不到时再回退到 daemon 内嵌资源；图标响应增加 “no-store”，便于现场手动替换验证。
- 修复朗读启动等待弹窗在点击“停止朗读”后不立即关闭的问题；停止时会主动发送关闭信号，弹窗也会在停止、错误、完成或播放开始后自动关闭。
- “文档朗读”选项卡 6 个按钮图标继续使用 128x128 PNG 图标，“getImage” 回调改为返回加载项内静态资源路径，例如 “assets/icons/start.png”。该方案规避 WPS Linux 对 Base64 图标返回值兼容性不稳定、显示问号占位图的问题。
- 关于朗读弹窗按成熟软件的“关于”信息重新整理，保留版本、发布日期、开发者、软件包、适用环境、服务地址、版权和第三方开源组件说明。
- Copyright 信息改为 “Copyright © 2026 Zhang Jingyao”，第三方组件版权归各自权利人所有。
- 第三方声明和模型许可说明改为中文排版，统一使用“组织或个人”等中性说法。
- 核实第三方许可信息：Sherpa-onnx 为 Apache License 2.0，ONNX Runtime 为 MIT License；“vits-zh-hf-fanchen-C” 模型权重当前未声明明确许可证，正式分发或商用前需单独完成授权确认，或替换为许可明确的中文模型。
- 状态检查弹窗保持固定摘要布局，只显示关键状态和播放器探测摘要，避免探测明细过多导致内容显示不全或出现整窗滚动条。
- 句内逗号、顿号、冒号、分号等语义标点继续通过文本预处理提示朗读节奏，不拆分为多个 TTS 合成任务，避免明显增加整句合成时间。
- 句末停顿改为在整句 wav 末尾追加精确静音：默认 “1.2x” 语速下句内标准停顿按约 “400ms” 设计，句末追加 “600ms” 静音；切换到其他语速时按比例缩放停顿时长。
- 朗读语速调整为 4 个选项：“0.75x”、“1x”、“1.2x”、“1.5x”，默认语速改为 “1.2x”。
- 朗读启动等待弹窗文案改为“朗读服务正在启动，请耐心等待...”，并适当增大提示字号。
- 文本预处理保留单句一次语音合成，不会因逗号、冒号等语义标点拆成多个 wav；语义标点会保留文本级停顿提示，双引号、单引号、书名号、括号等成对符号不额外增加停顿。
- 服务版本号从 Go 编译常量解耦，改为运行时读取安装目录中的 “version.json”；以后仅修改前端、图标、文档或发布信息时，可以复用现有 daemon 二进制重新打包，不必重新编译 Go 服务。
- 打包脚本会自动写入 “version.json”，并在 “dist/wps-tts-daemon” 不存在时从最近一个已生成的 “.deb” 中复用 daemon 二进制，减少前端小改版的构建步骤。
- GitHub 推送和 Release 发布脚本改为长期复用脚本，优先读取 “gh auth token”，其次读取本机 Git Credential Manager，并会校验 token 有效性；日志不会输出 token 或 Basic 认证头。
- 文档朗读选项卡的 6 个顶层控件继续使用 “size="large"”，并恢复 WPS JS 加载项更稳定的 “getImage="ribbon.GetImage"” 图标回调方式，避免静态 “image="..."” 在 WPS Linux 中不显示。
- 选项卡图标改为通过加载项静态资源路径返回，减少 Base64 和 data URI 在 WPS Linux Ribbon 中的兼容性风险。
- 朗读启动等待弹窗改为双层居中布局，去掉 compact 模式下的最小宽度限制，修复提示文字在小窗内偏右的问题。
- 朗读启动等待弹窗隐藏标题行，只保留正文提示，避免重复表达。
- 删除旧版单平台脚本和单平台部署说明，统一使用多平台构建入口和多平台交付文档。
- 软件包文件名统一为 Debian 习惯格式，例如 “wps-read-aloud-xc_1.0.29_arm64.deb” 和 “cn.wps-read-aloud-xc_1.0.29_arm64.deb”。
- 朗读启动等待弹窗改为更小的 compact 内容布局，降低标题、正文和内边距尺寸，并禁用 compact 弹窗内部滚动，避免小窗出现滚动条。
- 清理发布目录中的旧安装包、临时推送脚本和可能包含敏感认证输出的日志，避免交付目录混入无关文件。
- 开始朗读时的小提示窗改为紧凑布局，不再依赖固定倒计时自动关闭；进入实际播放后由加载项主动关闭。
- 朗读进行中禁用开始朗读、朗读方式、朗读语速、状态检查和关于朗读，停止朗读保持可用，避免连续点击造成状态不一致。
- 软件显示名称统一为“WPS 文档朗读助手”，选项卡名称仍保持“文档朗读”。
- 安装包增加 “Conflicts/Replaces: wps-read-aloud-zhangjingyao”，支持从旧包名升级，减少 dpkg 文件归属冲突。
- 关于朗读弹窗内容改为更紧凑布局，默认大小下尽量不出现滚动条。
- 预合成策略从固定前三句改为动态按句累计：启动阶段持续预处理句子，累计文本达到约 100 字即开始播放；若第一句超过 100 字，则只等待第一句。
- 播放中的后续预合成也按约 100 字窗口向前准备，兼顾长句启动速度和短句连续播放。

## 安装

    sudo dpkg -i wps-read-aloud-xc_1.0.29_arm64.deb

如果 WPS 已经打开，请安装完成后重启 WPS。

## 已知限制

- 当前句选中和翻页依赖 WPS Office for Linux 对 Range/Selection/Page 接口的支持；如果目标 WPS 版本接口行为不同，当页朗读会尽量回退到可读范围。
- 朗读过程中不能切换朗读方式和语速，需要停止朗读后再调整。
- 音量由麒麟系统音量、声卡或桌面环境声音设置控制，加载项内不再提供音量调节。




