package main

import (
	"log"

	"github.com/book-desk/route"
	"github.com/gin-gonic/gin"
)

func main() {

	err := route.PrepareDbConnection()

	if err == nil {
		server := gin.New()

		server.Use(
			gin.Recovery(),
			// gin.Logger(), // gin Default Logger
			// middlewares.Logger(),
		)

		route.InitRoutesGroups(server)
		server.Run(":8080")
	} else {
		log.Fatal(err)
	}

}
