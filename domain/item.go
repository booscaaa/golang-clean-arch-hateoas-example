package domain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/asaskevich/govalidator"
)

type Item struct {
	ID        int64  `json:"id"`
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
	Data      string `json:"data"`
	Sigla     string `json:"sigla"`
	Links     []Link `json:"links"`
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
	baseUrl := os.Getenv("BASE_URL")
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
