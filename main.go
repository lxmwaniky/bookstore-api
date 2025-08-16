package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lxmwaniky/bookstore-api/database"
	"github.com/lxmwaniky/bookstore-api/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Note: .env file not found, using environment variables")
	}

	database.Connect()
	defer database.DB.Close()

	router := gin.Default()

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
