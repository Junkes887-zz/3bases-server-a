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
	port := os.Getenv("PORT")
	context := context.Background()

	database := database.Context{CTX: context}
	client := database.CreateConnection()
	collection := client.Database("local").Collection("usuarios")

	controller := injects(collection, context)

	router := httprouter.New()
	router.GET("/:cpf", controller.Find)
	router.POST("/", controller.Save)

	c := cors.AllowAll()
	handlerCors := c.Handler(router)

	fmt.Println("Listem " + port + ".....")
	http.ListenAndServe(":"+port, handlerCors)
}
