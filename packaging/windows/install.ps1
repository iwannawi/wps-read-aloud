param(
  [string]$InstallDir = "$env:ProgramFiles\WPS Read Aloud XC"
)

$ErrorActionPreference = "Stop"
$LogDir = Join-Path $env:ProgramData "WPSReadAloudXC"
$LogFile = Join-Path $LogDir "install.log"
New-Item -ItemType Directory -Force -Path $LogDir | Out-Null
Start-Transcript -Path $LogFile -Append | Out-Null

function Backup-ConfigFile {
  param([string]$Path)
  if (Test-Path $Path) {
    $Stamp = Get-Date -Format "yyyyMMddHHmmss"
    Copy-Item -LiteralPath $Path -Destination "$Path.bak.$Stamp" -Force
  }
}

function Set-WpsPluginEntry {
  param(
    [string]$Path,
    [string]$Entry
  )

  Backup-ConfigFile -Path $Path
  if (Test-Path $Path) {
    $Content = Get-Content -Raw -Path $Path -Encoding UTF8
    $Content = [regex]::Replace($Content, '(?is)\s*<jspluginonline\b[^>]*name="wps-read-aloud"[^>]*/>', '')
    $Content = [regex]::Replace($Content, '(?is)\s*<jsplugin\b[^>]*name="wps-read-aloud"[\s\S]*?</jsplugin>', '')
    if ($Content -match '</jsplugins>') {
      $Content = $Content -replace '(?is)</jsplugins>', "  $Entry`r`n</jsplugins>"
    }
    else {
      $Content = "<?xml version=`"1.0`" encoding=`"UTF-8`"?>`r`n<jsplugins>`r`n  $Entry`r`n</jsplugins>`r`n"
    }
  }
  else {
    $Content = "<?xml version=`"1.0`" encoding=`"UTF-8`"?>`r`n<jsplugins>`r`n  $Entry`r`n</jsplugins>`r`n"
  }
  Set-Content -Path $Path -Value $Content -Encoding UTF8
}

try {
  $Source = Join-Path $PSScriptRoot "app"
  if (!(Test-Path $Source)) {
    throw "安装包不完整：未找到 app 目录。"
  }
  New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
  Copy-Item -Path (Join-Path $Source "*") -Destination $InstallDir -Recurse -Force

  $Daemon = Join-Path $InstallDir "daemon\wps-tts-daemon.exe"
  if (!(Test-Path $Daemon)) {
    throw "安装包不完整：未找到 wps-tts-daemon.exe。"
  }

  $TaskName = "WPSReadAloudXC"
  $Action = New-ScheduledTaskAction -Execute $Daemon -Argument "-config `"$InstallDir\config.yaml`"" -WorkingDirectory $InstallDir
  $Trigger = New-ScheduledTaskTrigger -AtLogOn
  $Principal = New-ScheduledTaskPrincipal -UserId $env:USERNAME -LogonType Interactive -RunLevel LeastPrivilege
  Register-ScheduledTask -TaskName $TaskName -Action $Action -Trigger $Trigger -Principal $Principal -Force | Out-Null
  Start-ScheduledTask -TaskName $TaskName

  $JsDir = Join-Path $env:APPDATA "Kingsoft\wps\jsaddons"
  New-Item -ItemType Directory -Force -Path $JsDir | Out-Null
  $VersionInfo = Get-Content -Raw -Path (Join-Path $InstallDir "version.json") | ConvertFrom-Json
  $AddinVersion = $VersionInfo.version
  $Target = Join-Path $JsDir "wps-read-aloud_$AddinVersion"
  New-Item -ItemType Directory -Force -Path $Target | Out-Null
  Copy-Item -Path (Join-Path $InstallDir "addin\*") -Destination $Target -Recurse -Force

  $Index = (Join-Path $Target "index.html").Replace("\", "/")
  $Ribbon = (Join-Path $Target "ribbon.xml").Replace("\", "/")
  $FileUrl = "file:///$Index"
  $LocalUrl = "http://127.0.0.1:19860/addin/index.html"
  $PublishXml = Join-Path $JsDir "publish.xml"
  $PluginsXml = Join-Path $JsDir "jsplugins.xml"
  $OnlineEntry = "<jspluginonline name=`"wps-read-aloud`" type=`"wps`" enable=`"enable_dev`" install=`"$LocalUrl`" url=`"$LocalUrl`" debug=`"`"/>"
  $LocalEntry = @"
<jsplugin name="wps-read-aloud" type="wps" url="$FileUrl" version="$AddinVersion" desc="WPS 文档朗读助手. Developer: Zhang Jingyao.">
    <ribbon file="$Ribbon"/>
  </jsplugin>
"@
  Set-WpsPluginEntry -Path $PublishXml -Entry $OnlineEntry
  Set-WpsPluginEntry -Path $PluginsXml -Entry $LocalEntry

  Write-Host "WPS 文档朗读助手安装完成。若 WPS 已打开，请重启 WPS。"
  Write-Host "安装日志：$LogFile"
}
finally {
  Stop-Transcript | Out-Null
}

