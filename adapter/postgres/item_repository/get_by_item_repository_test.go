package item_repository_test

import (
	"fmt"
	"testing"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/postgres/item_repository"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
	"github.com/bxcodec/faker/v3"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func TestGetItemByIDRepository(t *testing.T) {
	cols := []string{"id", "name", "description", "date", "initials"}
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}

	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	itemRepository := item_repository.NewItemRepository(mock)

	mock.ExpectQuery("SELECT (.+) FROM item where").WithArgs(
		1,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeItem.ID,
		fakeItem.Name,
		fakeItem.Description,
		fakeItem.Date,
		fakeItem.Initials,
	))

	item, err := itemRepository.GetByID(1)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, item.ID)
	require.Equal(t, item.Name, fakeItem.Name)
	require.Equal(t, item.Date, fakeItem.Date)
	require.Equal(t, item.Initials, fakeItem.Initials)

}

func TestGetItemByIDRepository_AnyDBError(t *testing.T) {
	cols := []string{"id", "name", "description", "date", "initials"}
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}

	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	itemRepository := item_repository.NewItemRepository(mock)

	mock.ExpectQuery("SELECT (.+) FROM item").WithArgs(
		1,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeItem.ID,
		1,
		fakeItem.Description,
		fakeItem.Date,
		fakeItem.Initials,
	))

	_, err = itemRepository.GetByID(1)

	if err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// func TestGetItemByIDRepository_ScanErro(t *testing.T) {
// 	fakeItem := domain.Item{}

// 	err := faker.FakeData(&fakeItem)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	mock, err := pgxmock.NewPool()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer mock.Close()

// 	itemRepository := item_repository.NewItemRepository(mock)

// 	mock.ExpectQuery("SELECT (.+) FROM item").WithArgs(
// 		1,
// 	).WillReturnError(fmt.Errorf("Any db problem"))

// 	_, err = itemRepository.GetByID(1)

// 	if err == nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

func TestGetItemByIDRepository_AnyItemError(t *testing.T) {
	cols := []string{"id", "name", "description", "date", "initials"}
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}

	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	itemRepository := item_repository.NewItemRepository(mock)

	mock.ExpectQuery("SELECT (.+) FROM item where").WithArgs(
		1,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeItem.ID,
		fakeItem.Name,
		"",
		fakeItem.Date,
		fakeItem.Initials,
	))

	_, err = itemRepository.GetByID(1)

	if err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
