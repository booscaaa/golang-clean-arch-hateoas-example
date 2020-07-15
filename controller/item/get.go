package item

import (
	"api/factory"
	"api/handler"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Get(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	db := factory.GetConnection()
	defer db.Close()
	item := Item{}
	id, _ := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)

	{
		rows, err := db.Query(
			`	SELECT id, nome, descricao, to_char(data, 'DD/MM/YYYY HH24:MI:SS') FROM item where id = $1 ORDER BY data asc;	`, id,
		)
		e, isEr := handler.CheckErr(err)

		if isEr {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(e.ReturnError())
			return
		}

		for rows.Next() {
			err = rows.Scan(
				&item.ID,
				&item.Nome,
				&item.Descricao,
				&item.Data,
			)
			e, isEr := handler.CheckErr(err)

			if isEr {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write(e.ReturnError())
				return
			}

			item = item.GetHateoas(response, request)
		}
	}

	payload, _ := json.Marshal(item)

	response.WriteHeader(http.StatusOK)
	response.Write(payload)
}
