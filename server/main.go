package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"server/config"
	"server/router"
	"server/storage"
	"server/utils"
	"syscall"
	"time"
)

func main() {
	cfg := config.Load()

	// 初始化日志系统
	logDir := "logs"
	if err := utils.InitLogger(logDir); err != nil {
		utils.LogError("初始化日志系统失败: %v", err)
		os.Exit(1)
	}

	utils.LogInfo("=========================================")
	utils.LogInfo("服务器启动中...")
	utils.LogInfo("环境: %s", cfg.Env)
	utils.LogInfo("端口: %s", cfg.HTTPPort)
	utils.LogInfo("数据库路径: %s", cfg.SQLitePath)

	// 初始化数据库
	utils.LogInfo("正在初始化数据库...")
	db, err := storage.OpenSQLite(cfg.SQLitePath)
	if err != nil {
		utils.LogError("打开SQLite数据库失败: %v", err)
		os.Exit(1)
	}
	utils.LogInfo("数据库连接成功")

	// 执行数据库迁移
	utils.LogInfo("正在执行数据库迁移...")
	if err := db.Migrate(); err != nil {
		utils.LogError("数据库迁移失败: %v", err)
		os.Exit(1)
	}
	utils.LogInfo("数据库迁移完成")
	defer func() {
		if err := db.Close(); err != nil {
			utils.LogError("关闭数据库连接失败: %v", err)
		} else {
			utils.LogInfo("数据库连接已关闭")
		}
	}()

	// 设置路由
	utils.LogInfo("正在设置路由...")
	r := router.SetupRouter()
	utils.LogInfo("路由设置完成")

	// 创建HTTP服务器
	srv := &http.Server{
		Addr:         ":" + cfg.HTTPPort,
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// 启动服务器
	utils.LogInfo("服务器正在启动，监听端口: %s", cfg.HTTPPort)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.LogError("服务器启动失败: %v", err)
			os.Exit(1)
		}
	}()

	utils.LogInfo("服务器启动成功")
	utils.LogInfo("=========================================")

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	utils.LogInfo("=========================================")
	utils.LogInfo("收到关闭信号，开始优雅关闭...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		utils.LogError("服务器强制关闭: %v", err)
	} else {
		utils.LogInfo("服务器已优雅关闭")
	}

	utils.LogInfo("服务器退出")
	utils.LogInfo("=========================================")
}
