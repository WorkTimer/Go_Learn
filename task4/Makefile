# Makefile for Blog System

.PHONY: help build run test clean deps migrate

# 默认目标
help:
	@echo "Available commands:"
	@echo "  deps     - 安装依赖"
	@echo "  build    - 编译项目"
	@echo "  run      - 运行项目"
	@echo "  test     - 运行测试"
	@echo "  clean    - 清理编译文件"
	@echo "  migrate  - 数据库迁移"

# 安装依赖
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

# 编译项目
build:
	@echo "Building project..."
	go build -o bin/blog-server cmd/server/main.go

# 运行项目
run:
	@echo "Running project..."
	go run cmd/server/main.go

# 运行测试
test:
	@echo "Running tests..."
	go test -v ./...

# 清理编译文件
clean:
	@echo "Cleaning build files..."
	rm -rf bin/

# 数据库迁移
migrate:
	@echo "Running database migration..."
	go run cmd/server/main.go
