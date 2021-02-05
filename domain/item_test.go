package domain_test

import (
	"golang-clean-arch-hateoas-example/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldCreateNewItem(t *testing.T) {
	id := int64(0)
	nome := "Tarefa 1"
	descricao := "Descrição da tarefa 1"
	data := "2020-02-02"
	sigla := "vin"

	item, err := domain.NewItem(id, nome, descricao, data, sigla)

	require.Nil(t, err)
	require.Empty(t, item.ID)
	require.Equal(t, item.Nome, nome)
	require.Equal(t, item.Descricao, descricao)
	require.Equal(t, item.Data, data)
	require.Equal(t, item.Sigla, sigla)
}

func TestShouldErrorCreateUserWithoutSigla(t *testing.T) {
	_, err := domain.NewItem(int64(0), "", "", "", "")
	require.NotNil(t, err)
}

func TestShouldCreateHateoasLinks(t *testing.T) {
	id := int64(0)
	nome := "Tarefa 1"
	descricao := "Descrição da tarefa 1"
	data := "2020-02-02"
	sigla := "vin"

	item, err := domain.NewItem(id, nome, descricao, data, sigla)

	require.Nil(t, err)

	itemHateoas := item.Hateoas()

	require.NotEmpty(t, itemHateoas.Links)

	for _, h := range itemHateoas.Links {
		require.NotEmpty(t, h.Href)
		require.NotEmpty(t, h.Method)
	}
}
