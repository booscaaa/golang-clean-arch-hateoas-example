package usecase

import (
	"encoding/json"
	"golang-clean-arch-hateoas-example/core/handler"
	"golang-clean-arch-hateoas-example/domain"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Update godoc
// @Summary Change tasks into database
// @Description Change tasks into database
// @Tags item
// @Accept  json
// @Produce  json
// @Param item body domain.Item true "item"
// @Param id path integer true "1"
// @Success 200 {object} domain.Item
// @Router /item/{id} [put]
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
