package item

import "github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"

func (usecase *itemUseCase) Create(i domain.Item) (*domain.Item, error) {
	item, err := usecase.repository.Create(i.Date, i.Description, i.Name, i.Initials)

	if err != nil {
		return nil, err
	}

	return item, nil
}
