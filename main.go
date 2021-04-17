package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Junkes887/3bases-server-a/controller"
	"github.com/Junkes887/3bases-server-a/database"
	"github.com/Junkes887/3bases-server-a/repository"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

func injects(collection *mongo.Collection, context context.Context) controller.Client {
	data := repository.Client{
		DB:  collection,
		CTX: context,
	}

	controller := controller.Client{
		DB:  collection,
		CTX: context,
		REP: data,
	}

	return controller
}

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	DATABASE := os.Getenv("DATABASE")
	COLLECTION := os.Getenv("COLLECTION")
	context := context.Background()

	database := database.Context{CTX: context}
	client := database.CreateConnection()

	collection := client.Database(DATABASE).Collection(COLLECTION)

	controller := injects(collection, context)

	router := httprouter.New()
	router.GET("/", controller.FindAll)
	router.GET("/:id", controller.Find)
	router.POST("/", controller.Save)
	router.PUT("/:id", controller.Upadate)
	router.DELETE("/:id", controller.Delete)

	c := cors.AllowAll()
	handlerCors := c.Handler(router)

	fmt.Println("Listem " + PORT + ".....")
	http.ListenAndServe(":"+PORT, handlerCors)
}
