package dbq

import (
	"context"
	"log"

	"github.com/book-desk/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	collection *mongo.Collection
}

func NewBookCollection(database *mongo.Database) Book {

	return Book{
		collection: database.Collection(booksC),
	}
}

func (b *Book) GetCollection() *mongo.Collection {
	return b.collection
}

func (b *Book) GetById(itemId string) (book model.Book) {

	objID, _ := primitive.ObjectIDFromHex(itemId)

	filter := bson.D{bson.E{Key: "_id", Value: objID}}

	err := b.collection.FindOne(
		context.TODO(),
		filter,
	).Decode(&book)

	if err != nil {
		log.Fatal(err)
	}

	return book
}

func (b *Book) GetByFilter(filter bson.D) ([]*model.Book, error) {

	cursor, err := b.collection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	var books []*model.Book

	for cursor.Next(context.TODO()) {

		var book model.Book
		err := cursor.Decode(&book)

		if err != nil {
			log.Fatal(err)
		}

		books = append(books, &book)
	}

	return books, nil
}

func (b *Book) Insert(dt model.Book) (*mongo.InsertOneResult, error) {

	insertOneResult, err := b.collection.InsertOne(
		context.TODO(),
		dt,
	)

	if err != nil {
		return nil, err
	}

	return insertOneResult, nil
}

func (b *Book) Update(filter bson.D, updatedData model.Book) (*mongo.UpdateResult, error) {
	updateResult, err := b.collection.UpdateOne(context.TODO(), filter, updatedData)

	if err != nil {
		return nil, err
	}

	return updateResult, nil
}

func (b *Book) Delete(filter bson.D, updatedData model.Book) (*mongo.DeleteResult, error) {
	deleteResult, err := b.collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	return deleteResult, nil
}

func (b *Book) Put(dt model.Book) (*mongo.UpdateResult, error) {

	dt.PreparePut()

	filter := bson.D{bson.E{Key: "_id", Value: dt.ID}}

	update := bson.D{bson.E{Key: "$set", Value: dt}}

	opts := options.Update().SetUpsert(true)

	updateResult, err := b.collection.UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		return nil, err
	}

	return updateResult, nil
}
