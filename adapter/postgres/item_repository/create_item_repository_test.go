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

func TestCreateItemRepository(t *testing.T) {
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

	mock.ExpectQuery("INSERT INTO item (.+)").WithArgs(
		fakeItem.Name, fakeItem.Description, fakeItem.Date, fakeItem.Initials,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeItem.ID,
		fakeItem.Name,
		fakeItem.Description,
		fakeItem.Initials,
		fakeItem.Date,
	))

	itemCreated, err := itemRepository.Create(
		fakeItem.Date,
		fakeItem.Description,
		fakeItem.Name,
		fakeItem.Initials,
	)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, itemCreated.ID)
	require.Equal(t, itemCreated.Name, fakeItem.Name)
	require.Equal(t, itemCreated.Date, fakeItem.Date)
	require.Equal(t, itemCreated.Initials, fakeItem.Initials)
}

func TestCreateItemRepository_NoRows(t *testing.T) {
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

	mock.ExpectQuery("INSERT INTO item (.+)").WithArgs(
		fakeItem.Name, fakeItem.Description, fakeItem.Date, fakeItem.Initials,
	).WillReturnError(pgx.ErrNoRows)

	_, err = itemRepository.Create(
		fakeItem.Date,
		fakeItem.Description,
		fakeItem.Name,
		fakeItem.Initials,
	)

	if err.Error() != "Item not created" {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateItemRepository_AnyDBError(t *testing.T) {
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

	mock.ExpectQuery("INSERT INTO item (.+)").WithArgs(
		fakeItem.Name, fakeItem.Description, fakeItem.Date, fakeItem.Initials,
	).WillReturnError(fmt.Errorf("Any db problem"))

	_, err = itemRepository.Create(
		fakeItem.Date,
		fakeItem.Description,
		fakeItem.Name,
		fakeItem.Initials,
	)

	if err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateItemRepository_AnyItemError(t *testing.T) {
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

	mock.ExpectQuery("INSERT INTO item (.+)").WithArgs(
		fakeItem.Name, "", fakeItem.Date, fakeItem.Initials,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeItem.ID,
		fakeItem.Name,
		"",
		fakeItem.Initials,
		fakeItem.Date,
	))

	_, err = itemRepository.Create(
		fakeItem.Date,
		"",
		fakeItem.Name,
		fakeItem.Initials,
	)

	if err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
