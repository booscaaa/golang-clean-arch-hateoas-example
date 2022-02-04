package item_service_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http/item_service"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain/mocks"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestUpdateItemService(t *testing.T) {
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}
	fakeItem.Date, _ = time.Parse("2006-01-02T15:04:00Z", "2022-01-13T15:04:00Z")

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemUseCase := mocks.NewMockItemUsecase(mockCtrl)
	mockItemUseCase.EXPECT().Update(
		fakeItem,
	).Return(&fakeItem, nil)

	itemService := item_service.NewItemService(mockItemUseCase)

	payload, _ := json.Marshal(fakeItem)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/item/1", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")

	vars := map[string]string{
		"id": fmt.Sprint(fakeItem.ID),
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
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}
	fakeItem.Date, _ = time.Parse("2006-01-02T15:04:00Z", "2022-01-13T15:04:00Z")

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
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}
	fakeItem.Date, _ = time.Parse("2006-01-02T15:04:00Z", "2022-01-13T15:04:00Z")

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemUseCase := mocks.NewMockItemUsecase(mockCtrl)
	mockItemUseCase.EXPECT().Update(
		fakeItem,
	).Return(nil, fmt.Errorf("Any item error"))

	itemService := item_service.NewItemService(mockItemUseCase)

	payload, _ := json.Marshal(fakeItem)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/item/1", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")

	vars := map[string]string{
		"id": fmt.Sprint(fakeItem.ID),
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
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}
	fakeItem.Date, _ = time.Parse("2006-01-02T15:04:00Z", "2022-01-13T15:04:00Z")

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemUseCase := mocks.NewMockItemUsecase(mockCtrl)

	itemService := item_service.NewItemService(mockItemUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/item/1", strings.NewReader("abc"))
	r.Header.Set("Content-Type", "application/json")

	vars := map[string]string{
		"id": fmt.Sprint(fakeItem.ID),
	}

	r = mux.SetURLVars(r, vars)

	itemService.UpdateItem(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}
