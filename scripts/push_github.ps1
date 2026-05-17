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
  $Gcm = Get-Command git-credential-manager.exe -ErrorAction SilentlyContinue
  if (!$Gcm) {
    $KnownGcm = "C:\Users\zhangjingyao\scoop\apps\git\2.53.0.2\mingw64\bin\git-credential-manager.exe"
    if (Test-Path $KnownGcm) {
      $GcmPath = $KnownGcm
    } else {
      return ""
    }
  } else {
    $GcmPath = $Gcm.Source
  }
  $InputText = "protocol=https`nhost=github.com`npath=$Owner/$Repo.git`n`n"
  $Cred = $InputText | & $GcmPath get
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

function Get-GitHubTokenFromGh {
  $Gh = Get-Command gh.exe -ErrorAction SilentlyContinue
  if (!$Gh) {
    return ""
  }
  $Token = & $Gh.Source auth token 2>$null
  if ($LASTEXITCODE -ne 0 -or !$Token) {
    return ""
  }
  return (($Token | Out-String).Trim())
}

function Get-GitHubToken {
  if (!$PromptToken) {
    $Token = Get-GitHubTokenFromGh
    if (![string]::IsNullOrWhiteSpace($Token)) {
      Write-Log "Using GitHub token from gh auth."
      return $Token
    }
    $Token = Get-GitHubTokenFromGcm
    if (![string]::IsNullOrWhiteSpace($Token)) {
      Write-Log "Using GitHub token from Git Credential Manager."
      return $Token
    }
  }
  Write-Log "No stored GitHub credential found; prompting for token."
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
