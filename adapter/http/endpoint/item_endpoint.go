package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http/util"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"

	"github.com/gorilla/mux"
)

type Item struct {
	domain.ItemUsecase
}

// CreateItem goDoc
// @Summary Create new item
// @Description Create new item
// @Tags item
// @Accept  json
// @Produce  json
// @Param company body domain.Item true "item"
// @Success 200 {object} domain.Item
// @Security ApiKeyAuth
// @Router /item [post]
func (i Item) CreateItem(response http.ResponseWriter, request *http.Request) {
	itemRequest, err := domain.FromJSONItem(request.Body)

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(err))
		return
	}

	item, err := i.Create(*itemRequest)

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(err))
		return
	}

	json.NewEncoder(response).Encode(item)
}

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
func (i Item) UpdateItem(response http.ResponseWriter, request *http.Request) {
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

	item, err := i.Update(*itemRequest)

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(err))
		return
	}

	json.NewEncoder(response).Encode(item)
}

// DeleteItem goDoc
// @Summary Delete item by id
// @Description Delete item by id
// @Tags item
// @Accept  json
// @Produce  json
// @Param id path int true "1"
// @Success 200 {object} domain.Item
// @Security ApiKeyAuth
// @Router /item/{id} [delete]
func (i Item) DeleteItem(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(fmt.Errorf("param route id is required and must be valid number")))
		return
	}

	item, err := i.Delete(id)

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(err))
		return
	}

	json.NewEncoder(response).Encode(item)
}

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
func (i Item) GetItemByID(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(fmt.Errorf("param route id is required and must be valid number")))
		return
	}

	item, err := i.GetByID(id)

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(err))
		return
	}

	json.NewEncoder(response).Encode(item)
}

// FetchItems goDoc
// @Summary Fetch items by initials
// @Description Fetch items by initials
// @Tags item
// @Accept  json
// @Produce  json
// @Param initials query string true "vnb"
// @Success 200 {object} domain.Item
// @Security ApiKeyAuth
// @Router /item [get]
func (i Item) FetchItems(response http.ResponseWriter, request *http.Request) {
	initials := request.FormValue("initials")

	items, err := i.Fetch(initials)

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(err))
		return
	}

	json.NewEncoder(response).Encode(items)
}
