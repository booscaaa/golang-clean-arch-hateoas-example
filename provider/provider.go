package provider

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	r = ItemProvider(r)
	r = SwaggerProvider(r)

	return r
}
