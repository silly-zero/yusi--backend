package svc

import (
	"log"

	"yusi-backend/internal/config"
	"yusi-backend/internal/database"
	"yusi-backend/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware
	DB     *gorm.DB
	// Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库（自动创建数据库和表）
	db, err := database.InitDB(c.Mysql.DataSource)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	return &ServiceContext{
		Config: c,
		Auth:   middleware.NewAuthMiddleware(c.Auth.AccessSecret).Handle,
		DB:     db,
	}
}
