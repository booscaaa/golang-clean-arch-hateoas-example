package usecase_test

import (
	"golang-clean-arch-hateoas-example/domain"
	"golang-clean-arch-hateoas-example/module/item/usecase"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-clean-arch-hateoas-example/domain/mocks"

	"github.com/stretchr/testify/mock"
)

func TestShould200FetchItensUseCase(t *testing.T) {
	mockResponseItem := []domain.Item{
		{
			Nome:      "Tarefa 1",
			Descricao: "Descrição da tarefa 1",
			Data:      "2020-02-02 00:00:00",
			Sigla:     "vin",
		},
	}

	r, _ := http.NewRequest("GET", "/item?sigla=vin", nil)
	w := httptest.NewRecorder()

	repository := new(mocks.ItemRepository)
	repository.On("Fetch", mock.AnythingOfType("string")).Return(&mockResponseItem, nil)

	usecase := usecase.ItemUseCaseImpl(repository)
	usecase.Fetch(w, r)

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

func TestShould500FetchItensUseCase_IncorrectURLParams(t *testing.T) {
	r, _ := http.NewRequest("GET", "/item", nil)
	w := httptest.NewRecorder()

	repository := new(mocks.ItemRepository)
	repository.On("Fetch", mock.AnythingOfType("string")).Return(nil, nil)

	usecase := usecase.ItemUseCaseImpl(repository)
	usecase.Fetch(w, r)

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
