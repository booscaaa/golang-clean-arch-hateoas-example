package item_repository

import (
	"context"
	"fmt"
	"golang-clean-arch-hateoas-example/core/domain"

	"github.com/jackc/pgx/v4"
)

func (db *itemRepository) Create(name, description, date, sigla string) (*domain.Item, error) {
	var idA int
	var nameA string
	var descriptionA string
	var dateA string
	var siglaA string

	ctx := context.Background()

	err := db.database.QueryRow(
		ctx,
		"INSERT INTO item (nome, descricao, data, sigla)  VALUES ($1, $2, to_timestamp($3, 'YYYY-MM-DD HH24:MI:SS'), $4) returning *;",
		name, description, date, sigla,
	).Scan(
		&idA,
		&nameA,
		&descriptionA,
		&dateA,
		&siglaA,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("Item not created")
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
