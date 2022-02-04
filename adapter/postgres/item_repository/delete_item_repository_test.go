package item_repository_test

import (
	"fmt"
	"testing"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/postgres/item_repository"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
	"github.com/bxcodec/faker/v3"
	"github.com/jackc/pgx/v4"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func TestDeleteItemRepository(t *testing.T) {
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

	mock.ExpectQuery("DELETE FROM item").WithArgs(
		fakeItem.ID,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeItem.ID,
		fakeItem.Name,
		fakeItem.Description,
		fakeItem.Initials,
		fakeItem.Date,
	))

	itemDeleted, err := itemRepository.Delete(
		fakeItem.ID,
	)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, itemDeleted.ID)
	require.Equal(t, itemDeleted.Name, fakeItem.Name)
	require.Equal(t, itemDeleted.Date, fakeItem.Date)
	require.Equal(t, itemDeleted.Initials, fakeItem.Initials)
}

func TestDeleteItemRepository_NoRows(t *testing.T) {
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

	mock.ExpectQuery("DELETE FROM item ").WithArgs(
		fakeItem.ID,
	).WillReturnError(pgx.ErrNoRows)

	_, err = itemRepository.Delete(
		fakeItem.ID,
	)

	if err.Error() != "Item not deleted" {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteItemRepository_AnyDBError(t *testing.T) {
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

	mock.ExpectQuery("DELETE FROM item").WithArgs(
		fakeItem.ID,
	).WillReturnError(fmt.Errorf("Any db problem"))

	_, err = itemRepository.Delete(
		fakeItem.ID,
	)

	if err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteItemRepository_AnyItemError(t *testing.T) {
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

	mock.ExpectQuery("DELETE FROM item").WithArgs(
		fakeItem.ID,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeItem.ID,
		fakeItem.Name,
		"",
		fakeItem.Date,
		fakeItem.Initials,
	))

	_, err = itemRepository.Delete(
		fakeItem.ID,
	)

	if err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
