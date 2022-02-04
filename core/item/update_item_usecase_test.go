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

func TestUpdateItemUseCase(t *testing.T) {
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemRepository := mocks.NewMockItemRepository(mockCtrl)
	mockItemRepository.EXPECT().Update(
		fakeItem.ID,
		fakeItem.Date,
		fakeItem.Description,
		fakeItem.Name,
		fakeItem.Initials,
	).Return(&fakeItem, nil)

	itemUseCase := item.NewItemUseCase(mockItemRepository)

	item, err := itemUseCase.Update(fakeItem)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, item.ID)
	require.Equal(t, item.Name, fakeItem.Name)
	require.Equal(t, item.Date, fakeItem.Date)
	require.Equal(t, item.Initials, fakeItem.Initials)
}

func TestUpdateItemUseCase_Error(t *testing.T) {
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockItemRepository := mocks.NewMockItemRepository(mockCtrl)
	mockItemRepository.EXPECT().Update(
		fakeItem.ID,
		fakeItem.Date,
		fakeItem.Description,
		fakeItem.Name,
		fakeItem.Initials,
	).Return(nil, fmt.Errorf("Any Error"))

	itemUseCase := item.NewItemUseCase(mockItemRepository)

	_, err = itemUseCase.Update(fakeItem)

	if err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}
