param(
  [string]$Owner = "iwannawi",
  [string]$Repo = "wps-read-aloud",
  [string]$Branch = "main",
  [string]$Tag = "",
  [switch]$PromptToken
)

$ErrorActionPreference = "Continue"
if (Get-Variable PSNativeCommandUseErrorActionPreference -ErrorAction SilentlyContinue) {
  $PSNativeCommandUseErrorActionPreference = $false
}

$Root = Resolve-Path (Join-Path $PSScriptRoot "..")
$Log = Join-Path $Root "dist\github-push.log"

function Write-Log($Text) {
  $Text | Tee-Object -FilePath $Log -Append
}

function ConvertFrom-SecureStringPlain($SecureString) {
  $Ptr = [Runtime.InteropServices.Marshal]::SecureStringToBSTR($SecureString)
  try {
    return [Runtime.InteropServices.Marshal]::PtrToStringBSTR($Ptr)
  } finally {
    [Runtime.InteropServices.Marshal]::ZeroFreeBSTR($Ptr)
  }
}

function Get-GitHubTokenFromGcm {
  $Gcm = "C:\Users\zhangjingyao\scoop\apps\git\2.53.0.2\mingw64\bin\git-credential-manager.exe"
  if (!(Test-Path $Gcm)) {
    return ""
  }
  $InputText = "protocol=https`nhost=github.com`npath=$Owner/$Repo.git`n`n"
  $Cred = $InputText | & $Gcm get
  if ($LASTEXITCODE -ne 0 -or !$Cred) {
    return ""
  }
  foreach ($Line in $Cred) {
    if ($Line.StartsWith("password=")) {
      return $Line.Substring("password=".Length)
    }
  }
  return ""
}

function Get-GitHubToken {
  if (!$PromptToken) {
    $Token = Get-GitHubTokenFromGcm
    if (![string]::IsNullOrWhiteSpace($Token)) {
      return $Token
    }
  }
  $Secure = Read-Host "Enter GitHub token" -AsSecureString
  return ConvertFrom-SecureStringPlain $Secure
}

function Invoke-Git($DisplayArgs, $ActualArgs) {
  Write-Log ("git " + ($DisplayArgs -join " "))
  & git @ActualArgs 2>&1 | Tee-Object -FilePath $Log -Append
  $ExitCode = $LASTEXITCODE
  Write-Log "exit code: $ExitCode"
  if ($ExitCode -ne 0) {
    throw "git failed with exit code $ExitCode"
  }
}

if (Test-Path $Log) {
  Remove-Item -Force -LiteralPath $Log
}
Set-Location $Root

$Token = Get-GitHubToken
if ([string]::IsNullOrWhiteSpace($Token)) {
  throw "GitHub token is empty."
}

$AuthBytes = [Text.Encoding]::ASCII.GetBytes("x-access-token:$Token")
$AuthHeader = "Authorization: Basic " + [Convert]::ToBase64String($AuthBytes)
$Token = $null

$CommonArgs = @(
  "-c", "safe.directory=$($Root.Path.Replace('\', '/'))",
  "-c", "http.sslBackend=openssl",
  "-c", "http.extraHeader=$AuthHeader"
)

Invoke-Git @("push", "origin", $Branch) ($CommonArgs + @("push", "origin", $Branch))
if (![string]::IsNullOrWhiteSpace($Tag)) {
  Invoke-Git @("push", "origin", $Tag) ($CommonArgs + @("push", "origin", $Tag))
}

Write-Log "Push completed."
Write-Host "Push completed."
