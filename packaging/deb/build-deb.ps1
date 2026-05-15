$ErrorActionPreference = "Stop"

$RootDir = Resolve-Path (Join-Path $PSScriptRoot "..\..")
$PkgName = "wps-read-aloud-zhangjingyao"
$Version = if ($env:VERSION) { $env:VERSION } else { "1.0.1" }
$Arch = if ($env:ARCH) { $env:ARCH } else { "arm64" }
$BuildDir = Join-Path $RootDir "build\deb\${PkgName}_${Version}_${Arch}"
$OutDir = Join-Path $RootDir "dist"
$DebPath = Join-Path $OutDir "${PkgName}_${Version}_${Arch}.deb"

function Require-Path($Path) {
  if (-not (Test-Path $Path)) {
    throw "missing required file: $Path"
  }
}

function Write-ArMember($Stream, [string]$Name, [byte[]]$Data) {
  $timestamp = "0"
  $owner = "0"
  $group = "0"
  $mode = "100644"
  $size = [string]$Data.Length
  $header = "{0,-16}{1,-12}{2,-6}{3,-6}{4,-8}{5,-10}`n" -f $Name, $timestamp, $owner, $group, $mode, $size
  $headerBytes = [System.Text.Encoding]::ASCII.GetBytes($header)
  if ($headerBytes.Length -ne 60) {
    throw "invalid ar header for $Name"
  }
  $Stream.Write($headerBytes, 0, $headerBytes.Length)
  $Stream.Write($Data, 0, $Data.Length)
  if (($Data.Length % 2) -ne 0) {
    $Stream.WriteByte(10)
  }
}

Require-Path (Join-Path $RootDir "dist\wps-tts-daemon")
Require-Path (Join-Path $RootDir "engines\piper\piper")
Require-Path (Join-Path $RootDir "engines\piper\lib")
Require-Path (Join-Path $RootDir "engines\espeak-ng\espeak-ng")
Require-Path (Join-Path $RootDir "engines\espeak-ng\espeak-ng-data")
Require-Path (Join-Path $RootDir "engines\espeak-ng\lib")
Require-Path (Join-Path $RootDir "voices\zh_CN.onnx")
Require-Path (Join-Path $RootDir "voices\zh_CN.onnx.json")

if (Test-Path $BuildDir) {
  Remove-Item -Recurse -Force $BuildDir
}

$DebianDir = Join-Path $BuildDir "DEBIAN"
$DataRoot = Join-Path $BuildDir "data"
New-Item -ItemType Directory -Force -Path $DebianDir, $DataRoot, $OutDir | Out-Null

$paths = @(
  "opt\wps-read-aloud\daemon",
  "opt\wps-read-aloud\addin",
  "opt\wps-read-aloud\engines",
  "opt\wps-read-aloud\voices",
  "etc\wps-read-aloud",
  "usr\bin",
  "lib\systemd\system"
)
foreach ($path in $paths) {
  New-Item -ItemType Directory -Force -Path (Join-Path $DataRoot $path) | Out-Null
}

Copy-Item (Join-Path $RootDir "packaging\deb\control") (Join-Path $DebianDir "control")
(Get-Content (Join-Path $DebianDir "control")) |
  ForEach-Object {
    $_ -replace "^Version:.*", "Version: $Version" -replace "^Architecture:.*", "Architecture: $Arch"
  } |
  Set-Content -Encoding ASCII (Join-Path $DebianDir "control")

Copy-Item (Join-Path $RootDir "packaging\deb\preinst") (Join-Path $DebianDir "preinst")
Copy-Item (Join-Path $RootDir "packaging\deb\postinst") (Join-Path $DebianDir "postinst")
Copy-Item (Join-Path $RootDir "packaging\deb\prerm") (Join-Path $DebianDir "prerm")
Copy-Item (Join-Path $RootDir "packaging\deb\postrm") (Join-Path $DebianDir "postrm")

Copy-Item (Join-Path $RootDir "dist\wps-tts-daemon") (Join-Path $DataRoot "opt\wps-read-aloud\daemon\wps-tts-daemon")
Copy-Item -Recurse -Force (Join-Path $RootDir "addin\*") (Join-Path $DataRoot "opt\wps-read-aloud\addin")
Copy-Item -Recurse -Force (Join-Path $RootDir "engines\*") (Join-Path $DataRoot "opt\wps-read-aloud\engines")
Copy-Item -Recurse -Force (Join-Path $RootDir "voices\*") (Join-Path $DataRoot "opt\wps-read-aloud\voices")
Copy-Item (Join-Path $RootDir "daemon\config.example.yaml") (Join-Path $DataRoot "etc\wps-read-aloud\config.yaml")
Copy-Item (Join-Path $RootDir "packaging\deb\wps-tts.service") (Join-Path $DataRoot "lib\systemd\system\wps-tts.service")
Copy-Item (Join-Path $RootDir "packaging\deb\wps-read-aloud-register") (Join-Path $DataRoot "usr\bin\wps-read-aloud-register")

$DebianBinary = Join-Path $BuildDir "debian-binary"
Set-Content -NoNewline -Encoding ASCII $DebianBinary "2.0`n"

$ControlTar = Join-Path $BuildDir "control.tar.gz"
$DataTar = Join-Path $BuildDir "data.tar.gz"
Push-Location $DebianDir
tar --format=ustar -czf $ControlTar control preinst postinst prerm postrm
Pop-Location
Push-Location $DataRoot
tar --format=ustar -czf $DataTar .
Pop-Location

$fs = [System.IO.File]::Open($DebPath, [System.IO.FileMode]::Create)
try {
  $magic = [System.Text.Encoding]::ASCII.GetBytes("!<arch>`n")
  $fs.Write($magic, 0, $magic.Length)
  Write-ArMember $fs "debian-binary" ([System.IO.File]::ReadAllBytes($DebianBinary))
  Write-ArMember $fs "control.tar.gz" ([System.IO.File]::ReadAllBytes($ControlTar))
  Write-ArMember $fs "data.tar.gz" ([System.IO.File]::ReadAllBytes($DataTar))
} finally {
  $fs.Dispose()
}

Write-Host "created $DebPath"
