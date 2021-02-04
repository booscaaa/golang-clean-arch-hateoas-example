package repository_test

import (
	"golang-clean-arch-hateoas-example/module/item/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestShouldGetByItensRepository(t *testing.T) {
	id := int64(3)
	tarefa := "Tarefa n"
	descricao := "Descrição tarefa n"
	data := "01/01/2021"
	sigla := "vinicius"
	query := "SELECT (.+) FROM item"
	columns := []string{"id", "nome", "descricao", "data", "sigla"}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(query).WithArgs(id).WillReturnRows(
		sqlmock.NewRows(columns).AddRow(
			3,
			tarefa,
			descricao,
			data,
			sigla,
		),
	)

	itemRepository := repository.ItemRepositoryImpl(db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when create new item", err)
	}

	item, err := itemRepository.GetByID(id)
	// now we execute our method
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, item.ID)
	require.Equal(t, item.ID, int64(3))
	require.Equal(t, item.Nome, tarefa)
	require.Equal(t, item.Data, data)
	require.Equal(t, item.Sigla, sigla)
}
