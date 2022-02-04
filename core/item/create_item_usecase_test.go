package item_test

import (
	"fmt"
	"testing"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain/mocks"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/item"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateItemUseCase(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemRepository := mocks.NewMockItemRepository(mockCtrl)
	mockItemRepository.EXPECT().Create(
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Date,
		fakeInsertItem.Initials,
	).Return(&fakeInsertItem, nil)

	itemUseCase := item.NewItemUseCase(mockItemRepository)

	item, err := itemUseCase.Create(fakeInsertItem)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, item.ID)
	require.Equal(t, item.Name, fakeInsertItem.Name)
	require.Equal(t, item.Date, fakeInsertItem.Date)
	require.Equal(t, item.Initials, fakeInsertItem.Initials)
}

func TestCreateItemUseCase_Error(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemRepository := mocks.NewMockItemRepository(mockCtrl)
	mockItemRepository.EXPECT().Create(
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Date,
		fakeInsertItem.Initials,
	).Return(nil, fmt.Errorf("Any Error"))

	itemUseCase := item.NewItemUseCase(mockItemRepository)

	_, err = itemUseCase.Create(fakeInsertItem)

	if err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}