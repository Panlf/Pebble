# Ops Workbench local server (ASCII-only content to avoid encoding issues)
# Serves the workspace folder at http://localhost:8730/ so the page can use
# the File System Access API (blocked on file:// pages by browsers).
$ErrorActionPreference = "SilentlyContinue"
$root = $PSScriptRoot
if (-not $root) { $root = Split-Path -Parent $MyInvocation.MyCommand.Path }
$port = 8730
$indexPage = "ops-workbench.html"

$listener = New-Object System.Net.HttpListener
$listener.Prefixes.Add("http://localhost:$port/")

try { $listener.Start() } catch {
  Write-Host "[ERROR] Port $port is already in use. Close the old window and retry."
  Read-Host "Press Enter to exit"
  exit 1
}

Write-Host "====================================================="
Write-Host " Ops Workbench server is running."
Write-Host " URL:  http://localhost:$port/$indexPage"
Write-Host " Keep this window OPEN while using the workbench."
Write-Host " Close this window (or press Ctrl+C) to stop."
Write-Host "====================================================="

Start-Process "http://localhost:$port/$indexPage"

$mime = @{
  ".html" = "text/html; charset=utf-8"
  ".htm"  = "text/html; charset=utf-8"
  ".js"   = "text/javascript; charset=utf-8"
  ".css"  = "text/css; charset=utf-8"
  ".json" = "application/json; charset=utf-8"
  ".png"  = "image/png"
  ".jpg"  = "image/jpeg"
  ".jpeg" = "image/jpeg"
  ".gif"  = "image/gif"
  ".svg"  = "image/svg+xml"
  ".ico"  = "image/x-icon"
  ".txt"  = "text/plain; charset=utf-8"
}

while ($listener.IsListening) {
  $ctx = $listener.GetContext()
  try {
    $path = [Uri]::UnescapeDataString($ctx.Request.Url.AbsolutePath)
    if ($path -eq "/") { $path = "/" + $indexPage }
    $rel = $path -replace "/", "\"
    $file = [IO.Path]::GetFullPath((Join-Path $root $rel))
    $rootFull = [IO.Path]::GetFullPath($root)
    if (-not $file.StartsWith($rootFull) -or -not (Test-Path -LiteralPath $file -PathType Leaf)) {
      $ctx.Response.StatusCode = 404
      $msg = [Text.Encoding]::UTF8.GetBytes("404 Not Found")
      $ctx.Response.OutputStream.Write($msg, 0, $msg.Length)
    } else {
      $ext = [IO.Path]::GetExtension($file).ToLowerInvariant()
      if ($mime.ContainsKey($ext)) { $ctx.Response.ContentType = $mime[$ext] }
      else { $ctx.Response.ContentType = "application/octet-stream" }
      $ctx.Response.Headers.Add("Cache-Control", "no-store")
      $bytes = [IO.File]::ReadAllBytes($file)
      $ctx.Response.ContentLength64 = $bytes.Length
      $ctx.Response.OutputStream.Write($bytes, 0, $bytes.Length)
    }
  } catch {
    $ctx.Response.StatusCode = 500
  } finally {
    try { $ctx.Response.Close() } catch { }
  }
}
