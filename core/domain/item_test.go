package domain_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
	"github.com/bxcodec/faker/v3"

	"github.com/stretchr/testify/require"
)

func TestNewItem(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	item, err := domain.NewItem(
		fakeInsertItem.ID,
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Date,
		fakeInsertItem.Initials,
	)

	require.Nil(t, err)
	require.Equal(t, item.ID, fakeInsertItem.ID)
	require.Equal(t, item.Name, fakeInsertItem.Name)
	require.Equal(t, item.Description, fakeInsertItem.Description)
	require.Equal(t, item.Date, fakeInsertItem.Date)
	require.Equal(t, item.Initials, fakeInsertItem.Initials)
}

func TestNewItemWithoutInitials(t *testing.T) {
	_, err := domain.NewItem(1, "", "", "", "")
	require.NotNil(t, err)
}

func TestNewItemHateoasLinks(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	item, err := domain.NewItem(
		fakeInsertItem.ID,
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Date,
		fakeInsertItem.Initials,
	)

	require.Nil(t, err)

	itemHateoas, err := item.Hateoas()

	require.Nil(t, err)

	require.NotEmpty(t, itemHateoas.Links)

	for _, h := range itemHateoas.Links {
		require.NotEmpty(t, h.Href)
		require.NotEmpty(t, h.Method)
	}
}

func TestNewItemHateoasLinksEmpty(t *testing.T) {

	item := domain.Item{}
	itemHateoas, err := item.Hateoas()

	require.NotNil(t, err)

	require.Nil(t, itemHateoas)
}

func TestNewItemJsonItem(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	item, err := domain.NewItem(
		fakeInsertItem.ID,
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Date,
		fakeInsertItem.Initials,
	)

	require.Nil(t, err)

	json, err := json.Marshal(item)
	require.Nil(t, err)

	nItem, err := domain.FromJSONItem(strings.NewReader(string(json)))

	require.Nil(t, err)
	require.Equal(t, nItem.ID, fakeInsertItem.ID)
	require.Equal(t, nItem.Name, fakeInsertItem.Name)
	require.Equal(t, nItem.Description, fakeInsertItem.Description)
	require.Equal(t, nItem.Date, fakeInsertItem.Date)
	require.Equal(t, nItem.Initials, fakeInsertItem.Initials)
}

func TestNewItemHateoasError(t *testing.T) {
	fakeInsertItem := domain.Item{}

	err := faker.FakeData(&fakeInsertItem)
	if err != nil {
		fmt.Println(err)
	}

	_, err = domain.NewItem(
		0,
		fakeInsertItem.Name,
		fakeInsertItem.Description,
		fakeInsertItem.Date,
		fakeInsertItem.Initials,
	)

	if err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func TestNewItemJsonItemError(t *testing.T) {
	json, err := json.Marshal([]byte("{"))
	require.Nil(t, err)

	_, err = domain.FromJSONItem(strings.NewReader(string(json)))

	if err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

func TestNewItemJsonItemInvalid(t *testing.T) {
	json, err := json.Marshal(nil)
	require.Nil(t, err)

	_, err = domain.FromJSONItem(strings.NewReader(string(json)))

	if err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}
