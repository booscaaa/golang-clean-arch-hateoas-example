package repository_test

import (
	"fmt"
	"golang-clean-arch-hateoas-example/domain"
	"golang-clean-arch-hateoas-example/module/item/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestShouldUpdateItemRepository(t *testing.T) {
	id := int64(3)
	tarefa := "Tarefa 1"
	descricao := "Descrição tarefa 1"
	data := "2021-01-01 12:00:00"
	sigla := "vinicius"
	query := "UPDATE item"
	columns := []string{"id", "nome", "descricao", "data", "sigla"}
	itemToUpdate, err := domain.NewItem(id, tarefa, descricao, data, sigla)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectPrepare(query)

	mock.ExpectQuery(query).WithArgs(
		itemToUpdate.Nome,
		itemToUpdate.Descricao,
		itemToUpdate.Data,
		itemToUpdate.ID,
		itemToUpdate.Sigla,
	).WillReturnRows(
		sqlmock.NewRows(columns).AddRow(
			3,
			tarefa,
			descricao,
			data,
			sigla,
		),
	)

	mock.ExpectCommit()

	itemRepository := repository.ItemRepositoryImpl(db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when create new item", err)
	}

	itemUpdated, err := itemRepository.Update(*itemToUpdate, itemToUpdate.ID)
	// now we execute our method
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, itemUpdated.ID)
	require.Equal(t, itemUpdated.ID, int64(3))
	require.Equal(t, itemUpdated.Nome, tarefa)
	require.Equal(t, itemUpdated.Data, data)
	require.Equal(t, itemUpdated.Sigla, sigla)
}

func TestShouldRollbackUpdateItemRepository(t *testing.T) {
	id := int64(3)
	tarefa := "Tarefa 1"
	descricao := "Descrição tarefa 1"
	data := "2021-01-01 12:00:00"
	sigla := "vinicius"
	query := "UPDATE item"
	itemToUpdate, err := domain.NewItem(id, tarefa, descricao, data, sigla)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectPrepare(query)

	mock.ExpectQuery(query).WithArgs(
		itemToUpdate.Nome,
		itemToUpdate.Descricao,
		itemToUpdate.Data,
		itemToUpdate.ID,
		itemToUpdate.Sigla,
	).WillReturnError(fmt.Errorf("Mock error"))

	mock.ExpectRollback()

	// mock.ExpectClose()

	itemRepository := repository.ItemRepositoryImpl(db)

	_, err = itemRepository.Update(*itemToUpdate, id)

	if err == nil {
		t.Error("Expected error, but got none")
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
