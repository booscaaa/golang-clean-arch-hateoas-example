package item_service

import "github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"

type itemService struct {
	usecase domain.ItemUsecase
}

func NewItemService(usecase domain.ItemUsecase) domain.ItemService {
	return &itemService{
		usecase: usecase,
	}
}
