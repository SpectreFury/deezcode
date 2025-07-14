package handlers

import (
	"context"
	"net/http"

	"github.com/SpectreFury/deezcode/server/internal/config"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	config *config.Config
}

func NewHealthHandler(config *config.Config) *HealthHandler {
	return &HealthHandler{
		config: config,
	}
}

// Health returns a simple health check response
func (h *HealthHandler) Health(c *gin.Context) {
	if err := h.config.Database.Conn.Ping(context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Database connection failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Database connection is healthy",
	})
}
