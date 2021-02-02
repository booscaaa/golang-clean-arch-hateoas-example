package usecase

import (
	"encoding/json"
	"golang-restfull-hateoas-example/domain"
	"golang-restfull-hateoas-example/handler"
	"net/http"
)

func (usecase *itemUseCase) Create(response http.ResponseWriter, request *http.Request) {
	itemRequest, err := domain.NewJSONItem(request.Body)
	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
	}

	insertedItem, err := usecase.repository.Create(*itemRequest)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
	}

	payload, err := json.Marshal(insertedItem)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
	}

	response.WriteHeader(200)
	response.Write(payload)
}
