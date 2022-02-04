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
)

func TestCreateItemService(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemUseCase := mocks.NewMockItemUsecase(mockCtrl)
	mockItemUseCase.EXPECT().Create(
		fakeInsertItem,
	).Return(&fakeInsertItem, nil)

	itemService := item_service.NewItemService(mockItemUseCase)

	payload, _ := json.Marshal(fakeInsertItem)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/item", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	itemService.CreateItem(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("status code is not correct")
	}
}

func TestCreateItemService_JsonErrorFormater(t *testing.T) {
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
	r := httptest.NewRequest(http.MethodPost, "/item", strings.NewReader("abc"))
	r.Header.Set("Content-Type", "application/json")
	itemService.CreateItem(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}

func TestCreateItemService_ItemError(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemUseCase := mocks.NewMockItemUsecase(mockCtrl)
	mockItemUseCase.EXPECT().Create(fakeInsertItem).Return(nil, fmt.Errorf("Any error"))

	itemService := item_service.NewItemService(mockItemUseCase)

	payload, _ := json.Marshal(fakeInsertItem)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/item", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	itemService.CreateItem(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}
