.PHONY: help api build run clean test docker

help: ## 显示帮助信息
	@echo "可用命令:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

api: ## 从 API 定义生成代码
	goctl api go -api api/yusi.api -dir . -style go_zero

build: ## 编译项目
	go build -o bin/yusi yusi.go

run: ## 运行服务
	go run yusi.go -f etc/yusi.yaml

clean: ## 清理编译产物
	rm -rf bin/
	go clean

test: ## 运行测试
	go test -v ./...

tidy: ## 整理依赖
	go mod tidy

fmt: ## 格式化代码
	go fmt ./...

lint: ## 代码检查
	golangci-lint run

docker-build: ## 构建 Docker 镜像
	docker build -t yusi-backend:latest .

docker-run: ## 运行 Docker 容器
	docker run -p 20611:20611 yusi-backend:latest

model: ## 从数据库生成 Model (需要修改数据库连接)
	@echo "请先修改下面命令中的数据库连接信息"
	@echo "goctl model mysql datasource -url \"root:password@tcp(127.0.0.1:3306)/yusi\" -table \"user,diary\" -dir ./model"

dev: ## 开发模式运行（带热重载，需要安装 air）
	air

install-tools: ## 安装开发工具
	go install github.com/zeromicro/go-zero/tools/goctl@latest
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "工具安装完成！"
