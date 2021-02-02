package repository

import (
	"golang-restfull-hateoas-example/domain"
)

func (repository *itemRepository) GetByID(id int64) (*domain.Item, error) {
	database := repository.database.Open()
	defer database.Close()

	var idItem int64
	var nome string
	var descricao string
	var data string
	var sigla string

	err := database.QueryRow(
		`SELECT id, nome, descricao, to_char(data, 'DD/MM/YYYY HH24:MI:SS'), sigla FROM item where id = $1 ORDER BY data asc;`,
		id,
	).Scan(
		&idItem,
		&nome,
		&descricao,
		&data,
		&sigla,
	)

	if err != nil {
		return nil, err
	}

	novoItem, err := domain.NewItem(idItem, nome, descricao, data, sigla)

	if err != nil {
		return nil, err
	}

	return novoItem, nil
}
