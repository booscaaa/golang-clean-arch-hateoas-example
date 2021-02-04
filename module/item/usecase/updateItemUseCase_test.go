package usecase_test

import (
	"encoding/json"
	"golang-clean-arch-hateoas-example/domain"
	"golang-clean-arch-hateoas-example/module/item/usecase"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang-clean-arch-hateoas-example/domain/mocks"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShould200UpdateItemUseCase(t *testing.T) {
	mockRequestItem := domain.Item{
		Nome:      "Tarefa 1",
		Descricao: "Descrição da tarefa 1",
		Data:      "2020-02-02 00:00:00",
		Sigla:     "vnb",
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

	r, _ := http.NewRequest("PUT", "/item/257", strings.NewReader(string(j)))
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "257",
	}

	r = mux.SetURLVars(r, vars)

	repository := new(mocks.ItemRepository)
	repository.On("Update", mock.Anything, mock.AnythingOfType("int64")).Return(&mockResponseItem, nil)

	usecase := usecase.ItemUseCaseImpl(repository)
	usecase.Update(w, r)

	result := w.Result()

	_, err = ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	result.Body.Close()

	if http.StatusOK != result.StatusCode {
		t.Error("expected", http.StatusOK, "got", result.StatusCode)
	}
}

func TestShould500UpdateItemUseCase_IncorrectURLParams(t *testing.T) {
	mockRequestItem := domain.Item{
		Nome:      "Tarefa 1",
		Descricao: "Descrição da tarefa 1",
		Data:      "2020-02-02 00:00:00",
		Sigla:     "vnb",
	}

	j, err := json.Marshal(mockRequestItem)
	assert.NoError(t, err)

	r, _ := http.NewRequest("PUT", "/item/", strings.NewReader(string(j)))
	w := httptest.NewRecorder()

	vars := map[string]string{}

	r = mux.SetURLVars(r, vars)

	repository := new(mocks.ItemRepository)
	repository.On("Update", mock.Anything, mock.AnythingOfType("int64")).Return(nil, nil)

	usecase := usecase.ItemUseCaseImpl(repository)
	usecase.Update(w, r)

	result := w.Result()

	_, err = ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	result.Body.Close()

	if http.StatusInternalServerError != result.StatusCode {
		t.Error("expected", http.StatusInternalServerError, "got", result.StatusCode)
	}
}

func TestShould500UpdateItemUseCase_IncorrectURLJson(t *testing.T) {
	r, _ := http.NewRequest("PUT", "/item/257", strings.NewReader(string("{\"key\":\"value\"}")))
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "257",
	}

	r = mux.SetURLVars(r, vars)

	repository := new(mocks.ItemRepository)
	repository.On("Update", mock.Anything, mock.AnythingOfType("int64")).Return(nil, nil)

	usecase := usecase.ItemUseCaseImpl(repository)
	usecase.Update(w, r)

	result := w.Result()

	_, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	result.Body.Close()

	if http.StatusInternalServerError != result.StatusCode {
		t.Error("expected", http.StatusInternalServerError, "got", result.StatusCode)
	}
}
