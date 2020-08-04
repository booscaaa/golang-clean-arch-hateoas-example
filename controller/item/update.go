package item

import (
	"api/factory"
	"api/handler"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Update(response http.ResponseWriter, request *http.Request) {
	db := factory.GetConnection()
	defer db.Close()
	item := Item{}

	var i Item
	if err := json.NewDecoder(request.Body).Decode(&i); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("500 - Something bad happened!"))
	}
	i.ID, _ = strconv.ParseInt(mux.Vars(request)["id"], 10, 64)

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
			`	UPDATE item SET 
				nome = $1, descricao = $2, 
				data =  to_timestamp($3, 'YYYY-MM-DD HH24:MI:SS')
				WHERE id = $4 and sigla = $5 returning *;	`)

		e, isEr = handler.CheckErr(err)

		if isEr {
			tx.Rollback()
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(e.ReturnError())
			return
		}

		err = stmt.QueryRow(
			i.Nome,
			i.Descricao,
			i.Data,
			i.ID,
			i.Sigla,
		).Scan(
			&item.ID,
			&item.Nome,
			&item.Descricao,
			&item.Data,
			&item.Sigla,
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

	item = item.GetHateoas(response, request)

	payload, err := json.Marshal(item)
	e, isEr := handler.CheckErr(err)

	if isEr {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(e.ReturnError())
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(payload)
}
