package item_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"

	"github.com/jackc/pgx/v4"
)

func (db *itemRepository) Update(id int, date time.Time, name, description, initials string) (*domain.Item, error) {
	var idA int
	var nameA string
	var descriptionA string
	var dateA time.Time
	var initialsA string

	ctx := context.Background()

	err := db.database.QueryRow(
		ctx,
		"UPDATE item SET name = $1, description = $2, date = to_timestamp($3, 'YYYY-MM-DD HH24:MI:SS'), initials = $4 WHERE id = $5",
		name, description, date, initials, id,
	).Scan(
		&idA,
		&nameA,
		&descriptionA,
		&initialsA,
		&dateA,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("Item not updated")
	}

	if err != nil {
		return nil, err
	}

	item, err := domain.NewItem(idA, nameA, descriptionA, dateA, initialsA)

	if err != nil {
		return nil, err
	}

	return item, nil
}
