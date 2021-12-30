package item

import "golang-clean-arch-hateoas-example/core/domain"

func (usecase *itemUseCase) Delete(id int) (*domain.Item, error) {
	item, err := usecase.repository.Delete(id)

	if err != nil {
		return nil, err
	}

	return item, nil
}
