package item

import (
	"api/factory"
	"api/handler"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Delete(response http.ResponseWriter, request *http.Request) {
	db := factory.GetConnection()
	defer db.Close()

	id, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)

	{
		tx, err := db.Begin()
		e, isEr := handler.CheckErr(err)

		if isEr {
			tx.Rollback()
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(e.ReturnError())
			return
		}
		stmt, err := tx.Prepare(
			`	DELETE FROM item where id = $1;	`)

		e, isEr = handler.CheckErr(err)

		if isEr {
			tx.Rollback()
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(e.ReturnError())
			return
		}

		_, err = stmt.Exec(
			id,
		)

		e, isEr = handler.CheckErr(err)

		if isEr {
			tx.Rollback()
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(e.ReturnError())
			return
		}

		tx.Commit()
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte(""))
}
