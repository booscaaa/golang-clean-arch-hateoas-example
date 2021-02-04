package repository

import (
	"database/sql"
	"golang-clean-arch-hateoas-example/domain"
)

type itemRepository struct {
	database *sql.DB
}

func ItemRepositoryImpl(connection *sql.DB) domain.ItemRepository {
	return &itemRepository{
		database: connection,
	}
}
