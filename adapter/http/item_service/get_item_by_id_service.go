package item_service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http/util"
	"github.com/gorilla/mux"
)

// GetItemByID goDoc
// @Summary Get item by id
// @Description Get item by id
// @Tags item
// @Accept  json
// @Produce  json
// @Param id path int true "1"
// @Success 200 {object} domain.Item
// @Security ApiKeyAuth
// @Router /item/{id} [get]
func (service itemService) GetItemByID(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(fmt.Errorf("param route id is required and must be valid number")))
		return
	}

	item, err := service.usecase.GetByID(id)

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(err))
		return
	}

	json.NewEncoder(response).Encode(item)
}
