package service

import (
	"github.com/book-desk/dbq"
	"github.com/book-desk/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookService interface {
	Save(model.Book) interface{}
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

func (bs *bookService) Save(book model.Book) interface{} {

	insertedBook, _ := bs.dbBook.Insert(book)

	return insertedBook.InsertedID
}

func (bs *bookService) Find(itemId string) model.Book {

	book := bs.dbBook.GetById(itemId)

	return book
}

func (bs *bookService) FindAll() []*model.Book {

	books, _ := bs.dbBook.GetByFilter(bson.D{{}})

	return books
}
