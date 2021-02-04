package repository_test

import (
	"fmt"
	"golang-clean-arch-hateoas-example/module/item/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestShouldFetchItensRepository(t *testing.T) {
	tarefa := "Tarefa n"
	descricao := "Descrição tarefa n"
	data := "01/01/2021"
	sigla := "vin"
	query := "SELECT (.+) FROM item"
	columns := []string{"id", "nome", "descricao", "data", "sigla"}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(query).WithArgs(sigla).WillReturnRows(
		sqlmock.NewRows(columns).AddRow(
			3,
			tarefa,
			descricao,
			data,
			sigla,
		).AddRow(
			4,
			tarefa,
			descricao,
			data,
			sigla,
		).AddRow(
			5,
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

	itens, err := itemRepository.Fetch(sigla)
	fmt.Println(len(*itens))
	// now we execute our method
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, itens)
	require.Equal(t, len(*itens), 3)
}
