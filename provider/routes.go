package provider

import (
	"net/http"

	"golang-restfull-hateoas-example/factory"
	"golang-restfull-hateoas-example/module/item/data/repository"
	"golang-restfull-hateoas-example/module/item/domain/usecase"

	"github.com/gorilla/mux"
)

//auth is a local function to control the session in middleware
func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE, PUT")
		response.Header().Set("Content-Type", "application/json")
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if request.Method == "OPTIONS" {
			response.WriteHeader(http.StatusOK)
		} else {
			next.ServeHTTP(response, request)
		}
	})
}

func Routes() *mux.Router {
	r := mux.NewRouter()

	connection := factory.GetConnection()
	itemRepository := repository.ItemRepositoryImpl(connection)
	itemUseCase := usecase.ItemUseCaseImpl(itemRepository)

	r.Handle("/item/{id}", auth(http.HandlerFunc(itemUseCase.Update))).Methods("PUT", "OPTIONS").Name("/item")
	r.Handle("/item/{id}", auth(http.HandlerFunc(itemUseCase.GetByID))).Methods("GET", "OPTIONS").Name("/item")
	r.Handle("/item/{id}", auth(http.HandlerFunc(itemUseCase.Delete))).Methods("DELETE", "OPTIONS").Name("/item")

	r.Handle("/item", auth(http.HandlerFunc(itemUseCase.Fetch))).Queries(
		"sigla", "{sigla}",
	).Methods("GET", "OPTIONS")

	return r
}
