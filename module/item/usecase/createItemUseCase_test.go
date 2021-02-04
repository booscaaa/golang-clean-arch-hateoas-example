package usecase_test

import (
	"encoding/json"
	"fmt"
	"golang-clean-arch-hateoas-example/domain"
	"golang-clean-arch-hateoas-example/module/item/usecase"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang-clean-arch-hateoas-example/domain/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShould200CreateItemUseCase(t *testing.T) {
	mockRequestItem := domain.Item{
		Nome:      "Tarefa 1",
		Descricao: "Descrição da tarefa 1",
		Data:      "2020-02-02",
		Sigla:     "vin",
	}

	mockResponseItem := domain.Item{
		ID:        1,
		Nome:      "Tarefa 1",
		Descricao: "Descrição da tarefa 1",
		Data:      "2020-02-02 00:00:00",
		Sigla:     "vin",
	}

	j, err := json.Marshal(mockRequestItem)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/item", strings.NewReader(string(j)))
	res := httptest.NewRecorder()

	repository := new(mocks.ItemRepository)

	repository.On("Create", mock.Anything).Return(&mockResponseItem, nil)

	usecase := usecase.ItemUseCaseImpl(repository)

	usecase.Create(res, req)

	result := res.Result()

	_, err = ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	result.Body.Close()

	if http.StatusOK != result.StatusCode {
		t.Error("expected", http.StatusOK, "got", result.StatusCode)
	}
}

func TestShould500CreateItemUseCase_IncorrectJSON(t *testing.T) {
	mockRequestItem := domain.Item{}

	j, err := json.Marshal(mockRequestItem)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/item", strings.NewReader(string(j)))
	res := httptest.NewRecorder()

	repository := new(mocks.ItemRepository)

	repository.On("Create", mock.Anything).Return(nil, fmt.Errorf("Mock error"))

	usecase := usecase.ItemUseCaseImpl(repository)

	usecase.Create(res, req)

	result := res.Result()

	_, err = ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	result.Body.Close()

	if http.StatusInternalServerError != result.StatusCode {
		t.Error("expected", http.StatusInternalServerError, "got", result.StatusCode)
	}
}
