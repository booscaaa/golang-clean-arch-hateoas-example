package repository

import (
	"golang-clean-arch-hateoas-example/domain"
)

func (repository *itemRepository) Fetch(sigla string) (*[]domain.Item, error) {
	// defer repository.database.Close()
	var itens []domain.Item

	rows, err := repository.database.Query(
		`SELECT id, nome, descricao, to_char(data, 'DD/MM/YYYY'), sigla FROM item where sigla = $1 ORDER BY data asc;`,
		sigla,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int64
		var nome, descricao, data, sigla string

		err = rows.Scan(
			&id,
			&nome,
			&descricao,
			&data,
			&sigla,
		)

		if err != nil {
			return nil, err
		}

		novoItem, err := domain.NewItem(id, nome, descricao, data, sigla)

		if err != nil {
			return nil, err
		}

		itens = append(itens, *novoItem)
	}

	return &itens, nil
}
