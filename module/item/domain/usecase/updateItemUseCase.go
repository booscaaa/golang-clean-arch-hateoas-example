package usecase

import (
	"encoding/json"
	"golang-restfull-hateoas-example/domain"
	"golang-restfull-hateoas-example/handler"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (usecase *itemUseCase) Update(response http.ResponseWriter, request *http.Request) {
	itemRequest, err := domain.NewJSONItem(request.Body)
	id, err := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)
	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
	}

	updatedItem, err := usecase.repository.Update(*itemRequest, id)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
	}

	payload, err := json.Marshal(updatedItem)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
	}

	response.WriteHeader(200)
	response.Write(payload)
}
