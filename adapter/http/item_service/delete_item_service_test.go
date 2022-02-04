package item_service_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http/item_service"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain/mocks"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestDeleteItemService(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemUseCase := mocks.NewMockItemUsecase(mockCtrl)
	mockItemUseCase.EXPECT().Delete(
		1,
	).Return(&fakeInsertItem, nil)

	itemService := item_service.NewItemService(mockItemUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/item/1", nil)
	r.Header.Set("Content-Type", "application/json")

	vars := map[string]string{
		"id": "1",
	}

	r = mux.SetURLVars(r, vars)

	itemService.DeleteItem(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("status code is not correct")
	}
}

func TestDeleteItemService_ParamsIDError(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemUseCase := mocks.NewMockItemUsecase(mockCtrl)

	itemService := item_service.NewItemService(mockItemUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/item/", nil)
	r.Header.Set("Content-Type", "application/json")

	itemService.DeleteItem(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}

func TestDeleteItemService_ItemError(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemUseCase := mocks.NewMockItemUsecase(mockCtrl)
	mockItemUseCase.EXPECT().Delete(
		1,
	).Return(nil, fmt.Errorf("Any item error"))

	itemService := item_service.NewItemService(mockItemUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/item/", nil)
	r.Header.Set("Content-Type", "application/json")

	vars := map[string]string{
		"id": "1",
	}

	r = mux.SetURLVars(r, vars)

	itemService.DeleteItem(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}
