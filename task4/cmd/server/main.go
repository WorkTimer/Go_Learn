package main

import (
	"blog-system/config"
	"blog-system/internal/routes"
	"blog-system/pkg/database"
	"blog-system/pkg/logger"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()

	cfg := config.Load()

	gin.SetMode(cfg.Server.GinMode)

	if err := database.Connect(cfg); err != nil {
		logger.Error("Failed to connect to database:", err)
		log.Fatal(err)
	}
	defer database.Close()

	if err := database.Migrate(); err != nil {
		logger.Error("Failed to migrate database:", err)
		log.Fatal(err)
	}

	r := gin.New()

	routes.SetupRoutes(r, cfg.JWT.Secret)

	port := ":" + cfg.Server.Port
	logger.Info("Starting server on port", port)

	go func() {
		if err := r.Run(port); err != nil {
			logger.Error("Failed to start server:", err)
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")
	fmt.Println("Server stopped")
}
