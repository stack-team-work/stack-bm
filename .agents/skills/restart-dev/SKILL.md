---
name: restart-dev
description: 重启本项目（stack-bm）的开发服务：后端 go run cmd/server/main.go（:8080）与前端 web/ 下 npm run dev（:3000）。凡是用户提到"重启"、"重启后端"、"重启前端"、"restart"、"把服务跑起来"、"接口不通/页面打不开（怀疑旧进程）"、"端口被占"，即使没说"重启"二字也应使用本技能。
---

# 重启 stack-bm 前后端开发服务

环境是 Windows 10 + Git Bash。已实测：`netstat`、`cmd` 均不在 Bash PATH 中，**不要**尝试 `netstat -ano` 或 `cmd //c`；查端口、杀进程一律走 `powershell -NoProfile -Command "..."`。

支持三种范围：不带限定词 = 前后端都重启；说"只重启后端/前端" = 只处理对应一侧。

## 1. 找到并杀掉旧进程

用一条 PowerShell 同时查两个端口（`-State Listen` 只匹配监听进程）。**命令必须用单引号包裹**：双引号会被 Bash 把 PowerShell 的 `$_` 展开成 Bash 自身变量（实测踩坑）：

```bash
powershell -NoProfile -Command 'Get-NetTCPConnection -State Listen -ErrorAction SilentlyContinue | Where-Object {$_.LocalPort -in 8080,3000} | Select-Object LocalPort,OwningProcess | Format-Table -HideTableHeaders'
```

只处理后目标范围的端口（8080=后端，3000=前端）。对每个查到的 PID 执行：

```bash
powershell -NoProfile -Command "Stop-Process -Id <PID> -Force"
```

- 查询结果为空 = 该服务本来就没在跑，直接进入下一步启动，不要报错。
- 注意：`go run` 启动的 8080 监听进程是 go 编译出的临时二进制（PID ≠ go run 的 shell 进程），按端口找 PID 才可靠，不要按进程名猜。

## 2. 后台启动服务

用 Bash 工具的 `run_in_background: true` 启动，**绝不能**前台运行（会一直阻塞到超时）。两条命令相互独立，可并行发起：

后端（工作目录 = 项目根）：

```bash
go run cmd/server/main.go
```

前端（工作目录 = 项目根下的 web/）：

```bash
npm run dev
```

## 3. 等待并验证端口恢复

`go run` 首次要编译、vite 启动也要几秒，先 `wait` 5 秒左右，再用第 1 步的同一条 PowerShell 查询确认目标端口重新出现 LISTENING。

- 端口已恢复 → 重启成功。后端启动日志可通过后台任务的输出文件（Read 工具读该文件）确认有无报错；前端通常无需再看日志。
- 端口未恢复 → 读后台任务输出排查（常见：`.env` 缺失/MySQL 或 Mongo 未启动/go 编译错误），修复后按需重跑第 2 步。

## 4. 汇报

向用户报告：杀掉了哪些旧 PID、两个服务是否都已后台拉起、端口验证结果。
