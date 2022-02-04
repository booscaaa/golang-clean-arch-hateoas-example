package item_repository

import (
	"context"
	"time"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
)

func (db *itemRepository) GetByID(id int) (*domain.Item, error) {
	var idA int
	var nameA string
	var descriptionA string
	var dateA time.Time
	var initialsA string

	ctx := context.Background()

	err := db.database.QueryRow(
		ctx,
		`SELECT id, name, description, date, initials FROM item where id = $1;`,
		id,
	).Scan(
		&idA,
		&nameA,
		&descriptionA,
		&dateA,
		&initialsA,
	)

	if err != nil {
		return nil, err
	}

	item, err := domain.NewItem(idA, nameA, descriptionA, dateA, initialsA)

	if err != nil {
		return nil, err
	}

	return item, nil
}
