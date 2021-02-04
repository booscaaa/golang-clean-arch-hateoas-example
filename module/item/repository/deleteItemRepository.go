package repository

import (
	"golang-clean-arch-hateoas-example/domain"
)

func (repository *itemRepository) Delete(id int64) (*domain.Item, error) {
	tx, err := repository.database.Begin()
	// defer repository.database.Close()

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	stmt, err := tx.Prepare(`DELETE FROM item where id = $1 returning *;`)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var idItem int64
	var nome string
	var descricao string
	var data string
	var sigla string

	err = stmt.QueryRow(id).Scan(
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
