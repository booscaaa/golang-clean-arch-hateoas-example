package di

import (
	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http_service/item_service"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/postgres/item_repository"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/item"
	"github.com/jackc/pgx/v4/pgxpool"
)

func ItemHttpInjection(sql *pgxpool.Pool) domain.ItemService {
	itemRepository := item_repository.NewItemRepository(sql)
	itemUseCase := item.NewItemUseCase(itemRepository)
	itemService := item_service.NewItemService(itemUseCase)

	return itemService
}

func ItemInjection(sql *pgxpool.Pool) domain.ItemUsecase {
	itemRepository := item_repository.NewItemRepository(sql)
	itemUseCase := item.NewItemUseCase(itemRepository)

	return itemUseCase
}
