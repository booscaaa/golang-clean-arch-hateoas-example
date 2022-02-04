package item

import (
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
)

func (usecase *itemUseCase) Update(i domain.Item) (*domain.Item, error) {
	item, err := usecase.repository.Update(i.ID, i.Date, i.Description, i.Name, i.Initials)

	if err != nil {
		return nil, err
	}

	return item, nil
}
