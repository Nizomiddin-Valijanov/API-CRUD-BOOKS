package service

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id       primitive.ObjectID `bson:"_id"`
	Title    string             `json:"title"`
	Author   string             `json:"author"`
	Quantity int                `json:"quantity"`
}

var Books = []Book{}

func GetBookById(id primitive.ObjectID) (*Book, error) {
	for i, b := range Books {
		if b.Id == id {
			return &Books[i], nil
		}
	}
	return nil, errors.New("Book not found")
}

func CheckoutBook(id primitive.ObjectID) (*Book, error) {
	book, err := GetBookById(id)
	if err != nil {
		return nil, err
	}

	if book.Quantity <= 0 {
		return nil, errors.New("Book not found")
	}

	book.Quantity--
	return book, nil
}

func DeleteBook(id primitive.ObjectID) error {
	index := -1
	for i, b := range Books {
		if b.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("Book not found")
	}

	Books = append(Books[:index], Books[index+1:]...)
	return nil
}
