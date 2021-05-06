package route

import (
	"context"
	"fmt"
	"log"

	"github.com/book-desk/api"
	"github.com/book-desk/config/constants"
	"github.com/book-desk/jhandler"
	"github.com/book-desk/service"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Db connection

var (
	clientOptions *options.ClientOptions
	client        *mongo.Client
	database      *mongo.Database

	bookService service.BookService
	bookApi     api.BookApi
)

func InitRoutes(s *gin.Engine) {

	routeConfigs := jhandler.GetRoutesConfig(bookApi)

	for _, data := range routeConfigs {
		if data.Method == constants.GET {
			s.GET(data.Path, data.Handler)
		} else if data.Method == constants.POST {
			s.POST(data.Path, data.Handler)
		}
	}
}

func PrepareDbConnection() (err error) {

	clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")

	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Connected to MongoDB!")

	database = client.Database("BookDesk")

	bookService = service.New(database)

	bookApi = api.New(bookService)

	return nil
}
