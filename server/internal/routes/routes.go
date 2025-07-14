package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy"})
	})
}
