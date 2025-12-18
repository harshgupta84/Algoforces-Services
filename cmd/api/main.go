package main

import (
	"algoforces/internal/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	authHandler := handlers.NewAuthHandler(authUseCase)
	r := gin.Default()

	r.GET("/health", handlers.GetHealth)
	r.POST("/signup", authHandler.Signup)
	r.POST("/login", authHandler.Login)
	// 4. Start the Server
	fmt.Println("Starting Algoforces API on :8080...")
	r.Run(":8080")
}
