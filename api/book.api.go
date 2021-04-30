package api

import (
	"github.com/book-desk/model"
	"github.com/book-desk/service"
	"github.com/gin-gonic/gin"
)

type BookApi interface {
	Find(string) model.Book
	FindAll() []*model.Book
	Save(ctx *gin.Context) model.Book
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

func (a *api) Save(ctx *gin.Context) model.Book {
	var book model.Book

	ctx.BindJSON(&book)
	a.service.Save(book)

	return book
}
