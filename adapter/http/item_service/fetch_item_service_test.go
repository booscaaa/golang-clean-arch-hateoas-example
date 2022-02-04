package item_service_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http/item_service"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain/mocks"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
)

func TestFetchItemService(t *testing.T) {
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}
	fakeItem.Date, _ = time.Parse("2006-01-02T15:04:00Z", "2022-01-13T15:04:00Z")

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemUseCase := mocks.NewMockItemUsecase(mockCtrl)
	mockItemUseCase.EXPECT().Fetch(
		fakeItem.Initials,
	).Return(&[]domain.Item{fakeItem}, nil)

	itemService := item_service.NewItemService(mockItemUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/item", nil)
	r.Header.Set("Content-Type", "application/json")

	q := r.URL.Query()
	q.Add("initials", fakeItem.Initials)
	r.URL.RawQuery = q.Encode()

	itemService.FetchItems(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("status code is not correct")
	}
}

func TestFetchItemService_ItemError(t *testing.T) {
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}
	fakeItem.Date, _ = time.Parse("2006-01-02T15:04:00Z", "2022-01-13T15:04:00Z")

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemUseCase := mocks.NewMockItemUsecase(mockCtrl)
	mockItemUseCase.EXPECT().Fetch(
		fakeItem.Initials,
	).Return(nil, fmt.Errorf("Any item error"))

	itemService := item_service.NewItemService(mockItemUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/item", nil)
	r.Header.Set("Content-Type", "application/json")

	q := r.URL.Query()
	q.Add("initials", fakeItem.Initials)
	r.URL.RawQuery = q.Encode()

	itemService.FetchItems(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode == 200 {
		t.Errorf("status code is not correct")
	}
}
