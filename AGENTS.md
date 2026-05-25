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

Dev mode (`.env` `SERVER_MODE=dev`) auto-creates `admin/admin123` on first login if no admin exists.

## Architecture

```
internal/
  handler/game/    package game    GameHandler, GameAppHandler, GameCpHandler
  handler/sys/     package sys     SysAdminHandler, SysAdminGroupHandler
  handler/         package handler AuthHandler
  model/game/      package game    Game, GameApp, GameCp
  model/sys/       package sys     SysAdmin, SysAdminGroup
  repository/game/ package game    GameRepository, GameAppRepository, GameCpRepository
  repository/sys/  package sys     SysAdminRepository, SysAdminGroupRepository
  service/game/    package game    GameService, GameAppService, GameCpService
  service/sys/     package sys     SysAdminService, SysAdminGroupService
  service/         package service AuthService
  router/          package router  central routing
```

**Same-name package conflict**: `handler/game`, `model/game`, `service/game`, `repository/game` all declare `package game`. When importing two same-named packages in one file, use aliases:

```go
import (
    gameRepo "stack-bm/internal/repository/game"
    gameSvc  "stack-bm/internal/service/game"
    "stack-bm/internal/model/game"    // used as: game.Game, game.GameApp
)
```

## Two databases

| Variable prefix | DB name | Connection var | Contains |
|-----------------|---------|---------------|----------|
| `DB_BM_*` | `stack_bm` | `database.DBBM` | sys_admin, sys_admin_group |
| `DB_API_*` | `stack_api` | `database.DBApi` | game, game_app, game_cp, users, user_orders, etc. |

New module `Foo` in `stack_api` uses `database.DBApi`; in `stack_bm` uses `database.DBBM`.

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

Each module needs files in 4 directories + router entry:

1. `internal/model/<domain>/` — struct + TableName()
2. `internal/repository/<domain>/repository.go` — CRUD methods, uses correct `database.DB*`
3. `internal/service/<domain>/service.go` — business logic
4. `internal/handler/<domain>/handler.go` — HTTP handler with `ShouldBindJSON`
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
│   └── game.js       game, gameApp, gameCp
├── composables/
│   ├── useTable.js   pagination + search + load
│   └── useModal.js   form open/edit/submit/delete
├── views/
│   ├── system/       SysAdmin, SysAdminGroup
│   ├── game/         Game, GameApp, GameCp, GameAppForm
│   ├── Login.vue
│   └── layouts/      MainLayout.vue
```
