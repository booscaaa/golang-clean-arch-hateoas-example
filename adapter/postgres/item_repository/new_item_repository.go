package item_repository

import (
	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/postgres"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
)

type itemRepository struct {
	database postgres.PoolInterface
}

func NewItemRepository(connection postgres.PoolInterface) domain.ItemRepository {
	return &itemRepository{
		database: connection,
	}
}
