package usecase

import (
	"encoding/json"
	"golang-restfull-hateoas-example/handler"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (usecase *itemUseCase) GetByID(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
	}

	item, err := usecase.repository.GetByID(id)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
	}

	payload, err := json.Marshal(item)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
	}

	response.WriteHeader(200)
	response.Write(payload)
}
