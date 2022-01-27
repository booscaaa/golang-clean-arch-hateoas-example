package item

import (
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
)

func (usecase *itemUseCase) Fetch(initials string) (*[]domain.Item, error) {
	items, err := usecase.repository.Fetch(initials)

	if err != nil {
		return nil, err
	}

	return items, nil
}
