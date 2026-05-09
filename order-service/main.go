package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"order-service/cache"
	"order-service/database"
	"order-service/messaging"
	"order-service/models"
)

func main() {
	// Initialize Database, Redis, and Kafka
	database.InitDB()
	cache.InitRedis()
	messaging.InitKafka()

	r := gin.Default()

	// Routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP", "service": "order-service"})
	})

	r.POST("/orders", placeOrder)
	r.GET("/orders", listOrders)
	r.GET("/orders/:id", getOrder)
	r.PATCH("/orders/:id", updateOrder)
	r.DELETE("/orders/:id", deleteOrder)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Order Service starting on port %s...", port)
	r.Run(":" + port)
}

// POST /orders
func placeOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Idempotency check
	if cache.CheckIdempotency(order.IdempotencyKey) {
		c.JSON(http.StatusConflict, gin.H{"error": "Duplicate order request"})
		return
	}

	order.Status = "PENDING"
	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save order"})
		return
	}

	// Save to Redis for idempotency
	cache.SetIdempotency(order.IdempotencyKey, 24*3600)

	// Publish to Kafka
	messaging.PublishOrderEvent(order)

	c.JSON(http.StatusCreated, order)
}

// GET /orders
func listOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

// GET /orders/:id
func getOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// PATCH /orders/:id
func updateOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&order).Updates(input)
	c.JSON(http.StatusOK, order)
}

// DELETE /orders/:id
func deleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Order{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
