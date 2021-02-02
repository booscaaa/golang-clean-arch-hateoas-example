package repository

import (
	"golang-restfull-hateoas-example/domain"
)

func (repository *itemRepository) Create(item domain.Item) (*domain.Item, error) {
	database := repository.database.Open()
	tx, err := database.Begin()
	defer database.Close()

	if err != nil {
		return nil, err
	}
	stmt, err := tx.Prepare(
		`INSERT INTO item (nome, descricao, data, sigla)  VALUES ($1, $2, to_timestamp($3, 'YYYY-MM-DD HH24:MI:SS'), $4) returning *;`,
	)

	if err != nil {
		return nil, err
	}

	var id int64
	var nome string
	var descricao string
	var data string
	var sigla string

	err = stmt.QueryRow(item.Nome, item.Descricao, item.Data, item.Sigla).Scan(
		&id,
		&nome,
		&descricao,
		&data,
		&sigla,
	)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	novoItem, err := domain.NewItem(id, nome, descricao, data, sigla)

	if err != nil {
		return nil, err
	}

	return novoItem, nil
}
