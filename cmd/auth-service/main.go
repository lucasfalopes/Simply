package main

import (
	"log"
	//"net/http"
	"Simply/internal/config"

	"github.com/gin-gonic/gin"
	//"github.com/golang-jwt/jwt/v5"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	r := gin.Default()
	auth := NewAuthService(cfg.JWTSecret)

	// O resto do código permanece o mesmo
	r.POST("/auth/login", func(c *gin.Context) {
		// ... código existente ...
	})

	log.Printf("Auth service starting on port %s", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}
