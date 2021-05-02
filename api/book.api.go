package api

import (
	"github.com/book-desk/model"
	"github.com/book-desk/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookApi interface {
	Find(string) model.Book
	FindAll() []*model.Book
	Save(ctx *gin.Context) *mongo.InsertOneResult
	Update(ctx *gin.Context) *mongo.UpdateResult
}

type api struct {
	service service.BookService
}

func New(service service.BookService) BookApi {
	return &api{
		service: service,
	}
}

func (a *api) Find(itemId string) model.Book {
	return a.service.Find(itemId)
}

func (a *api) FindAll() []*model.Book {
	return a.service.FindAll()
}

func (a *api) Save(ctx *gin.Context) *mongo.InsertOneResult {
	var book model.Book

	ctx.BindJSON(&book)
	insertResult := a.service.Save(book)

	return insertResult
}

func (a *api) Update(ctx *gin.Context) *mongo.UpdateResult {
	var book model.Book

	ctx.BindJSON(&book)
	updateResult := a.service.Upsert(book)

	return updateResult
}
