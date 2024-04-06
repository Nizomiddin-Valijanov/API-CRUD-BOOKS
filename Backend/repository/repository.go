package repository

import (
	"context"
	"fmt"
	"my-api/service"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

func GetBooksFromDb() ([]service.Book, error) {
	var books []service.Book

	cursor, err := Client.Database("Book").Collection("management").Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var book service.Book
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func AddNewBookToDB(book service.Book) (service.Book, error) {
	_, err := Client.Database("Book").Collection("management").InsertOne(context.Background(), book)
	if err != nil {
		return service.Book{}, err
	}

	return book, nil
}

func RemoveBookDB(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	res, err := Client.Database("Book").Collection("management").DeleteOne(context.Background(), filter)
	fmt.Println(res)
	if err != nil {
		return err
	}
	return nil
}

func GetOneBookFromDb(id primitive.ObjectID) (service.Book, error) {
	var book service.Book

	filter := bson.M{"_id": id}

	err := Client.Database("Book").Collection("management").FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		return service.Book{}, err
	}

	return book, nil
}

func UpdateBookDB(id string, updatedBook service.Book) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": updatedBook}

	_, err = Client.Database("Book").Collection("management").UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}
	return nil
}
