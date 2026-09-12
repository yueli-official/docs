param(
    [ValidateSet("Up", "Down", "Status")]
    [string]$Action = "Up",
    [switch]$SkipPrepare
)

# Product entry point: all preparation and process ownership stays with Workspace.
# Docs uses isolated Identity/Account/Asset ports, independent of the Blog session.
$ErrorActionPreference = "Stop"
$workspaceRoot = (Resolve-Path (Join-Path $PSScriptRoot "../../workspace")).Path
if (-not $env:LOCAL_IDENTITY_PORT) { $env:LOCAL_IDENTITY_PORT = "8781" }
if (-not $env:LOCAL_ACCOUNT_PORT) { $env:LOCAL_ACCOUNT_PORT = "3601" }
if (-not $env:LOCAL_ASSET_PORT) { $env:LOCAL_ASSET_PORT = "8782" }
if (-not $env:LOCAL_DOCS_API_PORT) { $env:LOCAL_DOCS_API_PORT = "8086" }
if (-not $env:LOCAL_DOCS_WEB_PORT) { $env:LOCAL_DOCS_WEB_PORT = "3003" }
if (-not $env:LOCAL_DOCS_DEV_SEED) { $env:LOCAL_DOCS_DEV_SEED = "false" }

$apiBase = "http://127.0.0.1:$($env:LOCAL_DOCS_API_PORT)"
$env:GF_DOCS_PERSONALTOKENS_SITEID = "docs-main-web"
$env:GF_DOCS_PERSONALTOKENS_ALLOWHTTP = "true"
$env:GF_PAT_APPLICATIONS = ConvertTo-Json -Compress -InputObject @(@{
    id = "docs-main-web"
    name = "月离文档"
    permissionsUrl = "$apiBase/api/v1/internal/personal-token/permissions"
    audience = "docs-main-web"
})
$env:GF_PAT_ALLOWHTTP = "true"
$env:GF_ASSET_PERSONALTOKENS_AUTHORITIES = ConvertTo-Json -Compress -InputObject @{
    "docs-main-web" = "$apiBase/api/v1/personal-token/media-authorization"
}
$env:GF_ASSET_PERSONALTOKENS_ALLOWHTTP = "true"

Push-Location $workspaceRoot
try {
    & ./environments/docs-local/run.ps1 -Action $Action -Mode Isolated -SkipPrepare:$SkipPrepare
} finally {
    Pop-Location
}
