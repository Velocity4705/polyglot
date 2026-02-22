$ErrorActionPreference = 'Stop'

$packageName = 'polyglot'
$toolsDir = "$(Split-Path -parent $MyInvocation.MyCommand.Definition)"
$url64 = 'https://github.com/yourusername/polyglot/releases/download/v0.6.0/polyglot-windows-amd64.exe'

$packageArgs = @{
  packageName   = $packageName
  unzipLocation = $toolsDir
  fileType      = 'exe'
  url64bit      = $url64
  softwareName  = 'Polyglot*'
  checksum64    = 'REPLACE_WITH_ACTUAL_CHECKSUM'
  checksumType64= 'sha256'
  silentArgs    = ''
  validExitCodes= @(0)
}

Install-ChocolateyPackage @packageArgs

# Add to PATH
Install-ChocolateyPath "$toolsDir" 'Machine'
