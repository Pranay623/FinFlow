package main

import (
	"log"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP", "service": "api-gateway"})
	})

	// Proxy for Order Service
	r.Any("/orders/*path", proxyTo("http://order-service:8081"))
	r.Any("/orders", proxyTo("http://order-service:8081"))

	log.Println("API Gateway running on port 8080")
	r.Run(":8080")
}

func proxyTo(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
