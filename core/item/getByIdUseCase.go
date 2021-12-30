package item

import (
	"golang-clean-arch-hateoas-example/core/domain"
)

func (usecase *itemUseCase) GetByID(id int) (*domain.Item, error) {
	item, err := usecase.repository.GetByID(id)

	if err != nil {
		return nil, err
	}

	return item, nil
}
