# 数据库操作练习

这个文件夹包含了SQL语句练习和Go语言数据库操作示例。

## 环境准备

使用Docker Compose启动PostgreSQL数据库（在项目根目录执行）：

```bash
cd ..  # 回到项目根目录
docker-compose up -d
```

数据库连接信息：
- 主机: localhost
- 端口: 5432
- 数据库: students, blog
- 用户名: postgres
- 密码: 123456

## 文件说明

- `main.go` - 主程序入口，运行SQLX和GORM示例
- `sqlx.go` - SQLX数据库操作示例
- `gorm.go` - GORM ORM框架示例
- `go.mod` - Go模块依赖管理
- `complete_sql_exercise.sql` - 包含所有SQL语句的完整文件，按顺序执行即可
- `README.md` - 使用说明文档

## 使用方法

### 1. 手动执行SQL练习

```bash
# 1. 启动数据库（在项目根目录执行）
cd ..  # 回到项目根目录
docker-compose up -d

# 2. 执行SQL练习
docker exec -i students_postgres psql -U postgres -d students < task3/complete_sql_exercise.sql

# 3. 查看结果
docker exec students_postgres psql -U postgres -d students -c "SELECT * FROM students;"
docker exec students_postgres psql -U postgres -d students -c "SELECT * FROM accounts;"
docker exec students_postgres psql -U postgres -d students -c "SELECT * FROM transactions;"
```

### 2：运行main.go

```bash
# 1. 启动数据库（在项目根目录执行）
cd ..  # 回到项目根目录
docker-compose up -d

# 2. 运行Go程序
cd task3
go run .

# 程序会自动执行：
# - SQLX 示例（员工和书籍查询）
# - GORM 示例（博客系统模型和关联查询）
```

## 项目结构
```
task3/
├── main.go                    # 主程序入口
├── sqlx.go                    # SQLX示例
├── gorm.go                    # GORM示例
├── go.mod                     # Go模块依赖
├── complete_sql_exercise.sql  # 完整的SQL练习文件
└── README.md                  # 使用说明

../docker-compose.yml          # 数据库Docker配置（位于项目根目录）
../init-databases.sql          # 数据库初始化脚本
```
