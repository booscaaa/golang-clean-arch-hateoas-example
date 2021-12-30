package item_repository

import (
	"golang-clean-arch-hateoas-example/core/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

type itemRepository struct {
	database *pgxpool.Pool
}

func ItemRepositoryImpl(connection *pgxpool.Pool) domain.ItemRepository {
	return &itemRepository{
		database: connection,
	}
}
