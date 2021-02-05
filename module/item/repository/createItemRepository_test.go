package repository_test

import (
	"fmt"
	"golang-clean-arch-hateoas-example/domain"
	"golang-clean-arch-hateoas-example/module/item/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestShouldCreateItemRepository(t *testing.T) {
	tarefa := "Tarefa 1"
	descricao := "Descrição tarefa 1"
	data := "2021-01-01 12:00:00"
	sigla := "vin"
	query := "INSERT INTO item"
	columns := []string{"id", "nome", "descricao", "data", "sigla"}
	itemToCreate, err := domain.NewItem(1, tarefa, descricao, data, sigla)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectPrepare(query)

	mock.ExpectQuery(query).WithArgs(
		itemToCreate.Nome,
		itemToCreate.Descricao,
		itemToCreate.Data,
		itemToCreate.Sigla,
	).WillReturnRows(
		sqlmock.NewRows(columns).AddRow(
			1,
			tarefa,
			descricao,
			data,
			sigla,
		),
	)

	mock.ExpectCommit()

	// mock.ExpectClose()

	itemRepository := repository.ItemRepositoryImpl(db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when create new item", err)
	}

	itemCreated, err := itemRepository.Create(*itemToCreate)
	// now we execute our method
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, itemCreated.ID)
	require.Equal(t, itemCreated.ID, int64(1))
	require.Equal(t, itemCreated.Nome, tarefa)
	require.Equal(t, itemCreated.Data, data)
	require.Equal(t, itemCreated.Sigla, sigla)
}

func TestShouldRollbackCreateItemRepository(t *testing.T) {
	tarefa := "Tarefa 1"
	descricao := "Descrição tarefa 1"
	data := "2021-01-01 12:00:00"
	sigla := "vin"
	query := "INSERT INTO item"
	itemToCreate, err := domain.NewItem(1, tarefa, descricao, data, sigla)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectPrepare(query)

	mock.ExpectQuery(query).WithArgs(
		tarefa,
		descricao,
		data,
		sigla,
	).WillReturnError(fmt.Errorf("Mock error"))

	mock.ExpectRollback()

	// mock.ExpectClose()

	itemRepository := repository.ItemRepositoryImpl(db)

	_, err = itemRepository.Create(*itemToCreate)

	if err == nil {
		t.Error("Expected error, but got none")
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldErrorCreateItemRepositoryEmpty(t *testing.T) {

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	itemRepository := repository.ItemRepositoryImpl(db)

	_, err = itemRepository.Create(*&domain.Item{})

	if err == nil {
		t.Error("Expected error, but got none")
	}
}
