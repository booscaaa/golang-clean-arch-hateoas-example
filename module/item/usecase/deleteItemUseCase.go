package usecase

import (
	"encoding/json"
	"golang-clean-arch-hateoas-example/core/handler"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Delete godoc
// @Summary Delete tasks
// @Description Delete tasks
// @Tags item
// @Accept  json
// @Produce  json
// @Param id path integer true "1"
// @Success 200 {object} domain.Item
// @Router /item/{id} [delete]
func (usecase *itemUseCase) Delete(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
	}

	item, err := usecase.repository.Delete(id)

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
