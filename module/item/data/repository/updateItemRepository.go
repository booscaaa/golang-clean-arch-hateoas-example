package repository

import (
	"golang-restfull-hateoas-example/domain"
)

func (repository *itemRepository) Update(item domain.Item, id int64) (*domain.Item, error) {
	database := repository.database.Open()
	tx, err := database.Begin()
	defer database.Close()

	if err != nil {
		return nil, err
	}
	stmt, err := tx.Prepare(
		`UPDATE item SET nome = $1, descricao = $2, data =  to_timestamp($3, 'YYYY-MM-DD HH24:MI:SS') WHERE id = $4 and sigla = $5 returning *;`,
	)

	if err != nil {
		return nil, err
	}

	var idItem int64
	var nome string
	var descricao string
	var data string
	var sigla string

	err = stmt.QueryRow(item.Nome, item.Descricao, item.Data, id, item.Sigla).Scan(
		&idItem,
		&nome,
		&descricao,
		&data,
		&sigla,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	novoItem, err := domain.NewItem(idItem, nome, descricao, data, sigla)

	if err != nil {
		return nil, err
	}

	return novoItem, nil
}
