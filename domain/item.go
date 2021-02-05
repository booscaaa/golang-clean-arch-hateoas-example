package domain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/spf13/viper"
)

type Item struct {
	ID        int64  `json:"id" swaggerignore:"true"`
	Nome      string `json:"nome" example:"Tarefa 1"`
	Descricao string `json:"descricao" example:"Descrição da tarefa 1"`
	Data      string `json:"data" example:"2021-02-02"`
	Sigla     string `json:"sigla" example:"vin" maxLength:"3"`
	Links     []Link `json:"links" swaggerignore:"true"`
}

type ItemUsecase interface {
	Create(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
	GetByID(response http.ResponseWriter, request *http.Request)
}

type ItemRepository interface {
	Create(item Item) (*Item, error)
	Update(item Item, id int64) (*Item, error)
	Delete(id int64) (*Item, error)
	Fetch(sigla string) (*[]Item, error)
	GetByID(id int64) (*Item, error)
}

func (item *Item) isValid() error {
	_, err := govalidator.ValidateStruct(item)
	if err != nil {
		return err
	}

	if item.Descricao == "" {
		return fmt.Errorf("No description")
	}
	return nil
}

func NewItem(id int64, nome string, descricao string, data string, sigla string) (*Item, error) {
	item := Item{
		ID:        id,
		Nome:      nome,
		Descricao: descricao,
		Data:      data,
		Sigla:     sigla,
	}
	item = item.Hateoas()

	err := item.isValid()
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func NewJSONItem(body io.ReadCloser) (*Item, error) {
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

func (item *Item) Hateoas() Item {
	baseUrl := viper.GetString(`hateoas.base`)
	item.Links = []Link{
		{
			Href:   fmt.Sprintf("%s/item/%d", baseUrl, item.ID),
			Method: "GET",
		},
		{
			Href:   fmt.Sprintf("%s/item/%d", baseUrl, item.ID),
			Method: "PUT",
		},
		{
			Href:   fmt.Sprintf("%s/item/%d", baseUrl, item.ID),
			Method: "DELETE",
		},
	}

	return *item
}
