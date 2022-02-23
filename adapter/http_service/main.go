package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http_service/util"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/postgres"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/di"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.elastic.co/apm/module/apmgorilla"

	_ "github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http_service/docs"
)

func init() {
	viper.SetConfigFile(`./config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Clean architecture and Level 3 of REST
// @version 2022.2.4.0
// @description An application of studies on the implementation of clean architecture with golang with a plus of REST level 3 implementations
// @termsOfService todo-list-hateoas.herokuapp.com
// @contact.name Vinícius Boscardin
// @contact.email boscardinvinicius@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host todo-list-hateoas.herokuapp.com
// @BasePath /
func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	itemService := di.ItemHttpInjection(conn)

	router := mux.NewRouter()
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	jsonApiRouter := router.PathPrefix("/").Subrouter()
	jsonApiRouter.Use(util.Cors)
	jsonApiRouter.Use(apmgorilla.Middleware())

	apmgorilla.Instrument(jsonApiRouter)

	jsonApiRouter.Handle("/item", http.HandlerFunc(itemService.CreateItem)).Methods("POST", "OPTIONS").Name("/item")
	jsonApiRouter.Handle("/item/{id}", http.HandlerFunc(itemService.UpdateItem)).Methods("PUT", "OPTIONS").Name("/item")
	jsonApiRouter.Handle("/item/{id}", http.HandlerFunc(itemService.GetItemByID)).Methods("GET", "OPTIONS").Name("/item")
	jsonApiRouter.Handle("/item/{id}", http.HandlerFunc(itemService.DeleteItem)).Methods("DELETE", "OPTIONS").Name("/item")
	jsonApiRouter.Handle("/item", http.HandlerFunc(itemService.FetchItems)).Queries(
		"initials", "{initials}",
	).Methods("GET", "OPTIONS")

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Printf("RUNNING ON PORT " + port)
	http.ListenAndServe(":"+port, handlers.CompressHandler(router))
}
