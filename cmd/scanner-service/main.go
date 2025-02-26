package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PortScanRequest struct {
	Target string `json:"target"`
	Ports  []int  `json:"ports"`
}

type PortScanResult struct {
	Port    int    `json:"port"`
	Status  string `json:"status"`
	Service string `json:"service"`
}

func scanPort(target string, port int, timeout time.Duration) PortScanResult {
	result := PortScanResult{Port: port}
	address := fmt.Sprintf("%s:%d", target, port)
	conn, err := net.DialTimeout("tcp", address, timeout)

	if err != nil {
		result.Status = "closed"
		return result
	}
	defer conn.Close()

	result.Status = "open"
	return result
}

func main() {
	r := gin.Default()

	r.POST("/scan/ports", func(c *gin.Context) {
		var req PortScanRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		results := make([]PortScanResult, 0)
		for _, port := range req.Ports {
			result := scanPort(req.Target, port, 2*time.Second)
			results = append(results, result)
		}

		c.JSON(http.StatusOK, gin.H{
			"target":  req.Target,
			"results": results,
		})
	})

	log.Fatal(r.Run(":8082"))
}
