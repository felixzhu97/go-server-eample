package main

import (
	"log"
	"os"

	"go-server-example/config"
	"go-server-example/database"
	"go-server-example/routes"
	"go-server-example/utils/logger"
)

func main() {
	// 加载环境变量
	if err := config.LoadEnv(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// 初始化日志
	logger.Init()

	// 初始化数据库
	if err := database.Init(); err != nil {
		logger.Error("Failed to initialize database:", err)
		os.Exit(1)
	}

	// 设置路由
	router := routes.SetupRoutes()

	// 获取端口
	port := config.GetPort()
	if port == "" {
		port = "8080"
	}

	logger.Info("Server starting on port:", port)
	
	// 启动服务器
	if err := router.Run(":" + port); err != nil {
		logger.Error("Failed to start server:", err)
		os.Exit(1)
	}
} 