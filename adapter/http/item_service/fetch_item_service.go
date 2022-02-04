package item_service

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http/util"
)

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
func (service itemService) FetchItems(response http.ResponseWriter, request *http.Request) {
	initials := request.FormValue("initials")

	items, err := service.usecase.Fetch(initials)

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(util.ResponseMessage(err))
		return
	}

	json.NewEncoder(response).Encode(items)
}
