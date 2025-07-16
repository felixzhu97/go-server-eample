# Go Server Example

一个使用 Go 语言构建的完整 Web 服务器示例，包含用户管理和文章管理功能。

## 功能特性

- 🚀 基于 Gin 框架的 HTTP 服务器
- 📊 SQLite 数据库集成（使用 GORM）
- 👥 用户管理（CRUD 操作）
- 📝 文章管理（CRUD 操作）
- 🔒 中间件支持（日志、CORS、恢复、速率限制）
- 📝 结构化日志记录
- ⚙️ 环境变量配置管理

## 项目结构

```
go-server-example/
├── main.go                 # 主程序入口
├── go.mod                  # Go模块文件
├── env.example             # 环境变量示例
├── config/                 # 配置管理
│   └── config.go
├── database/               # 数据库相关
│   └── database.go
├── models/                 # 数据模型
│   ├── user.go
│   └── post.go
├── handlers/               # HTTP处理器
│   ├── user_handler.go
│   └── post_handler.go
├── middleware/             # 中间件
│   └── middleware.go
├── routes/                 # 路由配置
│   └── routes.go
└── utils/                  # 工具函数
    └── logger/
        └── logger.go
```

## 快速开始

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 配置环境变量

复制环境变量示例文件：

```bash
cp env.example .env
```

根据需要修改 `.env` 文件中的配置。

### 3. 运行服务器

```bash
go run main.go
```

服务器将在 `http://localhost:8080` 启动。

## API 接口

### 健康检查

- `GET /health` - 服务器健康状态

### 用户管理

- `GET /api/v1/users` - 获取用户列表
- `GET /api/v1/users/:id` - 获取单个用户
- `POST /api/v1/users` - 创建用户
- `PUT /api/v1/users/:id` - 更新用户
- `DELETE /api/v1/users/:id` - 删除用户

### 文章管理

- `GET /api/v1/posts` - 获取文章列表
- `GET /api/v1/posts/:id` - 获取单个文章
- `POST /api/v1/posts` - 创建文章
- `PUT /api/v1/posts/:id` - 更新文章
- `DELETE /api/v1/posts/:id` - 删除文章
- `GET /api/v1/posts/user/:user_id` - 获取用户的文章

## 示例请求

### 创建用户

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "password123",
    "name": "John Doe"
  }'
```

### 创建文章

```bash
curl -X POST http://localhost:8080/api/v1/posts \
  -H "Content-Type: application/json" \
  -d '{
    "title": "我的第一篇文章",
    "content": "这是文章内容...",
    "summary": "文章摘要"
  }'
```

## 环境变量

| 变量名         | 描述           | 默认值        |
| -------------- | -------------- | ------------- |
| `PORT`         | 服务器端口     | `8080`        |
| `ENVIRONMENT`  | 运行环境       | `development` |
| `DATABASE_URL` | 数据库连接 URL | `app.db`      |
| `LOG_LEVEL`    | 日志级别       | `info`        |

## 开发

### 添加新的模型

1. 在 `models/` 目录下创建新的模型文件
2. 在 `database/database.go` 中添加模型到自动迁移列表
3. 创建对应的处理器和路由

### 添加新的中间件

在 `middleware/middleware.go` 中添加新的中间件函数，然后在 `routes/routes.go` 中注册使用。

## 许可证

MIT License
