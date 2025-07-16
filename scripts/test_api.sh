#!/bin/bash

# API测试脚本
BASE_URL="http://localhost:8080"

echo "🧪 开始API测试..."

# 测试健康检查
echo "1. 测试健康检查..."
curl -s "$BASE_URL/health" | jq .

# 测试创建用户
echo -e "\n2. 测试创建用户..."
USER_RESPONSE=$(curl -s -X POST "$BASE_URL/api/v1/users" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "test_user",
    "email": "test@example.com",
    "password": "password123",
    "name": "测试用户"
  }')
echo $USER_RESPONSE | jq .

# 获取用户ID
USER_ID=$(echo $USER_RESPONSE | jq -r '.data.id')

# 测试获取用户列表
echo -e "\n3. 测试获取用户列表..."
curl -s "$BASE_URL/api/v1/users" | jq .

# 测试获取单个用户
echo -e "\n4. 测试获取单个用户..."
curl -s "$BASE_URL/api/v1/users/$USER_ID" | jq .

# 测试创建文章
echo -e "\n5. 测试创建文章..."
POST_RESPONSE=$(curl -s -X POST "$BASE_URL/api/v1/posts" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "测试文章",
    "content": "这是一篇测试文章的内容...",
    "summary": "测试文章摘要"
  }')
echo $POST_RESPONSE | jq .

# 获取文章ID
POST_ID=$(echo $POST_RESPONSE | jq -r '.data.id')

# 测试获取文章列表
echo -e "\n6. 测试获取文章列表..."
curl -s "$BASE_URL/api/v1/posts" | jq .

# 测试获取单个文章
echo -e "\n7. 测试获取单个文章..."
curl -s "$BASE_URL/api/v1/posts/$POST_ID" | jq .

# 测试获取用户的文章
echo -e "\n8. 测试获取用户的文章..."
curl -s "$BASE_URL/api/v1/posts/user/$USER_ID" | jq .

echo -e "\n✅ API测试完成！" 