package http

import (
	"net/http"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http/endpoint"
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

	itemUseCase := endpoint.Item{
		ItemUsecase: item.ItemUseCaseImpl(itemRepository),
	}

	r.PathPrefix("/doc").Handler(httpSwagger.WrapHandler)

	jsonApi.Handle("/item", http.HandlerFunc(itemUseCase.CreateItem)).Methods("POST", "OPTIONS").Name("/item")
	jsonApi.Handle("/item/{id}", http.HandlerFunc(itemUseCase.UpdateItem)).Methods("PUT", "OPTIONS").Name("/item")
	jsonApi.Handle("/item/{id}", http.HandlerFunc(itemUseCase.GetItemByID)).Methods("GET", "OPTIONS").Name("/item")
	jsonApi.Handle("/item/{id}", http.HandlerFunc(itemUseCase.DeleteItem)).Methods("DELETE", "OPTIONS").Name("/item")
	jsonApi.Handle("/item", http.HandlerFunc(itemUseCase.FetchItems)).Queries(
		"initials", "{initials}",
	).Methods("GET", "OPTIONS")

	return r
}
