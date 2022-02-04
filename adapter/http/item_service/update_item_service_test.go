package item_service_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http/item_service"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain/mocks"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestUpdateItemService(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemUseCase := mocks.NewMockItemUsecase(mockCtrl)
	mockItemUseCase.EXPECT().Update(
		fakeInsertItem,
	).Return(nil, nil)

	itemService := item_service.NewItemService(mockItemUseCase)

	payload, _ := json.Marshal(fakeInsertItem)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/item/1", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")

	vars := map[string]string{
		"id": fmt.Sprint(fakeInsertItem.ID),
	}

	r = mux.SetURLVars(r, vars)

	itemService.UpdateItem(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("status code is not correct")
	}
}

func TestUpdateItemService_ParamsIDError(t *testing.T) {
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
	r := httptest.NewRequest(http.MethodPut, "/item/", nil)
	r.Header.Set("Content-Type", "application/json")

	itemService.UpdateItem(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}

func TestUpdateItemService_ItemError(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemUseCase := mocks.NewMockItemUsecase(mockCtrl)
	mockItemUseCase.EXPECT().Update(
		fakeInsertItem,
	).Return(nil, fmt.Errorf("Any item error"))

	itemService := item_service.NewItemService(mockItemUseCase)

	payload, _ := json.Marshal(fakeInsertItem)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/item/1", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")

	vars := map[string]string{
		"id": fmt.Sprint(fakeInsertItem.ID),
	}

	r = mux.SetURLVars(r, vars)

	itemService.UpdateItem(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}

func TestUpdateItemService_JsonErrorFormater(t *testing.T) {
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
	r := httptest.NewRequest(http.MethodPost, "/item/1", strings.NewReader("abc"))
	r.Header.Set("Content-Type", "application/json")

	vars := map[string]string{
		"id": fmt.Sprint(fakeInsertItem.ID),
	}

	r = mux.SetURLVars(r, vars)

	itemService.UpdateItem(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}
