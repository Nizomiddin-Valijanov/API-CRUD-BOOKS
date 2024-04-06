package main

import (
	"context"
	"fmt"
	"log"
	"my-api/handler"
	"my-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	repository.Client = client
	fmt.Println("Connected to MongoDB successfully")

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	router.GET("/books/:id", handler.BookById)
	router.GET("/books", handler.GetBooks)
	router.POST("/books", handler.CreateBook)
	router.PATCH("/books/:id", handler.UpdateBook)
	router.DELETE("/books/:id", handler.DeleteBook)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
