package main

import (
	"algoforces/internal/domain"
	"algoforces/internal/handlers"
	"algoforces/internal/middleware"
	"algoforces/internal/repository/postgres"
	"algoforces/internal/services"
	"algoforces/pkg/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize database connection
	db, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()


	// 2. Initialize dependencies
	userRepo := postgres.NewUserRepository(db.DB)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// 3. Setup router
	r := gin.Default()

	// Public routes
	r.GET("/health", handlers.GetHealth)
	r.POST("/api/auth/signup", authHandler.Signup)
	r.POST("/api/auth/login", authHandler.Login)


	// 5. Start the Server
	fmt.Println("Starting Algoforces API on :8080...")
	r.Run(":8080")
}
