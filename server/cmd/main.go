package main

import (
	"context"

	"github.com/SpectreFury/deezcode/server/internal/config"
	"github.com/SpectreFury/deezcode/server/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Connect to the database
	err = config.ConnectDB(context.Background())
	if err != nil {
		panic(err)
	}

	// Initialize the Gin router
	router := gin.Default()

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler(config)

	// Setup routes
	router.GET("/health", healthHandler.Health)

	// Listen and serve on the configured port
	router.Run(":" + config.Server.Port)
}
