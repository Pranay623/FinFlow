package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("Notification Service Started...")
	log.Println("Consuming events from Kafka topics...")

	// Keep the service running
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	log.Println("Waiting for order events... (Press Ctrl+C to stop)")
	<-sigChan
	log.Println("Notification Service shutting down...")
}
