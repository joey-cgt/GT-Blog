package main

import (
	"fmt"
	"log"
	"net/http"

	"gt-blog/backend/api"
	"gt-blog/backend/config"
	"gt-blog/backend/internal/handler"
	"gt-blog/backend/internal/middleware"
	"gt-blog/backend/internal/model"
	"gt-blog/backend/internal/repository"
)

func main() {
	// 1. 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("无法加载配置:", err)
	}

	// 2. 初始化数据库连接
	gormDB, err := repository.InitDB(cfg.Database)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	defer func() {
		if db, _ := gormDB.DB(); db != nil {
			db.Close()
		}
	}()

	// 自动迁移模型
	err = gormDB.AutoMigrate(
		&model.Profile{},
		// 其他模型...
	)
	if err != nil {
		log.Fatal("自动迁移失败:", err)
	}

	// 3. 初始化路由
	router := handler.NewRouter()

	// 4. 注册中间件
	router.Use(middleware.CORS())
	router.Use(middleware.AuthMiddleware)
	router.Use(middleware.Logger)

	// 5. 注册API路由
	router.GET("/api/profiles/:id", api.GetProfile)
	router.POST("/api/profiles/update", handler.UpdateProfile)
	// 其他路由...

	// 6. 启动服务器
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("服务器启动在 %s", serverAddr)

	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
