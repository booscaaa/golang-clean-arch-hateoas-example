package item

import (
	"api/factory"
	"api/handler"
	"encoding/json"
	"fmt"
	"net/http"
)

func Create(response http.ResponseWriter, request *http.Request) {
	db := factory.GetConnection()
	defer db.Close()
	item := Item{}

	var i Item
	if err := json.NewDecoder(request.Body).Decode(&i); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("500 - Something bad happened!"))
	}

	fmt.Println(i)

	{
		tx, err := db.Begin()
		e, isEr := handler.CheckErr(err)

		if isEr {
			tx.Rollback()
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(e.ReturnError())
			return
		}
		stmt, err := tx.Prepare(`INSERT INTO item (nome, descricao, data) VALUES ($1, $2, to_timestamp($3, 'YYYY-MM-DD HH24:MI:SS')) returning *;`)

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
		).Scan(
			&item.ID,
			&item.Nome,
			&item.Descricao,
			&item.Data,
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
