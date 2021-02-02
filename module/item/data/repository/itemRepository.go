package repository

import (
	"golang-restfull-hateoas-example/domain"
	"golang-restfull-hateoas-example/factory"
)

type itemRepository struct {
	database factory.Connection
}

func ItemRepositoryImpl(connection factory.Connection) domain.ItemRepository {
	return &itemRepository{
		database: connection,
	}
}
