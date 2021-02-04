package provider

import (
	"database/sql"
	"net/http"

	"golang-clean-arch-hateoas-example/middleware"
	"golang-clean-arch-hateoas-example/module/item/repository"
	"golang-clean-arch-hateoas-example/module/item/usecase"

	"github.com/gorilla/mux"
)

func ItemProvider(r *mux.Router, db *sql.DB) *mux.Router {

	itemRepository := repository.ItemRepositoryImpl(db)
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
