package item_repository

import (
	"context"
	"golang-clean-arch-hateoas-example/core/domain"
)

func (db *itemRepository) Fetch(sigla string) (*[]domain.Item, error) {
	// defer repository.database.Close()
	items := []domain.Item{}

	ctx := context.Background()

	rows, err := db.database.Query(
		ctx,
		`SELECT id, nome, descricao, to_char(data, 'DD/MM/YYYY'), sigla FROM item where sigla = $1 ORDER BY data asc;`,
		sigla,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var idA int
		var nameA, desctiptionA, dateA, siglaA string

		err = rows.Scan(
			&idA,
			&nameA,
			&desctiptionA,
			&dateA,
			&siglaA,
		)

		if err != nil {
			return nil, err
		}

		item, err := domain.NewItem(idA, nameA, desctiptionA, dateA, siglaA)

		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}
