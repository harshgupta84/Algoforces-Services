package main

import (
	"algoforces/internal/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", handlers.GetHealth)

	// 4. Start the Server
	fmt.Println("Starting Algoforces API on :8080...")
	r.Run(":8080")
}
