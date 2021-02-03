package repository

import (
	"golang-clean-arch-hateoas-example/core/factory"
	"golang-clean-arch-hateoas-example/domain"
)

type itemRepository struct {
	database factory.Connection
}

func ItemRepositoryImpl(connection factory.Connection) domain.ItemRepository {
	return &itemRepository{
		database: connection,
	}
}
