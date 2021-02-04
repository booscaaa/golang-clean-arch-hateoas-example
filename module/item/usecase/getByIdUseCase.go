package usecase

import (
	"encoding/json"
	"golang-clean-arch-hateoas-example/core/handler"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetByID godoc
// @Summary Search tasks by ID
// @Description Search tasks by ID
// @Tags item
// @Accept  json
// @Produce  json
// @Param id path integer true "1"
// @Success 200 {object} domain.Item{links=[]domain.Link}
// @Router /item/{id} [get]
func (usecase *itemUseCase) GetByID(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(request)["id"], 10, 64)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
		return
	}

	item, err := usecase.repository.GetByID(id)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
		return
	}

	payload, err := json.Marshal(item)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
		return
	}

	response.WriteHeader(200)
	response.Write(payload)
}
