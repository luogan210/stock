package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"server/config"
	"server/router"
	"server/storage"
	"syscall"
	"time"
)

func main() {
	cfg := config.Load()
	fmt.Printf("Starting server in %s on :%s\n", cfg.Env, cfg.HTTPPort)

	// init db
	db, err := storage.OpenSQLite(cfg.SQLitePath)
	if err != nil {
		fmt.Printf("sqlite open failed: %v\n", err)
		return
	}
	if err := db.Migrate(); err != nil {
		fmt.Printf("sqlite migrate failed: %v\n", err)
		return
	}
	defer db.Close()

	r := router.SetupRouter(db)

	srv := &http.Server{
		Addr:         ":" + cfg.HTTPPort,
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// Run server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("server error: %v\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	}

	fmt.Println("Server exiting")
}
