package item_service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http/util"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
	"github.com/gorilla/mux"
)

// UpdateItem goDoc
// @Summary Update item by id
// @Description Update item by id
// @Tags item
// @Accept  json
// @Produce  json
// @Param item body domain.Item true "item"
// @Param id path int true "1"
// @Success 200 {object} domain.Item
// @Security ApiKeyAuth
// @Router /item/{id} [put]
func (service itemService) UpdateItem(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(fmt.Errorf("param route id is required and must be valid number")))
		return
	}

	itemRequest, err := domain.FromJSONItem(request.Body)

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(err))
		return
	}

	itemRequest.ID = id

	item, err := service.usecase.Update(*itemRequest)

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(err))
		return
	}

	json.NewEncoder(response).Encode(item)
}
