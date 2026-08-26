---
name: dev-stack-bm-restart
description: >
  Restart stack-bm local development services (Go backend + Vue frontend).
  Use when user says "restart dev", "重启项目", "restart", "重启开发环境",
  or similar.
---

# Dev Stack BM Restart

Restart all local development services for the stack-bm project
(stop any running instances, then start fresh).

## Prerequisites
- Project root: the current working directory (where `go.mod` exists)
- MySQL running on 127.0.0.1:3306
- Go and Node.js installed

## Workflow

### Step 1: Stop existing services

**Windows (PowerShell):**
```powershell
Write-Host "=== Stopping stack-bm dev services ===" -ForegroundColor Cyan

$p8080 = Get-NetTCPConnection -LocalPort 8080 -ErrorAction SilentlyContinue | Select-Object -ExpandProperty OwningProcess -Unique
if ($p8080) {
    foreach ($p in $p8080) { Write-Host "  Stopping backend (PID: $p)"; Stop-Process -Id $p -Force }
    Write-Host "  Backend stopped."
} else { Write-Host "  No process on port 8080" }

$p3000 = Get-NetTCPConnection -LocalPort 3000 -ErrorAction SilentlyContinue | Select-Object -ExpandProperty OwningProcess -Unique
if ($p3000) {
    foreach ($p in $p3000) { Write-Host "  Stopping frontend (PID: $p)"; Stop-Process -Id $p -Force }
    Write-Host "  Frontend stopped."
} else { Write-Host "  No process on port 3000" }

Get-Process -Name "go" -ErrorAction SilentlyContinue | Where-Object { $_.MainWindowTitle -eq "" } | Stop-Process -Force
Get-Process -Name "node" -ErrorAction SilentlyContinue | Where-Object { $_.CommandLine -like "*vite*" -and $_.CommandLine -like "*stack-bm*" } | Stop-Process -Force -ErrorAction SilentlyContinue
```

**Linux/Mac:**
```bash
kill $(lsof -ti:8080) 2>/dev/null && echo "  Backend stopped" || echo "  No process on port 8080"
kill $(lsof -ti:3000) 2>/dev/null && echo "  Frontend stopped" || echo "  No process on port 3000"
```

### Step 2: Start Go backend

Run in background at project root.

**Windows (PowerShell):**
```powershell
Start-Process -WindowStyle Hidden -FilePath "go" -ArgumentList "run", "cmd/server/main.go" -WorkingDirectory "{project_root}"
```

**Linux/Mac:**
```bash
cd {project_root} && nohup go run cmd/server/main.go > /dev/null 2>&1 &
```

Then wait 3 seconds and verify port 8080:
```powershell
Start-Sleep -Seconds 3
$portCheck = Get-NetTCPConnection -LocalPort 8080 -ErrorAction SilentlyContinue
if ($portCheck) { Write-Host "  Backend started on :8080" } else { Write-Host "  Backend start failed" }
```

### Step 3: Start Vue frontend

**Windows (PowerShell):**
```powershell
Start-Process -WindowStyle Hidden -FilePath "cmd" -ArgumentList "/c", "npm run dev" -WorkingDirectory "{project_root}\web"
```

**Linux/Mac:**
```bash
cd {project_root}/web && nohup npm run dev > /dev/null 2>&1 &
```

Then wait 3 seconds and verify port 3000:
```powershell
Start-Sleep -Seconds 3
$portCheck = Get-NetTCPConnection -LocalPort 3000 -ErrorAction SilentlyContinue
if ($portCheck) { Write-Host "  Frontend started on :3000" } else { Write-Host "  Frontend start failed" }
```

### Step 4: Report

```
=== Dev services restarted ===
  Backend:  http://localhost:8080
  Frontend: http://localhost:3000
  Login:    use a manually-created sys_admin account (no auto-create)
```