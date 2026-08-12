# AGENTS.md

## Dev commands

```bash
# Backend — run from project root
go run cmd/server/main.go         # start backend (:8080)
go build ./...                     # verify all packages compile

# Frontend — run from web/
npm run dev                        # start frontend (:3000)
npm run build                      # production build check
```

Dev mode (`.env` `SERVER_MODE=dev`) only affects Gin's run mode; the default admin account must be created manually (no auto-create on login).

## Architecture

Each table prefix is an independent module under 4 layers:

```
internal/
  handler/
    game/          package game          GameHandler
    gameapp/       package gameapp       GameAppHandler
    gamecp/        package gamecp        GameCpHandler
    gametag/       package gametag       GameTagHandler
    gamevariable/  package gamevariable  GameVariableHandler
    gameplatform/  package gameplatform  GamePlatformHandler
    sysadmin/      package sysadmin      SysAdminHandler
    sysadmingroup/ package sysadmingroup SysAdminGroupHandler
    syslog/        package syslog        SysLogHandler
    sysmenu/       package sysmenu       SysMenuHandler
    media/         package media         MediaHandler
    mediasub/      package mediasub      MediaSubHandler
    handler/       package handler       AuthHandler, DashboardHandler
  model/           (mirrors handler/ structure)
  repository/      (mirrors handler/ structure)
  service/         (mirrors handler/ structure)
    auth.go        package service       AuthService
  router/          package router        central routing
```

**Same-name package conflict**: Within a service or handler, model and repo/service share the same package name. Use aliases:

```go
// In service/gameapp/gameapp.go:
import (
    "stack-bm/internal/model/gameapp"
    gameappRepo "stack-bm/internal/repository/gameapp"
)

// In handler/gameapp/gameapp.go:
import (
    "stack-bm/internal/model/gameapp"
    gameappSvc "stack-bm/internal/service/gameapp"
)

// In router.go, each handler package has a unique name — no aliases needed:
import (
    "stack-bm/internal/handler/game"
    "stack-bm/internal/handler/gameapp"
    "stack-bm/internal/handler/media"
)
```

## Three databases

| Variable prefix | DB name | Connection var | Contains |
|-----------------|---------|---------------|----------|
| `DB_BM_*` | `stack_bm` | `database.DBBM` | sys_admin, sys_admin_group, sys_logs, sys_menu |
| `DB_API_*` | `stack_api` | `database.DBApi` | game, game_app, game_cp, game_tags, game_variable, game_platform, users, user_orders, etc. |
| `DB_MKT_*` | `stack_mkt` | `database.DBMkt` | media, media_sub |

New module uses the correct `database.DB*` for its table's database.

**Local connection strings** (from `.env`):

| Database | DSN |
|----------|-----|
| `stack_bm` | `root:root@tcp(127.0.0.1:3306)/stack_bm?charset=utf8mb4&parseTime=True&loc=Local` |
| `stack_api` | `root:root@tcp(127.0.0.1:3306)/stack_api?charset=utf8mb4&parseTime=True&loc=Local` |
| `stack_mkt` | `root:root@tcp(127.0.0.1:3306)/stack_mkt?charset=utf8mb4&parseTime=True&loc=Local` |

## Model quirks

- **Timestamps**: All tables use `int(11)` Unix timestamps (seconds), NOT `time.Time`. GORM tags: `autoCreateTime` / `autoUpdateTime`.
- **No soft delete**: Actual tables lack `deleted_at` columns — do not add `gorm.DeletedAt` to models.
- **Password JSON binding**: `Password` field uses `json:"password,omitempty"` (not `json:"-"`) so Gin's `ShouldBindJSON` can read it. Salt stays `json:"-"`.
- **AppKey/AppSecret**: Auto-generated in `GameAppService.Create()` via `utils.RandomString()`. Not updatable — `Update()` skips these fields.

## API conventions

- **All endpoints use POST**, including list/detail queries. No GET.
- **Response format**: `{ code: 0, data: ..., msg: "success" }`
- **Paginated**: `{ code: 0, data: { list: [...], total: N }, msg: "success" }`

## Adding a new module

1. `internal/model/<table-prefix>/` — struct + TableName(), `package <table-prefix>`
2. `internal/repository/<table-prefix>/` — CRUD methods, uses correct `database.DB*`, `package <table-prefix>`
3. `internal/service/<table-prefix>/` — business logic, import repo with alias `<prefix>Repo`
4. `internal/handler/<table-prefix>/` — HTTP handler, import service with alias `<prefix>Svc`
5. `internal/router/router.go` — register handler and routes

Frontend: add API functions in `web/src/api/<domain>.js`, create Vue page under `web/src/views/<domain>/`, add route in `web/src/router/index.js`, add menu entry in `web/src/layouts/MainLayout.vue`.

## Frontend gotchas

- **Naive UI `n-form-item-gi` + `n-switch`**: Does not reliably bind `v-model:value`. Use `n-form-item` directly instead.
- **All API calls**: Use `request.post(url, data)` from `src/utils/request.js`. The `request.js` interceptor expects `res.data.code === 0`.
- **Naive UI select**: `n-select` uses `{ label, value }` option format. `n-input-number` should be replaced with `n-select` when picking from a related table.
- **Composables**: Use `useTable(fetchFn)` for list pages (pagination, search, loading). Use `useModal()` for modal form pages (create/edit, submit, delete). Both in `src/composables/`.

## Frontend structure

```
web/src/
├── api/
│   ├── index.js      login, getUserInfo
│   ├── system.js     admin, adminGroup
│   ├── game.js       game, gameApp, gameCp, gamePlatform, gameTag, gameVariable
│   └── mkt.js        media, mediaSub
├── composables/
│   ├── useTable.js   pagination + search + load
│   └── useModal.js   form open/edit/submit/delete
├── views/
│   ├── system/       SysAdmin, SysAdminGroup, SysLogs, SysMenu, Dashboard
│   ├── game/         Game, GameApp, GameAppForm, GameCp, GameTag, GameVariable, GamePlatform
│   ├── mkt/          Media, MediaSub
│   ├── Login.vue
│   └── layouts/      MainLayout.vue
```
