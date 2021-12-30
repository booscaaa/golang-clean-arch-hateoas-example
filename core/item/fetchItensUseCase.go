package item

import (
	"golang-clean-arch-hateoas-example/core/domain"
)

func (usecase *itemUseCase) Fetch(sigla string) (*[]domain.Item, error) {
	items, err := usecase.repository.Fetch(sigla)

	if err != nil {
		return nil, err
	}

	return items, nil
}
