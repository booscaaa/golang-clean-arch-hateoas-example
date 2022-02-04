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
)

func TestFetchItemService(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemUseCase := mocks.NewMockItemUsecase(mockCtrl)
	mockItemUseCase.EXPECT().Fetch(
		fakeInsertItem.Initials,
	).Return(&[]domain.Item{fakeInsertItem}, nil)

	itemService := item_service.NewItemService(mockItemUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/item", nil)
	r.Header.Set("Content-Type", "application/json")

	q := r.URL.Query()
	q.Add("initials", fakeInsertItem.Initials)
	r.URL.RawQuery = q.Encode()

	itemService.FetchItems(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("status code is not correct")
	}
}

func TestFetchItemService_ItemError(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemUseCase := mocks.NewMockItemUsecase(mockCtrl)
	mockItemUseCase.EXPECT().Fetch(
		fakeInsertItem.Initials,
	).Return(nil, fmt.Errorf("Any item error"))

	itemService := item_service.NewItemService(mockItemUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/item", nil)
	r.Header.Set("Content-Type", "application/json")

	q := r.URL.Query()
	q.Add("initials", fakeInsertItem.Initials)
	r.URL.RawQuery = q.Encode()

	itemService.FetchItems(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}
