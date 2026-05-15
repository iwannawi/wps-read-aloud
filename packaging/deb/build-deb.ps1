$ErrorActionPreference = "Stop"

$RootDir = Resolve-Path (Join-Path $PSScriptRoot "..\..")
$Python = Get-Command python -ErrorAction SilentlyContinue
if (-not $Python) {
  $Python = Get-Command python3 -ErrorAction SilentlyContinue
}
if (-not $Python) {
  throw "missing required command: python"
}

& $Python.Source (Join-Path $RootDir "packaging\deb\build_deb.py")
