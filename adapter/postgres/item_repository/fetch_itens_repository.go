package item_repository

import (
	"context"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
)

func (db *itemRepository) Fetch(initials string) (*[]domain.Item, error) {
	// defer repository.database.Close()
	items := []domain.Item{}

	ctx := context.Background()

	rows, err := db.database.Query(
		ctx,
		`SELECT id, nome, descricao, to_char(data, 'DD/MM/YYYY'), initials FROM item where initials = $1 ORDER BY data asc;`,
		initials,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var idA int
		var nameA, desctiptionA, dateA, initialsA string

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
