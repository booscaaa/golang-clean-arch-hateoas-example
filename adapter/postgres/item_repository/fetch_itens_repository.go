package item_repository

import (
	"context"
	"time"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
)

func (db *itemRepository) Fetch(initials string) (*[]domain.Item, error) {
	// defer repository.database.Close()
	items := []domain.Item{}

	ctx := context.Background()

	rows, err := db.database.Query(
		ctx,
		`SELECT id, name, description, date, initials FROM item where initials = $1 ORDER BY date asc;`,
		initials,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var idA int
		var nameA, desctiptionA, initialsA string
		var dateA time.Time

		err = rows.Scan(
			&idA,
			&nameA,
			&desctiptionA,
			&dateA,
			&initialsA,
		)

		if err != nil {
			return nil, err
		}

		item, err := domain.NewItem(idA, nameA, desctiptionA, dateA, initialsA)

		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}
