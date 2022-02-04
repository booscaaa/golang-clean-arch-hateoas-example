package http

import (
	"net/http"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http/item_service"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http/util"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/postgres/item_repository"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/item"

	_ "github.com/booscaaa/golang-clean-arch-hateoas-example/docs"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"
)

func GenerateRoutes(conn *pgxpool.Pool) *mux.Router {
	r := mux.NewRouter()
	jsonApi := r.PathPrefix("/").Subrouter()

	cors := util.Cors
	jsonApi.Use(cors)

	itemRepository := item_repository.NewItemRepository(conn)
	itemUseCase := item.NewItemUseCase(itemRepository)
	itemService := item_service.NewItemService(itemUseCase)

	r.PathPrefix("/doc").Handler(httpSwagger.WrapHandler)

	jsonApi.Handle("/item", http.HandlerFunc(itemService.CreateItem)).Methods("POST", "OPTIONS").Name("/item")
	jsonApi.Handle("/item/{id}", http.HandlerFunc(itemService.UpdateItem)).Methods("PUT", "OPTIONS").Name("/item")
	jsonApi.Handle("/item/{id}", http.HandlerFunc(itemService.GetItemByID)).Methods("GET", "OPTIONS").Name("/item")
	jsonApi.Handle("/item/{id}", http.HandlerFunc(itemService.DeleteItem)).Methods("DELETE", "OPTIONS").Name("/item")
	jsonApi.Handle("/item", http.HandlerFunc(itemService.FetchItems)).Queries(
		"initials", "{initials}",
	).Methods("GET", "OPTIONS")

	return r
}
