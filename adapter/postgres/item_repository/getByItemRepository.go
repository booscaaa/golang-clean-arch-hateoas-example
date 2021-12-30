package item_repository

import (
	"context"
	"golang-clean-arch-hateoas-example/core/domain"
)

func (db *itemRepository) GetByID(id int) (*domain.Item, error) {
	var idA int
	var nameA string
	var descriptionA string
	var dateA string
	var siglaA string

	ctx := context.Background()

	err := db.database.QueryRow(
		ctx,
		`SELECT id, nome, descricao, to_char(data, 'DD/MM/YYYY HH24:MI:SS'), sigla FROM item where id = $1 ORDER BY data asc;`,
		id,
	).Scan(
		&idA,
		&nameA,
		&descriptionA,
		&dateA,
		&siglaA,
	)

	if err != nil {
		return nil, err
	}

	item, err := domain.NewItem(idA, nameA, descriptionA, dateA, siglaA)

	if err != nil {
		return nil, err
	}

	return item, nil
}
