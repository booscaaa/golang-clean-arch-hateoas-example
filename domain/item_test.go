package domain_test

import (
	"encoding/json"
	"golang-clean-arch-hateoas-example/domain"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldCreateNewItem(t *testing.T) {
	id := int64(1)
	nome := "Tarefa 1"
	descricao := "Descrição da tarefa 1"
	data := "2020-02-02"
	sigla := "vin"

	item, err := domain.NewItem(id, nome, descricao, data, sigla)

	require.Nil(t, err)
	require.Equal(t, item.ID, id)
	require.Equal(t, item.Nome, nome)
	require.Equal(t, item.Descricao, descricao)
	require.Equal(t, item.Data, data)
	require.Equal(t, item.Sigla, sigla)
}

func TestShouldErrorCreateUserWithoutSigla(t *testing.T) {
	_, err := domain.NewItem(int64(1), "", "", "", "")
	require.NotNil(t, err)
}

func TestShouldCreateHateoasLinks(t *testing.T) {
	id := int64(1)
	nome := "Tarefa 1"
	descricao := "Descrição da tarefa 1"
	data := "2020-02-02"
	sigla := "vin"

	item, err := domain.NewItem(id, nome, descricao, data, sigla)

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
	id := int64(1)
	nome := "Tarefa 1"
	descricao := "Descrição da tarefa 1"
	data := "2020-02-02"
	sigla := "vin"

	item, err := domain.NewItem(id, nome, descricao, data, sigla)

	require.Nil(t, err)

	json, err := json.Marshal(item)
	require.Nil(t, err)

	nItem, err := domain.NewJSONItem(strings.NewReader(string(json)))

	require.Nil(t, err)
	require.Equal(t, nItem.ID, id)
	require.Equal(t, nItem.Nome, nome)
	require.Equal(t, nItem.Descricao, descricao)
	require.Equal(t, nItem.Data, data)
	require.Equal(t, nItem.Sigla, sigla)
}
