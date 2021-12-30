package item

import "golang-clean-arch-hateoas-example/core/domain"

func (usecase *itemUseCase) Create(i domain.Item) (*domain.Item, error) {
	item, err := usecase.repository.Create(i.Name, i.Description, i.Date, i.Sigla)

	if err != nil {
		return nil, err
	}

	return item, nil
}
