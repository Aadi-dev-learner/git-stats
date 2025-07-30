$ErrorActionPreference = "Stop"

# Define system and binary info
$OS = "windows"
$ARCH = "amd64"
$BinaryName = "git-stats-$OS-$ARCH.exe"
$InstallPath = "$env:USERPROFILE\git-stats.exe"

# GitHub Release URL
$URL = "https://github.com/Aadi-dev-learner/git-stats/releases/latest/download/$BinaryName"

Write-Host "⬇️  Downloading git-stats for $OS/$ARCH..."
Invoke-WebRequest -Uri $URL -OutFile $InstallPath

Write-Host "✅ Installed at $InstallPath"
Write-Host "`nTo run: `n  & '$InstallPath' -r <repo_link>"
