package usecase

import (
	"encoding/json"
	"golang-restfull-hateoas-example/handler"
	"net/http"
)

func (usecase *itemUseCase) Fetch(response http.ResponseWriter, request *http.Request) {
	sigla := request.FormValue("sigla")
	itens, err := usecase.repository.Fetch(sigla)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
	}

	payload, err := json.Marshal(itens)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
	}

	response.WriteHeader(200)
	response.Write(payload)
}
