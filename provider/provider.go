package provider

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func Routes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	r = ItemProvider(r, db)
	r = SwaggerProvider(r)

	return r
}
