package handler

import (
	"fmt"
	"net/http"

	"my-api/repository"
	"my-api/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBooks(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	books, err := repository.GetBooksFromDb()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get books"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var newBook *service.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	createdBook, err := repository.AddNewBookToDB(*newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, createdBook)
}

func BookById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not found"})
		return
	}
	book, err := repository.GetOneBookFromDb(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, book)
}

func CheckoutBook(ctx *gin.Context) {
	idSting := ctx.Query("id")
	id, _ := primitive.ObjectIDFromHex(idSting)
	book, err := service.CheckoutBook(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, book)
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	if id == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Kitob ID'si belirtilmagan"})
		return
	}

	if err := repository.RemoveBookDB(id); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Ma'lumotlar bazasidan kitobni o'chirishda xatolik yuz berdi"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedBook service.Book

	if err := ctx.BindJSON(&updatedBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	if err := repository.UpdateBookDB(id, updatedBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update book"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
