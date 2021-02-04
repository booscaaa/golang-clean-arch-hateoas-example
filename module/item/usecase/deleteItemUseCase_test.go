package usecase_test

import (
	"golang-clean-arch-hateoas-example/domain"
	"golang-clean-arch-hateoas-example/module/item/usecase"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-clean-arch-hateoas-example/domain/mocks"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func TestShould200DeleteItemUseCase(t *testing.T) {
	mockResponseItem := domain.Item{
		ID:        10,
		Nome:      "Tarefa 1",
		Descricao: "Descrição da tarefa 1",
		Data:      "2020-02-02 00:00:00",
		Sigla:     "vin",
	}

	r, _ := http.NewRequest("DELETE", "/item/10", nil)
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "10",
	}

	r = mux.SetURLVars(r, vars)

	repository := new(mocks.ItemRepository)
	repository.On("Delete", mock.Anything).Return(&mockResponseItem, nil)

	usecase := usecase.ItemUseCaseImpl(repository)
	usecase.Delete(w, r)

	result := w.Result()

	_, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	result.Body.Close()

	if http.StatusOK != result.StatusCode {
		t.Error("expected", http.StatusOK, "got", result.StatusCode)
	}
}

func TestShould500DeleteItemUseCase_IncorrectURLParams(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/item/", nil)
	w := httptest.NewRecorder()

	vars := map[string]string{}

	r = mux.SetURLVars(r, vars)

	repository := new(mocks.ItemRepository)
	repository.On("Delete", mock.Anything).Return(nil, nil)

	usecase := usecase.ItemUseCaseImpl(repository)
	usecase.Delete(w, r)

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
