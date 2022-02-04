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

func TestGetByIDUseCase(t *testing.T) {
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemRepository := mocks.NewMockItemRepository(mockCtrl)
	mockItemRepository.EXPECT().GetByID(1).Return(&fakeItem, nil)

	itemUseCase := item.NewItemUseCase(mockItemRepository)

	item, err := itemUseCase.GetByID(1)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, item.ID)
	require.Equal(t, item.Name, fakeItem.Name)
	require.Equal(t, item.Date, fakeItem.Date)
	require.Equal(t, item.Initials, fakeItem.Initials)
}

func TestGetByIDUseCase_Error(t *testing.T) {
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemRepository := mocks.NewMockItemRepository(mockCtrl)
	mockItemRepository.EXPECT().GetByID(1).Return(nil, fmt.Errorf("Any Error"))

	itemUseCase := item.NewItemUseCase(mockItemRepository)

	_, err = itemUseCase.GetByID(1)

	if err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}
