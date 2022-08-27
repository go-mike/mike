[CmdletBinding()]
param()

$Command = "go test ./..."
if ($VerbosePreference) {
    $Command = $Command + " -v"
}
Write-Host $Command -ForegroundColor Green
Invoke-Expression $Command
