package item

import "github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"

type itemUseCase struct {
	repository domain.ItemRepository
}

func ItemUseCaseImpl(repository domain.ItemRepository) domain.ItemUsecase {
	return &itemUseCase{
		repository: repository,
	}
}
