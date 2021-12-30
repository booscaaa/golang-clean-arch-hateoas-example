package item_repository

import (
	"context"
	"fmt"
	"golang-clean-arch-hateoas-example/core/domain"

	"github.com/jackc/pgx/v4"
)

func (db *itemRepository) Update(id int, name, description, date, sigla string) (*domain.Item, error) {
	var idA int
	var nameA string
	var descriptionA string
	var dateA string
	var siglaA string

	ctx := context.Background()

	err := db.database.QueryRow(
		ctx,
		"UPDATE ITEM SET nome = $1, descricao = $2, data = to_timestamp($3, 'YYYY-MM-DD HH24:MI:SS'), sigla = $4 WHERE id = $5",
		name, description, date, sigla, id,
	).Scan(
		&idA,
		&nameA,
		&descriptionA,
		&dateA,
		&siglaA,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("Item not updated")
	}

	if err != nil {
		return nil, err
	}

	item, err := domain.NewItem(idA, nameA, descriptionA, dateA, siglaA)

	if err != nil {
		return nil, err
	}

	return item, nil
}
