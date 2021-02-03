package provider

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "golang-clean-arch-hateoas-example/docs"
)

func SwaggerProvider(r *mux.Router) *mux.Router {
	r.PathPrefix("/doc").Handler(httpSwagger.WrapHandler)

	return r
}
