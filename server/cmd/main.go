package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	router := gin.Default()

	// Initialize database connection
	conn, err := pgx.Connect(context.Background(), "")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	router.GET("/health", func(c *gin.Context) {
		var name int

		row := conn.QueryRow(context.Background(), "SELECT id FROM health_check").Scan(&name)

		if row != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data": gin.H{
				"id": name,
			},
		})
	})

	fmt.Println("Starting server on :8080")
	router.Run()
}
