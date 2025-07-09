package main

import (
	"context"

	"github.com/SpectreFury/deezcode/server/internal/config"
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

	r := gin.Default()

}
