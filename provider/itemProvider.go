package provider

import (
	"net/http"

	"golang-clean-arch-hateoas-example/core/factory"
	"golang-clean-arch-hateoas-example/middleware"
	"golang-clean-arch-hateoas-example/module/item/repository"
	"golang-clean-arch-hateoas-example/module/item/usecase"

	"github.com/gorilla/mux"
)

func ItemProvider(r *mux.Router) *mux.Router {
	connection := factory.GetConnection()
	itemRepository := repository.ItemRepositoryImpl(connection)
	itemUseCase := usecase.ItemUseCaseImpl(itemRepository)

	r.Handle("/item", middleware.Cors(http.HandlerFunc(itemUseCase.Create))).Methods("POST", "OPTIONS").Name("/item")
	r.Handle("/item/{id}", middleware.Cors(http.HandlerFunc(itemUseCase.Update))).Methods("PUT", "OPTIONS").Name("/item")
	r.Handle("/item/{id}", middleware.Cors(http.HandlerFunc(itemUseCase.GetByID))).Methods("GET", "OPTIONS").Name("/item")
	r.Handle("/item/{id}", middleware.Cors(http.HandlerFunc(itemUseCase.Delete))).Methods("DELETE", "OPTIONS").Name("/item")
	r.Handle("/item", middleware.Cors(http.HandlerFunc(itemUseCase.Fetch))).Queries(
		"sigla", "{sigla}",
	).Methods("GET", "OPTIONS")

	return r
}
