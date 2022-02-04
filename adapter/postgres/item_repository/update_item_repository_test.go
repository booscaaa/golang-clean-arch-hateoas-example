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

func TestUpdateItemRepository(t *testing.T) {
	cols := []string{"id", "name", "description", "date", "initials"}
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	itemRepository := item_repository.NewItemRepository(mock)

	mock.ExpectQuery("UPDATE item SET").WithArgs(
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Date,
		fakeInsertItem.Initials,
		fakeInsertItem.ID,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeInsertItem.ID,
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Date,
		fakeInsertItem.Initials,
	))

	item, err := itemRepository.Update(
		fakeInsertItem.ID,
		fakeInsertItem.Date,
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Initials,
	)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, item.ID)
	require.Equal(t, item.Name, fakeInsertItem.Name)
	require.Equal(t, item.Date, fakeInsertItem.Date)
	require.Equal(t, item.Initials, fakeInsertItem.Initials)

}

func TestUpdateItemRepository_AnyDBError(t *testing.T) {
	cols := []string{"id", "name", "description", "date", "initials"}
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	itemRepository := item_repository.NewItemRepository(mock)

	mock.ExpectQuery("UPDATE item SET").WithArgs(
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Date,
		fakeInsertItem.Initials,
		fakeInsertItem.ID,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeInsertItem.ID,
		1,
		fakeInsertItem.Description,
		fakeInsertItem.Date,
		fakeInsertItem.Initials,
	))

	_, err = itemRepository.Update(
		fakeInsertItem.ID,
		fakeInsertItem.Date,
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Initials,
	)

	if err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateItemRepository_NoRows(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	itemRepository := item_repository.NewItemRepository(mock)

	mock.ExpectQuery("UPDATE item SET").WithArgs(
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Date,
		fakeInsertItem.Initials,
		fakeInsertItem.ID,
	).WillReturnError(pgx.ErrNoRows)

	_, err = itemRepository.Update(
		fakeInsertItem.ID,
		fakeInsertItem.Date,
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Initials,
	)

	if err.Error() != "Item not updated" {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestUpdateItemRepository_AnyItemError(t *testing.T) {
	cols := []string{"id", "name", "description", "date", "initials"}
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()

	itemRepository := item_repository.NewItemRepository(mock)

	mock.ExpectQuery("UPDATE item SET").WithArgs(
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Date,
		fakeInsertItem.Initials,
		fakeInsertItem.ID,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeInsertItem.ID,
		fakeInsertItem.Name,
		"",
		fakeInsertItem.Date,
		fakeInsertItem.Initials,
	))

	_, err = itemRepository.Update(
		fakeInsertItem.ID,
		fakeInsertItem.Date,
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Initials,
	)

	if err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
