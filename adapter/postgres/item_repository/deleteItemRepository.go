package item_repository

import (
	"context"
	"fmt"
	"golang-clean-arch-hateoas-example/core/domain"

	"github.com/jackc/pgx/v4"
)

func (db *itemRepository) Delete(id int) (*domain.Item, error) {
	var idA int
	var nameA string
	var descriptionA string
	var dateA string
	var siglaA string

	ctx := context.Background()

	err := db.database.QueryRow(
		ctx,
		"DELETE FROM item WHERE id = $1 returning *;",
	).Scan(
		&idA,
		&nameA,
		&descriptionA,
		&dateA,
		&siglaA,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("Item not deleted")
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
