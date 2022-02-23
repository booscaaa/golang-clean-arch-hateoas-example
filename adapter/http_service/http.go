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

	_ "github.com/booscaaa/golang-clean-arch-hateoas-example/docs"
)

func Run() {
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
