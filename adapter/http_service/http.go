package http_service

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
	httpSwagger "github.com/swaggo/http-swagger"
	"go.elastic.co/apm/module/apmgorilla"

	_ "github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http_service/docs"
)

// @title Clean architecture and Level 3 of REST
// @version 2021.12.5.0
// @description An application of studies on the implementation of clean architecture with golang with a plus of REST level 3 implementations
// @termsOfService todo-list-hateoas.herokuapp.com
// @contact.name Vinícius Boscardin
// @contact.email boscardinvinicius@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host todo-list-hateoas.herokuapp.com
// @BasePath /
func Run() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	itemService := di.ItemHttpInjection(conn)

	router := mux.NewRouter()
	jsonApiRouter := router.PathPrefix("/").Subrouter()
	jsonApiRouter.Use(util.Cors)
	jsonApiRouter.Use(apmgorilla.Middleware())

	apmgorilla.Instrument(jsonApiRouter)

	router.PathPrefix("/doc").Handler(httpSwagger.WrapHandler)

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
