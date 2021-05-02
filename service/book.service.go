package service

import (
	"github.com/book-desk/dbq"
	"github.com/book-desk/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookService interface {
	Save(model.Book) *mongo.InsertOneResult
	Upsert(book model.Book) *mongo.UpdateResult
	Find(string) model.Book
	FindAll() []*model.Book
}

type bookService struct {
	dbBook dbq.Book
}

func New(database *mongo.Database) BookService {
	return &bookService{
		dbBook: dbq.NewBookCollection(database),
	}
}

func (bs *bookService) Save(book model.Book) *mongo.InsertOneResult {

	insertedBook, _ := bs.dbBook.Insert(book)

	return insertedBook
}

func (bs *bookService) Find(itemId string) model.Book {

	book := bs.dbBook.GetById(itemId)

	return book
}

func (bs *bookService) FindAll() []*model.Book {

	books, _ := bs.dbBook.GetByFilter(bson.D{{}})

	return books
}

func (bs *bookService) Upsert(book model.Book) *mongo.UpdateResult {

	updateResult, _ := bs.dbBook.Put(book)

	return updateResult
}
