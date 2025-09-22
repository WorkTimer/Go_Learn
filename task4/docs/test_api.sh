#!/bin/bash

# 博客系统API测试脚本
# 使用方法: ./test_api.sh

BASE_URL="http://localhost:8081/api/v1"
TOKEN=""

echo "=== 博客系统API测试 ==="
echo ""

# 健康检查
echo "1. 健康检查..."
curl -s http://localhost:8081/health | jq .
echo ""

# 用户注册
echo "2. 用户注册..."
REGISTER_RESPONSE=$(curl -s -X POST $BASE_URL/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123",
    "email": "test@example.com"
  }')
echo $REGISTER_RESPONSE | jq .
echo ""

# 用户登录
echo "3. 用户登录..."
LOGIN_RESPONSE=$(curl -s -X POST $BASE_URL/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }')
echo $LOGIN_RESPONSE | jq .

# 提取token
TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.data.token')
echo "Token: $TOKEN"
echo ""

# 获取用户信息
echo "4. 获取用户信息..."
curl -s -X GET $BASE_URL/profile \
  -H "Authorization: Bearer $TOKEN" | jq .
echo ""

# 创建文章
echo "5. 创建文章..."
POST_RESPONSE=$(curl -s -X POST $BASE_URL/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "我的第一篇文章",
    "content": "这是文章内容，可以包含多行文本。\n\n支持Markdown格式。"
  }')
echo $POST_RESPONSE | jq .

# 提取文章ID
POST_ID=$(echo $POST_RESPONSE | jq -r '.data.id')
echo "文章ID: $POST_ID"
echo ""

# 获取文章列表
echo "6. 获取文章列表..."
curl -s -X GET "$BASE_URL/posts?page=1&page_size=10" | jq .
echo ""

# 获取单个文章
echo "7. 获取单个文章..."
curl -s -X GET $BASE_URL/posts/$POST_ID | jq .
echo ""

# 创建评论
echo "8. 创建评论..."
COMMENT_RESPONSE=$(curl -s -X POST $BASE_URL/comments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d "{
    \"content\": \"这是一条评论\",
    \"post_id\": $POST_ID
  }")
echo $COMMENT_RESPONSE | jq .
echo ""

# 获取文章评论
echo "9. 获取文章评论..."
curl -s -X GET $BASE_URL/posts/$POST_ID/comments | jq .
echo ""

# 更新文章
echo "10. 更新文章..."
curl -s -X PUT $BASE_URL/posts/$POST_ID \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "更新后的标题",
    "content": "这是更新后的内容..."
  }' | jq .
echo ""

echo "=== 测试完成 ==="
