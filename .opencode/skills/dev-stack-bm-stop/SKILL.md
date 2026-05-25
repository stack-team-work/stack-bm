---
name: dev-stack-bm-stop
description: >
  Stop stack-bm local development services. Kill all backend and frontend processes.
  Use when user says "关闭项目", "stop dev", "dev-stop", "停止开发环境",
  "停止服务", or similar.
---

# Dev Stack BM Stop

Stop all local development services for the stack-bm project.

## Workflow

### Step 1: Kill processes on target ports

**Windows (PowerShell):**
```powershell
Write-Host "=== Stopping stack-bm dev services ===" -ForegroundColor Cyan

$p8080 = Get-NetTCPConnection -LocalPort 8080 -ErrorAction SilentlyContinue | Select-Object -ExpandProperty OwningProcess -Unique
if ($p8080) {
    foreach ($p in $p8080) {
        Write-Host "  Stopping backend on port 8080 (PID: $p)..."
        Stop-Process -Id $p -Force
    }
    Write-Host "  Backend stopped."
} else {
    Write-Host "  No process on port 8080"
}

$p3000 = Get-NetTCPConnection -LocalPort 3000 -ErrorAction SilentlyContinue | Select-Object -ExpandProperty OwningProcess -Unique
if ($p3000) {
    foreach ($p in $p3000) {
        Write-Host "  Stopping frontend on port 3000 (PID: $p)..."
        Stop-Process -Id $p -Force
    }
    Write-Host "  Frontend stopped."
} else {
    Write-Host "  No process on port 3000"
}
```

**Linux/Mac:**
```bash
echo "=== Stopping stack-bm dev services ==="
kill $(lsof -ti:8080) 2>/dev/null && echo "  Backend stopped" || echo "  No process on port 8080"
kill $(lsof -ti:3000) 2>/dev/null && echo "  Frontend stopped" || echo "  No process on port 3000"
```

### Step 2: Cleanup lingering processes

**Windows (PowerShell):**
```powershell
Get-Process -Name "go" -ErrorAction SilentlyContinue | Where-Object { $_.Path -like "*stack-bm*" -or $_.MainWindowTitle -eq "" } | Stop-Process -Force
Get-Process -Name "node" -ErrorAction SilentlyContinue | Where-Object { $_.CommandLine -like "*vite*" -and $_.CommandLine -like "*stack-bm*" } | Stop-Process -Force -ErrorAction SilentlyContinue
```

### Step 3: Verify ports are free

```powershell
$p8080 = Get-NetTCPConnection -LocalPort 8080 -ErrorAction SilentlyContinue
$p3000 = Get-NetTCPConnection -LocalPort 3000 -ErrorAction SilentlyContinue
if (-not $p8080 -and -not $p3000) {
    Write-Host "=== All services stopped ===" -ForegroundColor Green
} else {
    Write-Host "WARNING: Some ports may still be occupied" -ForegroundColor Yellow
}
```
