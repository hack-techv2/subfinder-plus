[CmdletBinding()]
param()

$ErrorActionPreference = 'Stop'
$repoRoot = (Resolve-Path (Join-Path $PSScriptRoot '..')).Path
$workflowDir = Join-Path $repoRoot '.github/workflows'
$ciPath = Join-Path $workflowDir 'ci.yml'
$releasePath = Join-Path $workflowDir 'subfinder-plus-release.yml'

$expected = @('ci.yml', 'subfinder-plus-release.yml')
$actual = @(Get-ChildItem -LiteralPath $workflowDir -Filter '*.yml' | Select-Object -ExpandProperty Name | Sort-Object)
$workflowDifference = @(Compare-Object -ReferenceObject $expected -DifferenceObject $actual)
if ($workflowDifference.Count -ne 0) {
    throw "Unexpected workflow set: $($actual -join ', ')"
}

$ci = Get-Content -LiteralPath $ciPath -Raw
$release = Get-Content -LiteralPath $releasePath -Raw
$readme = Get-Content -Raw (Join-Path $repoRoot 'README.md')
$maintenance = Get-Content -Raw (Join-Path $repoRoot 'docs/fork-maintenance.md')
$security = Get-Content -Raw (Join-Path $repoRoot 'SECURITY.md')
$issueConfig = Get-Content -Raw (Join-Path $repoRoot '.github/ISSUE_TEMPLATE/config.yml')
$allWorkflows = $ci + $release

if (($security + $issueConfig) -match 'github\.com/hack-techv2/subfinder-plus/security/advisories/new') {
    throw 'Security guidance must not assume private vulnerability reporting is enabled'
}
if ($issueConfig -notmatch [regex]::Escape('https://github.com/hack-techv2/subfinder-plus/security/policy')) {
    throw 'Security contact must point to the repository security policy'
}
foreach ($required in @('when the repository exposes that option', 'minimal public issue', 'private contact channel')) {
    if ($security -notmatch [regex]::Escape($required)) {
        throw "SECURITY.md is missing conditional reporting guidance: $required"
    }
}

$initialUpstreamMatches = [regex]::Matches($maintenance, '(?m)^- Initial upstream base: `([0-9a-f]{40})`\.$')
if ($initialUpstreamMatches.Count -ne 1) {
    throw 'Fork maintenance must record exactly one 40-character Initial upstream base'
}
if ($initialUpstreamMatches[0].Groups[1].Value -ne 'd0ea1029cf87ff965804fc399d12e4502c7436b2') {
    throw 'Fork maintenance changed the immutable Initial upstream base'
}

$currentUpstreamMatches = [regex]::Matches($maintenance, '(?m)^- Current upstream base: `([0-9a-f]{40})`\.$')
if ($currentUpstreamMatches.Count -ne 1) {
    throw 'Fork maintenance must record exactly one 40-character Current upstream base'
}

foreach ($required in @('go test ./...', 'go vet ./...', 'go build ./...')) {
    if ($ci -notmatch [regex]::Escape($required)) { throw "CI missing $required" }
}

foreach ($required in @('subfinder-plus-v*.*.*', 'subfinder-plus-linux-amd64', 'subfinder-plus-windows-amd64.exe', 'checksums.txt')) {
    if ($release -notmatch [regex]::Escape($required)) { throw "release missing $required" }
}

foreach ($notice in @('LICENSE.md', 'DISCLAIMER.md', 'THANKS.md')) {
    if ($release -notmatch [regex]::Escape($notice)) { throw "release missing notice $notice" }
}

foreach ($forbidden in @('dockerhub', 'slack', 'discord', 'SECURITYTRAILS_API_KEY', 'SHODAN_API_KEY', 'dabsterz', 'Dockerfile.plus')) {
    if ($allWorkflows -match $forbidden) { throw "forbidden workflow dependency: $forbidden" }
}

if ($release -notmatch 'contents:\s*write') { throw 'release missing contents: write permission' }
if ($release -notmatch '\^\[0-9\]\+\\\.\[0-9\]\+\\\.\[0-9\]\+\$') { throw 'release missing independent SemVer validation' }

foreach ($required in @(
    'Subfinder-plus',
    'ProjectDiscovery Subfinder',
    'BBOT',
    'docs/free-source-coverage.md',
    'docs/fork-maintenance.md'
)) {
    if ($readme -notmatch [regex]::Escape($required)) {
        throw "README.md is missing required documentation marker: $required"
    }
}

if ($release -match 'HACKTECH_CHANGES\.md') {
    throw 'Release workflow still references HACKTECH_CHANGES.md'
}
if ($release -notmatch 'docs/fork-maintenance\.md') {
    throw 'Release workflow does not reference docs/fork-maintenance.md'
}
if ($release -notmatch [regex]::Escape('Current upstream base: `\([0-9a-f]\{40\}\)`')) {
    throw 'Release workflow does not parse the 40-character Current upstream base'
}
if ($release -match [regex]::Escape('Initial upstream base: `\([0-9a-f]\{40\}\)`')) {
    throw 'Release workflow must not parse the immutable Initial upstream base'
}

Write-Host 'workflow contract passed'
