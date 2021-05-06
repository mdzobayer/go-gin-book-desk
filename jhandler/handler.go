package jhandler

import (
	"github.com/book-desk/api"
	"github.com/book-desk/config/constants"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	Name    string
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func GetRoutesConfig(bookApi api.BookApi) []RouteConfig {
	var routeConfigs []RouteConfig

	routeConfigs = append(routeConfigs, RouteConfig{
		Name:   "Find All books",
		Method: constants.GET,
		Path:   "/books",
		Handler: func(ctx *gin.Context) {
			ctx.JSON(200, bookApi.FindAll())
		},
	})

	routeConfigs = append(routeConfigs, RouteConfig{
		Name:   "Find book by id",
		Method: constants.GET,
		Path:   "/book/:id",
		Handler: func(ctx *gin.Context) {
			id := ctx.Param("id")
			ctx.JSON(200, bookApi.Find(id))
		},
	})

	routeConfigs = append(routeConfigs, RouteConfig{
		Name:   "Insert a book",
		Method: constants.POST,
		Path:   "/book",
		Handler: func(ctx *gin.Context) {
			ctx.JSON(200, bookApi.Save(ctx))
		},
	})

	routeConfigs = append(routeConfigs, RouteConfig{
		Name:   "Update a book",
		Method: constants.POST,
		Path:   "/bookupdate",
		Handler: func(ctx *gin.Context) {
			ctx.JSON(200, bookApi.Update(ctx))
		},
	})

	routeConfigs = append(routeConfigs, RouteConfig{
		Name:   "Find books by filter",
		Method: constants.POST,
		Path:   "/booksbyfilter",
		Handler: func(ctx *gin.Context) {
			ctx.JSON(200, bookApi.FindByFilter(ctx))
		},
	})

	return routeConfigs
}
