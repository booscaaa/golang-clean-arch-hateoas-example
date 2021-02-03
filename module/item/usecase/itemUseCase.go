package usecase

import "golang-clean-arch-hateoas-example/domain"

type itemUseCase struct {
	repository domain.ItemRepository
}

func ItemUseCaseImpl(repository domain.ItemRepository) domain.ItemUsecase {
	return &itemUseCase{
		repository: repository,
	}
}
