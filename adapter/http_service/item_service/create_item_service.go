package item_service

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http_service/util"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
)

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
func (service itemService) CreateItem(response http.ResponseWriter, request *http.Request) {
	itemRequest, err := domain.FromJSONItem(request.Body)

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(err))
		return
	}

	item, err := service.usecase.Create(*itemRequest)

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(err))
		return
	}

	json.NewEncoder(response).Encode(item)
}
