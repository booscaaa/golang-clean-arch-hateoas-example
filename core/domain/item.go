package domain

import (
	"encoding/json"
	"fmt"
	"io"
)

type Item struct {
	ID          int    `json:"id" swaggerignore:"true"`
	Name        string `json:"name" example:"Tarefa 1"`
	Description string `json:"description" example:"Descrição da tarefa 1"`
	Date        string `json:"date" example:"2021-02-02"`
	Sigla       string `json:"sigla" example:"vin" maxLength:"3"`
	Links       []Link `json:"_links"`
}

type ItemUsecase interface {
	Create(item Item) (*Item, error)
	Update(item Item) (*Item, error)
	Delete(id int) (*Item, error)
	Fetch(sigla string) (*[]Item, error)
	GetByID(id int) (*Item, error)
}

type ItemRepository interface {
	Create(date, description, name, sigla string) (*Item, error)
	Update(id int, date, description, name, sigla string) (*Item, error)
	Delete(id int) (*Item, error)
	Fetch(sigla string) (*[]Item, error)
	GetByID(id int) (*Item, error)
}

func (item *Item) isValid() error {
	if item.Description == "" {
		return fmt.Errorf("No description")
	}
	return nil
}

func NewItem(id int, name string, description string, date string, sigla string) (*Item, error) {
	item := Item{
		ID:          id,
		Name:        name,
		Description: description,
		Date:        date,
		Sigla:       sigla,
	}

	err := item.isValid()
	if err != nil {
		return nil, err
	}

	hItem, err := item.Hateoas()
	if err != nil {
		return nil, err
	}

	return hItem, nil
}

func FromJSONItem(body io.Reader) (*Item, error) {
	item := Item{}
	if err := json.NewDecoder(body).Decode(&item); err != nil {
		return nil, err
	}

	err := item.isValid()
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (item *Item) Hateoas() (*Item, error) {
	if item.ID == 0 {
		return nil, fmt.Errorf("No item to generate hateoas")
	}

	item.Links = GenerateHateoasLinks("item", item.ID)

	return item, nil
}
