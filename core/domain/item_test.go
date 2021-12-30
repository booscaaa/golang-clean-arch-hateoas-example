package domain_test

import (
	"encoding/json"
	"golang-clean-arch-hateoas-example/core/domain"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldCreateNewItem(t *testing.T) {
	id := 1
	name := "Tarefa 1"
	description := "Descrição da tarefa 1"
	date := "2020-02-02"
	sigla := "vin"

	item, err := domain.NewItem(id, name, description, date, sigla)

	require.Nil(t, err)
	require.Equal(t, item.ID, id)
	require.Equal(t, item.Name, name)
	require.Equal(t, item.Description, description)
	require.Equal(t, item.Date, date)
	require.Equal(t, item.Sigla, sigla)
}

func TestShouldErrorCreateUserWithoutSigla(t *testing.T) {
	_, err := domain.NewItem(1, "", "", "", "")
	require.NotNil(t, err)
}

func TestShouldCreateHateoasLinks(t *testing.T) {
	id := 1
	name := "Tarefa 1"
	description := "Descrição da tarefa 1"
	date := "2020-02-02"
	sigla := "vin"

	item, err := domain.NewItem(id, name, description, date, sigla)

	require.Nil(t, err)

	itemHateoas, err := item.Hateoas()

	require.Nil(t, err)

	require.NotEmpty(t, itemHateoas.Links)

	for _, h := range itemHateoas.Links {
		require.NotEmpty(t, h.Href)
		require.NotEmpty(t, h.Method)
	}
}

func TestShouldNotCreateHateoasLinksEmpty(t *testing.T) {

	item := domain.Item{}
	itemHateoas, err := item.Hateoas()

	require.NotNil(t, err)

	require.Nil(t, itemHateoas)
}

func TestShowCreateJsonItem(t *testing.T) {
	id := 1
	name := "Tarefa 1"
	description := "Descrição da tarefa 1"
	date := "2020-02-02"
	sigla := "vin"

	item, err := domain.NewItem(id, name, description, date, sigla)

	require.Nil(t, err)

	json, err := json.Marshal(item)
	require.Nil(t, err)

	nItem, err := domain.FromJSONItem(strings.NewReader(string(json)))

	require.Nil(t, err)
	require.Equal(t, nItem.ID, id)
	require.Equal(t, nItem.Name, name)
	require.Equal(t, nItem.Description, description)
	require.Equal(t, nItem.Date, date)
	require.Equal(t, nItem.Sigla, sigla)
}
