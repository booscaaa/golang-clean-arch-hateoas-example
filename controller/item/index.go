package item

import (
	"api/factory"
	"api/handler"
	"encoding/json"
	"net/http"
)

func Index(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	db := factory.GetConnection()
	defer db.Close()
	item := Item{}

	item.Sigla = request.FormValue("sigla")

	var itens []Item

	{
		rows, err := db.Query(
			`	SELECT id, nome, descricao, to_char(data, 'DD/MM/YYYY HH24:MI:SS'), sigla FROM item where sigla = $1 ORDER BY data asc;	`,
			item.Sigla,
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
				&item.Sigla,
			)
			e, isEr := handler.CheckErr(err)

			if isEr {
				response.WriteHeader(http.StatusInternalServerError)
				response.Write(e.ReturnError())
				return
			}

			item = item.GetHateoas(response, request)

			itens = append(itens, item)
		}
	}

	payload, _ := json.Marshal(itens)

	response.WriteHeader(http.StatusOK)
	response.Write(payload)
}
