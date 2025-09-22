# 个人博客系统后端

一个基于 Go 语言、Gin 框架和 GORM 库开发的个人博客系统后端，实现了博客文章的基本管理功能，包括文章的创建、读取、更新和删除（CRUD）操作，同时支持用户认证和评论功能。

## 功能特性

- ✅ 用户注册和登录（JWT认证）
- ✅ 文章CRUD操作
- ✅ 评论功能
- ✅ 分页查询
- ✅ 权限控制
- ✅ 错误处理和日志记录
- ✅ 跨域支持
- ✅ 数据库自动迁移

## 技术栈

- **语言**: Go 1.21+
- **Web框架**: Gin
- **ORM**: GORM
- **数据库**: PostgreSQL
- **认证**: JWT
- **密码加密**: bcrypt

## 项目结构

```
task4/
├── cmd/
│   └── server/
│       └── main.go              # 主程序入口
├── config/
│   └── config.go                # 配置管理
├── internal/
│   ├── handlers/                # 处理器层
│   │   ├── user_handler.go
│   │   ├── post_handler.go
│   │   └── comment_handler.go
│   ├── middleware/              # 中间件
│   │   ├── auth.go
│   │   └── cors.go
│   ├── models/                  # 数据模型
│   │   ├── user.go
│   │   ├── post.go
│   │   └── comment.go
│   ├── routes/                  # 路由配置
│   │   └── routes.go
│   ├── services/                # 服务层
│   │   ├── user_service.go
│   │   ├── post_service.go
│   │   └── comment_service.go
│   └── utils/                   # 工具函数
│       └── response.go
├── pkg/
│   ├── auth/                    # 认证相关
│   │   ├── jwt.go
│   │   └── password.go
│   ├── database/                # 数据库连接
│   │   └── database.go
│   └── logger/                  # 日志
│       └── logger.go
├── docs/                        # 文档
├── config.env                   # 环境配置
├── go.mod                       # Go模块文件
├── go.sum                       # 依赖校验文件
├── Makefile                     # 构建脚本
└── README.md                    # 项目说明
```

## 环境要求

- Go 1.21 或更高版本
- PostgreSQL 12 或更高版本
- Docker 和 Docker Compose（可选）

## 快速开始

### 1. 克隆项目

```bash
cd /Users/macos/Projects/Go_Learn/task4
```

### 2. 安装依赖

```bash
make deps
# 或者
go mod tidy
```

### 3. 配置环境

复制并编辑配置文件：

```bash
cp config.env.example config.env
```

编辑 `config.env` 文件，设置数据库连接信息：

**重要：生成新的JWT密钥**
```bash
# 生成安全的JWT密钥
openssl rand -base64 32
```

```env
# 数据库配置
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=123456
DB_NAME=blog

# JWT配置
JWT_SECRET=ktjnCkMI6GgMN3w6Nein+BFSl7YThGzlmwuomDSvkzo=

# 服务器配置
SERVER_PORT=8080
GIN_MODE=debug
```

### 4. 启动数据库

使用 Docker Compose 启动 PostgreSQL：

```bash
# 在项目根目录（Go_Learn）运行
docker-compose up -d postgres
```

### 5. 运行项目

```bash
make run
# 或者
go run cmd/server/main.go
```

服务器将在 `http://localhost:8081` 启动。

## API 文档

### 基础信息

- **Base URL**: `http://localhost:8081/api/v1`
- **Content-Type**: `application/json`

### 认证

所有需要认证的接口都需要在请求头中包含 JWT token：

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzU5MTY4MzU0LCJuYmYiOjE3NTg1NjM1NTQsImlhdCI6MTc1ODU2MzU1NH0.ui_qifxELc-oRQk3iJCurUjK_u_YGyrewjQ6m8iYnAs
```

### 用户相关接口

#### 用户注册

```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123",
  "email": "test@example.com"
}
```

#### 用户登录

```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

响应：
```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzU5MTY4MzU0LCJuYmYiOjE3NTg1NjM1NTQsImlhdCI6MTc1ODU2MzU1NH0.ui_qifxELc-oRQk3iJCurUjK_u_YGyrewjQ6m8iYnAs"
  }
}
```

#### 获取用户信息

```http
GET /api/v1/profile
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzU5MTY4MzU0LCJuYmYiOjE3NTg1NjM1NTQsImlhdCI6MTc1ODU2MzU1NH0.ui_qifxELc-oRQk3iJCurUjK_u_YGyrewjQ6m8iYnAs
```

### 文章相关接口

#### 创建文章

```http
POST /api/v1/posts
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzU5MTY4MzU0LCJuYmYiOjE3NTg1NjM1NTQsImlhdCI6MTc1ODU2MzU1NH0.ui_qifxELc-oRQk3iJCurUjK_u_YGyrewjQ6m8iYnAs
Content-Type: application/json

{
  "title": "我的第一篇文章",
  "content": "这是文章内容..."
}
```

#### 获取文章列表

```http
GET /api/v1/posts?page=1&page_size=10
```

#### 获取单个文章

```http
GET /api/v1/posts/1
```

#### 更新文章

```http
PUT /api/v1/posts/1
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzU5MTY4MzU0LCJuYmYiOjE3NTg1NjM1NTQsImlhdCI6MTc1ODU2MzU1NH0.ui_qifxELc-oRQk3iJCurUjK_u_YGyrewjQ6m8iYnAs
Content-Type: application/json

{
  "title": "更新后的标题",
  "content": "更新后的内容..."
}
```

#### 删除文章

```http
DELETE /api/v1/posts/1
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzU5MTY4MzU0LCJuYmYiOjE3NTg1NjM1NTQsImlhdCI6MTc1ODU2MzU1NH0.ui_qifxELc-oRQk3iJCurUjK_u_YGyrewjQ6m8iYnAs
```

### 评论相关接口

#### 创建评论

```http
POST /api/v1/comments
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzU5MTY4MzU0LCJuYmYiOjE3NTg1NjM1NTQsImlhdCI6MTc1ODU2MzU1NH0.ui_qifxELc-oRQk3iJCurUjK_u_YGyrewjQ6m8iYnAs
Content-Type: application/json

{
  "content": "这是一条评论",
  "post_id": 1
}
```

#### 获取文章评论

```http
GET /api/v1/posts/1/comments
```

#### 删除评论

```http
DELETE /api/v1/comments/1
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzU5MTY4MzU0LCJuYmYiOjE3NTg1NjM1NTQsImlhdCI6MTc1ODU2MzU1NH0.ui_qifxELc-oRQk3iJCurUjK_u_YGyrewjQ6m8iYnAs
```

## 测试用例

### 使用 curl 测试

1. **注册用户**：
```bash
curl -X POST http://localhost:8081/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123","email":"test@example.com"}'
```

2. **用户登录**：
```bash
curl -X POST http://localhost:8081/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'
```

3. **创建文章**（需要先登录获取token）：
```bash
curl -X POST http://localhost:8081/api/v1/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzU5MTY4MzU0LCJuYmYiOjE3NTg1NjM1NTQsImlhdCI6MTc1ODU2MzU1NH0.ui_qifxELc-oRQk3iJCurUjK_u_YGyrewjQ6m8iYnAs" \
  -d '{"title":"测试文章","content":"这是测试内容"}'
```

### 使用 Postman 测试

1. 导入 Postman 集合（如果有的话）
2. 设置环境变量 `base_url` 为 `http://localhost:8081`
3. 先调用登录接口获取 token
4. 在需要认证的请求头中添加 `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzU5MTY4MzU0LCJuYmYiOjE3NTg1NjM1NTQsImlhdCI6MTc1ODU2MzU1NH0.ui_qifxELc-oRQk3iJCurUjK_u_YGyrewjQ6m8iYnAs`

## 开发命令

```bash
# 安装依赖
make deps

# 编译项目
make build

# 运行项目
make run

# 运行测试
make test

# 清理编译文件
make clean
```

## 数据库设计

### Users 表
- `id` (主键)
- `username` (用户名，唯一)
- `password` (密码，加密存储)
- `email` (邮箱，唯一)
- `created_at` (创建时间)
- `updated_at` (更新时间)
- `deleted_at` (删除时间，软删除)

### Posts 表
- `id` (主键)
- `title` (标题)
- `content` (内容)
- `user_id` (用户ID，外键)
- `created_at` (创建时间)
- `updated_at` (更新时间)
- `deleted_at` (删除时间，软删除)

### Comments 表
- `id` (主键)
- `content` (评论内容)
- `user_id` (用户ID，外键)
- `post_id` (文章ID，外键)
- `created_at` (创建时间)
- `updated_at` (更新时间)
- `deleted_at` (删除时间，软删除)

## 错误处理

系统使用统一的错误响应格式：

```json
{
  "code": 400,
  "message": "错误描述"
}
```

常见HTTP状态码：
- `200` - 成功
- `400` - 请求参数错误
- `401` - 未认证
- `403` - 无权限
- `404` - 资源不存在
- `500` - 服务器内部错误

## 日志记录

系统使用结构化日志记录，包括：
- 请求日志
- 错误日志
- 数据库操作日志

日志输出到标准输出和标准错误流。

## 部署说明

### 生产环境配置

1. 修改 `config.env` 中的配置：
   - 设置强密码
   - 修改 JWT 密钥
   - 设置 `GIN_MODE=release`

2. 编译生产版本：
```bash
make build
```

3. 运行：
```bash
./bin/blog-server
```

### Docker 部署

可以创建 Dockerfile 和 docker-compose.yml 进行容器化部署。

## 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request



### 接口测试

使用 Postman 或其他工具对接口进行测试测，以下是测试结果：

#### 测试工具

1. **Postman Collection**
   - 导入 `docs/postman_collection.json`
   - 包含所有 API 接口的测试用例
   - 支持环境变量和自动化测试

2. **命令行测试脚本**
   - 运行 `docs/test_api.sh` 进行自动化测试
   - 包含完整的测试流程和结果验证

#### 测试结果

```bash
# 运行测试脚本
./docs/test_api.sh

# 实际输出
=== 博客系统API测试 ===

1. 健康检查...
{
  "status": "ok"
}

2. 用户注册...
{
  "code": 400,
  "message": "用户名已存在"
}

3. 用户登录...
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzU5MTY4MzU0LCJuYmYiOjE3NTg1NjM1NTQsImlhdCI6MTc1ODU2MzU1NH0.ui_qifxELc-oRQk3iJCurUjK_u_YGyrewjQ6m8iYnAs",
    "user": {
      "id": 2,
      "username": "testuser",
      "email": "test@example.com",
      "created_at": "2025-09-22T10:24:07.855733-04:00",
      "updated_at": "2025-09-22T10:24:07.855733-04:00"
    }
  }
}
Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzU5MTY4MzU0LCJuYmYiOjE3NTg1NjM1NTQsImlhdCI6MTc1ODU2MzU1NH0.ui_qifxELc-oRQk3iJCurUjK_u_YGyrewjQ6m8iYnAs

4. 获取用户信息...
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 2,
    "username": "testuser",
    "email": "test@example.com",
    "created_at": "2025-09-22T10:24:07.855733-04:00",
    "updated_at": "2025-09-22T10:24:07.855733-04:00"
  }
}

5. 创建文章...
{
  "code": 200,
  "message": "文章创建成功",
  "data": {
    "id": 3,
    "title": "我的第一篇文章",
    "content": "这是文章内容，可以包含多行文本。\n\n支持Markdown格式。",
    "user_id": 2,
    "user": {
      "id": 2,
      "username": "testuser",
      "email": "test@example.com",
      "created_at": "2025-09-22T10:24:07.855733-04:00",
      "updated_at": "2025-09-22T10:24:07.855733-04:00"
    },
    "created_at": "2025-09-22T13:47:27.516576-04:00",
    "updated_at": "2025-09-22T13:47:27.516576-04:00"
  }
}
文章ID: 3

6. 获取文章列表...
{
  "code": 200,
  "message": "success",
  "data": {
    "page": 1,
    "page_size": 10,
    "posts": [
      {
        "id": 3,
        "title": "我的第一篇文章",
        "content": "这是文章内容，可以包含多行文本。\n\n支持Markdown格式。",
        "user_id": 2,
        "user": {
          "id": 2,
          "username": "testuser",
          "email": "test@example.com",
          "created_at": "2025-09-22T10:24:07.855733-04:00",
          "updated_at": "2025-09-22T10:24:07.855733-04:00"
        },
        "created_at": "2025-09-22T13:47:27.516576-04:00",
        "updated_at": "2025-09-22T13:47:27.516576-04:00"
      },
      {
        "id": 2,
        "title": "更新后的标题",
        "content": "这是更新后的内容...",
        "user_id": 2,
        "user": {
          "id": 2,
          "username": "testuser",
          "email": "test@example.com",
          "created_at": "2025-09-22T10:24:07.855733-04:00",
          "updated_at": "2025-09-22T10:24:07.855733-04:00"
        },
        "created_at": "2025-09-22T10:24:07.93434-04:00",
        "updated_at": "2025-09-22T10:24:07.989604-04:00"
      },
      {
        "id": 1,
        "title": "测试文章",
        "content": "这是重新启动后的测试内容",
        "user_id": 1,
        "user": {
          "id": 1,
          "username": "testuser2",
          "email": "test2@example.com",
          "created_at": "2025-09-22T10:23:31.248007-04:00",
          "updated_at": "2025-09-22T10:23:31.248007-04:00"
        },
        "created_at": "2025-09-22T10:23:46.985228-04:00",
        "updated_at": "2025-09-22T10:23:46.985228-04:00"
      }
    ],
    "total": 3,
    "total_page": 1
  }
}

7. 获取单个文章...
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 3,
    "title": "我的第一篇文章",
    "content": "这是文章内容，可以包含多行文本。\n\n支持Markdown格式。",
    "user_id": 2,
    "user": {
      "id": 2,
      "username": "testuser",
      "email": "test@example.com",
      "created_at": "2025-09-22T10:24:07.855733-04:00",
      "updated_at": "2025-09-22T10:24:07.855733-04:00"
    },
    "created_at": "2025-09-22T13:47:27.516576-04:00",
    "updated_at": "2025-09-22T13:47:27.516576-04:00"
  }
}

8. 创建评论...
{
  "code": 200,
  "message": "评论创建成功",
  "data": {
    "id": 3,
    "content": "这是一条评论",
    "user_id": 2,
    "post_id": 3,
    "user": {
      "id": 2,
      "username": "testuser",
      "email": "test@example.com",
      "created_at": "2025-09-22T10:24:07.855733-04:00",
      "updated_at": "2025-09-22T10:24:07.855733-04:00"
    },
    "created_at": "2025-09-22T13:47:27.553605-04:00",
    "updated_at": "2025-09-22T13:47:27.553605-04:00"
  }
}

9. 获取文章评论...
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 3,
      "content": "这是一条评论",
      "user_id": 2,
      "post_id": 3,
      "user": {
        "id": 2,
        "username": "testuser",
        "email": "test@example.com",
        "created_at": "2025-09-22T10:24:07.855733-04:00",
        "updated_at": "2025-09-22T10:24:07.855733-04:00"
      },
      "created_at": "2025-09-22T13:47:27.553605-04:00",
      "updated_at": "2025-09-22T13:47:27.553605-04:00"
    }
  ]
}

10. 更新文章...
{
  "code": 200,
  "message": "文章更新成功",
  "data": {
    "id": 3,
    "title": "更新后的标题",
    "content": "这是更新后的内容...",
    "user_id": 2,
    "user": {
      "id": 2,
      "username": "testuser",
      "email": "test@example.com",
      "created_at": "2025-09-22T10:24:07.855733-04:00",
      "updated_at": "2025-09-22T10:24:07.855733-04:00"
    },
    "created_at": "2025-09-22T13:47:27.516576-04:00",
    "updated_at": "2025-09-22T13:47:27.57331-04:00"
  }
}

=== 测试完成 ===
```



