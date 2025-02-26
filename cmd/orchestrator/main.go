// cmd/orchestrator/main.go
package main

import (
	"Simply/internal/config"
	"github.com/gin-gonic/gin"
	"log"

	"net/http"
)

type ServiceConfig struct {
	ScannerServiceURL string
	AuthServiceURL    string
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Configure endpoints for local development
	services := ServiceConfig{
		ScannerServiceURL: "http://localhost:8082",
		AuthServiceURL:    "http://localhost:8080",
	}

	r := gin.Default()
	orchestrator := NewOrchestratorService()

	// ... resto do código ...

	log.Printf("Orchestrator service starting on port 8081")
	log.Fatal(r.Run(":8081"))
}
