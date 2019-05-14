$name = "pack-repo"
$version = "0.4.3"
#$url = "http://git.greedyint.com/devops/draft-pack-repo/raw/scaffold/bin/$name-v$version-windows-amd64.zip"

if ($env:TEMP -eq $null) {
  $env:TEMP = Join-Path $env:SystemDrive 'temp'
}
$tempDir = Join-Path $env:TEMP $name
#if (![System.IO.Directory]::Exists($tempDir)) {[void][System.IO.Directory]::CreateDirectory($tempDir)}
#$file = Join-Path $env:TEMP "$name-v$version-windows-amd64.zip"

#Write-Output "Downloading $url"
#(new-object System.Net.WebClient).DownloadFile($url, $file)

$installPath = "$env:SCAFFOLD_HOME\plugins\draft-pack-repo\bin"
if (![System.IO.Directory]::Exists($installPath)) {[void][System.IO.Directory]::CreateDirectory($installPath)}
Write-Output "Preparing to install into $installPath"

$binaryPath = "$installPath\$name.exe"
Expand-Archive -Path "$file" -DestinationPath "$tempDir" -Force
if ([System.IO.File]::Exists("$binaryPath")) {[void][System.IO.File]::Delete("$binaryPath")}
Move-Item -Path "$tempDir\windows-amd64\$name.exe" -Destination "$binaryPath"

Write-Output "$name installed into $installPath\$name.exe"
