# Stack BM - 游戏发行后台管理系统

## 项目简介
游戏发行后台管理系统，提供管理员管理、游戏管理和游戏应用管理等基础功能。

## 技术栈
- **后端**: Go 1.25+, Gin, GORM, MySQL 5.7+
- **前端**: Vue 3, Naive UI, Vite
- **认证**: JWT + MD5(password + salt)

## 快速开始

### 环境要求
- Go 1.25+
- MySQL 5.7+
- Node.js 18+

### 数据库初始化
确保 MySQL 中存在以下数据库和表：
- `stack_bm` 库: `sys_admin`, `sys_admin_group`
- `stack_api` 库: `game`, `game_app`

```sql
-- stack_bm
CREATE TABLE `sys_admin` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  `salt` varchar(10) NOT NULL,
  `realname` varchar(50) DEFAULT '',
  `email` varchar(100) DEFAULT '',
  `phone` varchar(20) DEFAULT '',
  `group_id` int unsigned DEFAULT 0,
  `status` tinyint DEFAULT 1,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`)
);

CREATE TABLE `sys_admin_group` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `description` varchar(200) DEFAULT '',
  `status` tinyint DEFAULT 1,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

-- stack_api
CREATE TABLE `game` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `code` varchar(50) NOT NULL,
  `description` varchar(500) DEFAULT '',
  `icon` varchar(255) DEFAULT '',
  `status` tinyint DEFAULT 1,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_code` (`code`)
);

CREATE TABLE `game_app` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `game_id` int unsigned NOT NULL,
  `name` varchar(100) NOT NULL,
  `code` varchar(50) NOT NULL,
  `app_key` varchar(100) DEFAULT '',
  `description` varchar(500) DEFAULT '',
  `status` tinyint DEFAULT 1,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_code` (`code`),
  KEY `idx_game_id` (`game_id`)
);
```

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
| `DB_BM_HOST` | stack-bm 数据库主机 | 127.0.0.1 |
| `DB_BM_PORT` | stack-bm 数据库端口 | 3306 |
| `DB_BM_USER` | stack-bm 数据库用户 | root |
| `DB_BM_PASSWORD` | stack-bm 数据库密码 | root |
| `DB_BM_NAME` | stack_bm 数据库名 | stack_bm |
| `DB_API_HOST` | stack_api 数据库主机 | 127.0.0.1 |
| `DB_API_PORT` | stack_api 数据库端口 | 3306 |
| `DB_API_USER` | stack_api 数据库用户 | root |
| `DB_API_PASSWORD` | stack_api 数据库密码 | root |
| `DB_API_NAME` | stack_api 数据库名 | stack_api |
| `JWT_SECRET` | JWT 签名密钥 | - |
| `JWT_EXPIRE_HOURS` | Token 过期时间(小时) | 24 |

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
