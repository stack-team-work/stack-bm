---
name: dev-stack-bm-start
description: >
  Start stack-bm local development services (Go backend + Vue frontend).
  Use when user says "启动项目", "start dev", "dev-start", "启动开发环境",
  "升始服务", or similar.
---

# Dev Stack BM Start

Start all local development services for the stack-bm project.

## Prerequisites
- Project root: the current working directory (where `go.mod` exists)
- MySQL running on 127.0.0.1:3306
- Go and Node.js installed

## Workflow

### Step 1: Kill existing processes on ports 8080 and 3000

**Windows (PowerShell):**
```powershell
$p8080 = Get-NetTCPConnection -LocalPort 8080 -ErrorAction SilentlyContinue | Select-Object -ExpandProperty OwningProcess -Unique
if ($p8080) { foreach ($p in $p8080) { Stop-Process -Id $p -Force } }
$p3000 = Get-NetTCPConnection -LocalPort 3000 -ErrorAction SilentlyContinue | Select-Object -ExpandProperty OwningProcess -Unique
if ($p3000) { foreach ($p in $p3000) { Stop-Process -Id $p -Force } }
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

Then wait 3 seconds and verify the backend is running by checking port 8080:
```powershell
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

Then wait 3 seconds and verify:
```powershell
$portCheck = Get-NetTCPConnection -LocalPort 3000 -ErrorAction SilentlyContinue
if ($portCheck) { Write-Host "  Frontend started on :3000" } else { Write-Host "  Frontend start failed" }
```

### Step 4: Report

```
=== Dev services started ===
  Backend:  http://localhost:8080
  Frontend: http://localhost:3000
  Login:    admin / admin123 (dev mode)
  Stop:     dev-stack-bm-stop
```
