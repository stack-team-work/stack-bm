# Stack BM - 游戏发行后台管理系统

## 项目简介
游戏发行后台管理系统，提供管理员管理、游戏管理、游戏应用管理，以及 B站(BiliAds) / 快手(KsAds) 的广告模板、定向包模板、标题包模板等投放配置管理。

## 技术栈
- **后端**: Go 1.25+, Gin, GORM, MySQL 5.7+, MongoDB
- **前端**: Vue 3, Naive UI, Vite
- **认证**: JWT + MD5(password + salt)

## 快速开始

### 环境要求
- Go 1.25+
- MySQL 5.7+
- MongoDB（用于 B站/快手 模板，默认 127.0.0.1:27017，未安装时后端可正常启动，仅模板相关接口不可用）
- Node.js 18+

### 数据库初始化
共 3 个 MySQL 库 + 1 个 MongoDB 库，表结构以 `migrations/` 下 SQL 为准：

| 数据库 | 说明 | 主要表 |
|--------|------|--------|
| `stack_bm` | 后台管理 | sys_admin, sys_admin_group, sys_logs, sys_menu |
| `stack_sdk` | SDK / 游戏 | game, game_app, game_app_template, game_gift, game_voucher 等 |
| `stack_mkt` | 发行渠道 | media, media_sub, media_agent, media_application 等 |
| `channel_template`(MongoDB) | 投放模板 | bili_ad_template, bili_audience_template, bili_title_template, ks_ad_template, ks_audience_template, ks_title_template |

MySQL 表结构执行：

```bash
mysql -uroot -p < migrations/stack_bm.sql
mysql -uroot -p < migrations/stack_sdk.sql
mysql -uroot -p < migrations/stack_mkt.sql
```

MongoDB 集合由程序自动创建，无需手动建表；如从参考项目迁移历史数据，可将 `channel_template` 库下 `bili_*_template` / `ks_*_template` 集合直接导入。

### 启动后端服务

```bash
# 1. 安装依赖
go mod tidy

# 2. 首次运行需复制配置（根据需要修改数据库连接信息）
cp .env.example .env

# 3. 启动开发服务器
go run cmd/server/main.go
```

### 启动前端服务

```bash
cd web

# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 构建生产版本
npm run build
```

### 访问系统
- 前端: http://localhost:3000
- 后端: http://localhost:8080
- 管理员账号需手动在 `sys_admin` 表中创建

## API 接口列表

所有接口统一使用 POST 方法。

### 认证
| 接口 | 说明 | 认证 |
|------|------|------|
| /api/login | 登录 | 否 |

### 管理员管理
| 接口 | 说明 |
|------|------|
| /api/admin/create | 创建管理员 |
| /api/admin/list | 管理员列表 |
| /api/admin/detail/:id | 管理员详情 |
| /api/admin/update/:id | 更新管理员 |
| /api/admin/delete/:id | 删除管理员 |

### 管理员分组
| 接口 | 说明 |
|------|------|
| /api/admin-group/create | 创建分组 |
| /api/admin-group/list | 分组列表 |
| /api/admin-group/all | 全部分组 |
| /api/admin-group/detail/:id | 分组详情 |
| /api/admin-group/update/:id | 更新分组 |
| /api/admin-group/delete/:id | 删除分组 |

### 游戏管理
| 接口 | 说明 |
|------|------|
| /api/game/create | 创建游戏 |
| /api/game/list | 游戏列表 |
| /api/game/all | 全部游戏 |
| /api/game/detail/:id | 游戏详情 |
| /api/game/update/:id | 更新游戏 |
| /api/game/delete/:id | 删除游戏 |

### 游戏应用管理
| 接口 | 说明 |
|------|------|
| /api/game-app/create | 创建应用 |
| /api/game-app/list | 应用列表 |
| /api/game-app/detail/:id | 应用详情 |
| /api/game-app/update/:id | 更新应用 |
| /api/game-app/delete/:id | 删除应用 |

### B站模板（数据存 MongoDB channel_template.bili_*_template）
三类模板接口一致，以广告模板为例（定向包、标题包同理，前缀分别为 `bili-audience-template` / `bili-title-template`）：

| 接口 | 说明 |
|------|------|
| /api/bili-ad-template/create | 创建广告模板 |
| /api/bili-ad-template/list | 模板列表（分页 + 名称模糊） |
| /api/bili-ad-template/detail/:id | 模板详情 |
| /api/bili-ad-template/update/:id | 更新模板 |
| /api/bili-ad-template/delete/:id | 删除模板（软删 display=0） |
| /api/bili-ad-template/copy | 复制模板 |

### 快手模板（数据存 MongoDB channel_template.ks_*_template）
同 B站，前缀分别为 `ks-ad-template` / `ks-audience-template` / `ks-title-template`。

### 字典
| 接口 | 说明 |
|------|------|
| /api/dict | 全部字典（状态、游戏类型、B站/快手枚举等） |
| /api/dict/:key | 指定字典 |

## 部署

### 交叉编译
```bash
sh build.sh
```
编译产物在 `build/` 目录。

### Linux 部署
```bash
sh deploy.sh
```

### Windows 编译
```bash
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o build/stack-bm.exe cmd/server/main.go
```

## 配置说明

配置文件 `.env` 位于项目根目录（首次从 `.env.example` 复制）：

| 变量 | 说明 | 默认值 |
|------|------|--------|
| `APP_ENV` | 运行环境 | dev |
| `SERVER_PORT` | 服务端口 | 8080 |
| `SERVER_MODE` | 运行模式(dev/prod) | dev |
| `DB_BM_HOST` | stack_bm 数据库主机 | 127.0.0.1 |
| `DB_BM_PORT` | stack_bm 数据库端口 | 3306 |
| `DB_BM_USER` | stack_bm 数据库用户 | root |
| `DB_BM_PASSWORD` | stack_bm 数据库密码 | root |
| `DB_BM_NAME` | stack_bm 数据库名 | stack_bm |
| `DB_SDK_HOST` | stack_sdk 数据库主机 | 127.0.0.1 |
| `DB_SDK_PORT` | stack_sdk 数据库端口 | 3306 |
| `DB_SDK_USER` | stack_sdk 数据库用户 | root |
| `DB_SDK_PASSWORD` | stack_sdk 数据库密码 | root |
| `DB_SDK_NAME` | stack_sdk 数据库名 | stack_sdk |
| `DB_MKT_HOST` | stack_mkt 数据库主机 | 127.0.0.1 |
| `DB_MKT_PORT` | stack_mkt 数据库端口 | 3306 |
| `DB_MKT_USER` | stack_mkt 数据库用户 | root |
| `DB_MKT_PASSWORD` | stack_mkt 数据库密码 | root |
| `DB_MKT_NAME` | stack_mkt 数据库名 | stack_mkt |
| `JWT_SECRET` | JWT 签名密钥 | - |
| `JWT_EXPIRE_HOURS` | Token 过期时间(小时) | 24 |
| `MONGO_URI` | MongoDB 连接串 | mongodb://127.0.0.1:27017 |
| `MONGO_CHANNEL_DB` | 投放模板 MongoDB 库名 | channel_template |

## 项目结构

```
stack-bm/
├── cmd/server/          # 程序入口
├── internal/
│   ├── config/          # 配置管理
│   ├── database/        # 数据库初始化
│   ├── model/           # 数据模型
│   ├── repository/      # 数据访问层
│   ├── service/         # 业务逻辑层
│   ├── handler/         # HTTP处理层
│   ├── middleware/       # 中间件
│   └── router/          # 路由注册
├── pkg/
│   ├── jwt/             # JWT工具
│   ├── response/        # 响应封装
│   └── utils/           # 工具函数
├── build.sh             # 交叉编译脚本
├── deploy.sh            # 部署脚本
├── stack-bm.service     # systemd 服务文件
├── .env.example         # 配置模板
└── web/                 # 前端项目
```
