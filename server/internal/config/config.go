package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

type ServerConfig struct {
	Port         string        `json:"port"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
}

type DatabaseConfig struct {
	Conn *pgx.Conn `json:"-"` // Don't serialize the connection
	URI  string    `json:"-"` // Don't expose credentials in JSON
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	config := &Config{
		Server: ServerConfig{
			Port:         getEnvWithDefault("PORT", "8080"),
			ReadTimeout:  getDurationEnvWithDefault("READ_TIMEOUT", 10*time.Second),
			WriteTimeout: getDurationEnvWithDefault("WRITE_TIMEOUT", 10*time.Second),
		},
		Database: DatabaseConfig{
			URI: os.Getenv("DATABASE_URI"),
		},
	}

	if config.Database.URI == "" {
		return nil, fmt.Errorf("DATABASE_URI environment variable is required")
	}

	return config, nil
}

// ConnectDB creates a connection to the database
func (c *Config) ConnectDB(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, c.Database.URI)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test the connection
	if err := conn.Ping(ctx); err != nil {
		conn.Close(ctx)
		return fmt.Errorf("failed to ping database: %w", err)
	}

	c.Database.Conn = conn
	return nil
}

// Close gracefully closes the database connection
func (c *Config) Close() {
	if c.Database.Conn != nil {
		c.Database.Conn.Close(context.Background())
	}
}

// Helper functions for environment variable parsing
func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getDurationEnvWithDefault(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if seconds, err := strconv.Atoi(value); err == nil {
			return time.Duration(seconds) * time.Second
		}
	}
	return defaultValue
}
