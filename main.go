package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lxmwaniky/bookstore-api/database"
	"github.com/lxmwaniky/bookstore-api/handlers"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Note: .env file not found, using environment variables")
	}

	database.Connect()
	defer database.DB.Close()

	router := gin.Default()
	
	// Add CORS middleware
	router.Use(corsMiddleware())

	api := router.Group("/api")
	{
		api.GET("/books", handlers.GetBooks)
		api.GET("/books/:id", handlers.GetBookByID)
		api.POST("/books", handlers.CreateBook)
		api.PUT("/books/:id", handlers.UpdateBook)
		api.DELETE("/books/:id", handlers.DeleteBook)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	router.Run(":" + port)
}
