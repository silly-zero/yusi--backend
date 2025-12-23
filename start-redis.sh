#!/bin/bash

# 启动 Redis 服务脚本

echo "正在检查 Redis 状态..."

# 检查 Redis 是否已经在运行
if redis-cli ping > /dev/null 2>&1; then
    echo "✅ Redis 已在运行"
    redis-cli INFO server | grep redis_version
    exit 0
fi

echo "正在启动 Redis..."

# 启动 Redis（后台运行，无密码，绑定到本地）
redis-server --daemonize yes \
    --protected-mode no \
    --bind 127.0.0.1 \
    --port 6379 \
    --save 900 1 \
    --save 300 10 \
    --save 60 10000

# 等待 Redis 启动
sleep 1

# 验证 Redis 是否成功启动
if redis-cli ping > /dev/null 2>&1; then
    echo "✅ Redis 启动成功"
    redis-cli INFO server | grep redis_version
    echo "Redis 运行在 127.0.0.1:6379"
    echo "无密码访问"
else
    echo "❌ Redis 启动失败"
    exit 1
fi
