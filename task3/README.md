# SQL语句练习

这个文件夹包含了两个SQL练习题目的实现。

## 环境准备

使用Docker Compose启动PostgreSQL数据库：

```bash
docker-compose up -d
```

数据库连接信息：
- 主机: localhost
- 端口: 5432
- 数据库: students
- 用户名: postgres
- 密码: 123456

## 文件说明

- `complete_sql_exercise.sql` - 包含所有SQL语句的完整文件，按顺序执行即可
- `docker-compose.yml` - PostgreSQL数据库Docker配置
- `README.md` - 使用说明文档

## 使用方法

```bash
# 1. 启动数据库
docker-compose up -d

# 2. 执行SQL练习
docker exec -i students_postgres psql -U postgres -d students < complete_sql_exercise.sql

# 3. 查看结果
docker exec students_postgres psql -U postgres -d students -c "SELECT * FROM students;"
docker exec students_postgres psql -U postgres -d students -c "SELECT * FROM accounts;"
docker exec students_postgres psql -U postgres -d students -c "SELECT * FROM transactions;"
```

## 题目内容

### 题目1：基本CRUD操作
- 创建students表（id, name, age, grade字段）
- 插入学生记录（张三，20岁，三年级）
- 查询年龄大于18岁的学生
- 更新学生年级（张三改为四年级）
- 删除年龄小于15岁的学生

### 题目2：事务语句
- 创建accounts表和transactions表
- 实现账户间转账功能（从账户A向账户B转账100元）
- 包含余额检查和事务回滚机制
- 记录转账交易信息

## 项目结构
```
task3/
├── complete_sql_exercise.sql  # 完整的SQL练习文件
├── docker-compose.yml         # 数据库Docker配置
└── README.md                  # 使用说明
```
