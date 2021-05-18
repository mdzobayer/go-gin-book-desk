package api

import (
	"fmt"

	"github.com/book-desk/model"
	"github.com/book-desk/service"
	"github.com/gin-gonic/gin"
)

type BookApi interface {
	Find(string) model.Book
	FindAll() []*model.Book
	FindByFilter(ctx *gin.Context) []*model.Book
	Save(ctx *gin.Context) interface{}
	Update(ctx *gin.Context) interface{}
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

func (a *api) FindByFilter(ctx *gin.Context) []*model.Book {
	var filter model.Filter

	err := ctx.BindJSON(&filter)

	if err != nil {
		_ = fmt.Errorf("\nError: FindByFilter %v", err)
		return nil
	}

	return a.service.FindByFilter(filter)
}

func (a *api) Save(ctx *gin.Context) interface{} {
	var book model.Book

	err := ctx.BindJSON(&book)

	if err != nil {
		_ = fmt.Errorf("\nError: Save %v", err)
		return nil
	}

	insertResult := a.service.Save(book)

	return insertResult
}

func (a *api) Update(ctx *gin.Context) interface{} {
	var book model.Book

	err := ctx.BindJSON(&book)

	if err != nil {
		_ = fmt.Errorf("\nError: Update %v", err)
		return nil
	}

	updateResult := a.service.Upsert(book)

	return updateResult
}
