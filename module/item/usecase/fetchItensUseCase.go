package usecase

import (
	"encoding/json"
	"golang-clean-arch-hateoas-example/core/handler"
	"net/http"
)

// Fetch godoc
// @Summary Search tasks by acronym
// @Description Search tasks by acronym
// @Tags item
// @Accept  json
// @Produce  json
// @Param sigla query string true "vinicius"
// @Success 200 {array} domain.Item{links=[]domain.Link}
// @Router /item [get]
func (usecase *itemUseCase) Fetch(response http.ResponseWriter, request *http.Request) {
	sigla := request.FormValue("sigla")
	itens, err := usecase.repository.Fetch(sigla)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
	}

	payload, err := json.Marshal(itens)

	if err, isErr := handler.CheckErr(err); isErr {
		response.WriteHeader(500)
		response.Write(err.ReturnError())
	}

	response.WriteHeader(200)
	response.Write(payload)
}
