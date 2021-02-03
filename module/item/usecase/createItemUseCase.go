package usecase

import (
	"encoding/json"
	"golang-clean-arch-hateoas-example/core/handler"
	"golang-clean-arch-hateoas-example/domain"
	"net/http"
)

// Create godoc
// @Summary Include tasks into database
// @Description Include tasks into database
// @Tags item
// @Accept  json
// @Produce  json
// @Param item body domain.Item true "item"
// @Success 200 {object} domain.Item
// @Router /item [post]
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
