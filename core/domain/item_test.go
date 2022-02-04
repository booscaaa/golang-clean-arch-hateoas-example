package domain_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
	"github.com/bxcodec/faker/v3"

	"github.com/stretchr/testify/require"
)

func TestNewItem(t *testing.T) {
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}
	fakeItem.Date, _ = time.Parse("2006-01-02T15:04:00Z", "2022-01-13T15:04:00Z")

	item, err := domain.NewItem(
		fakeItem.ID,
		fakeItem.Name,
		fakeItem.Description,
		fakeItem.Date,
		fakeItem.Initials,
	)

	require.Nil(t, err)
	require.Equal(t, item.ID, fakeItem.ID)
	require.Equal(t, item.Name, fakeItem.Name)
	require.Equal(t, item.Description, fakeItem.Description)
	require.Equal(t, item.Date, fakeItem.Date)
	require.Equal(t, item.Initials, fakeItem.Initials)
}

func TestNewItemWithoutInitials(t *testing.T) {
	_, err := domain.NewItem(1, "", "", time.Now(), "")
	require.NotNil(t, err)
}

func TestNewItemHateoasLinks(t *testing.T) {
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}
	fakeItem.Date, _ = time.Parse("2006-01-02T15:04:00Z", "2022-01-13T15:04:00Z")

	item, err := domain.NewItem(
		fakeItem.ID,
		fakeItem.Name,
		fakeItem.Description,
		fakeItem.Date,
		fakeItem.Initials,
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
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}
	fakeItem.Date, _ = time.Parse("2006-01-02T15:04:00Z", "2022-01-13T15:04:00Z")

	item, err := domain.NewItem(
		fakeItem.ID,
		fakeItem.Name,
		fakeItem.Description,
		fakeItem.Date,
		fakeItem.Initials,
	)

	require.Nil(t, err)

	json, err := json.Marshal(item)
	require.Nil(t, err)

	nItem, err := domain.FromJSONItem(strings.NewReader(string(json)))

	require.Nil(t, err)
	require.Equal(t, nItem.ID, fakeItem.ID)
	require.Equal(t, nItem.Name, fakeItem.Name)
	require.Equal(t, nItem.Description, fakeItem.Description)
	require.Equal(t, nItem.Date, fakeItem.Date)
	require.Equal(t, nItem.Initials, fakeItem.Initials)
}

func TestNewItemHateoasError(t *testing.T) {
	fakeItem := domain.Item{}

	err := faker.FakeData(&fakeItem)
	if err != nil {
		fmt.Println(err)
	}
	fakeItem.Date, _ = time.Parse("2006-01-02T15:04:00Z", "2022-01-13T15:04:00Z")

	_, err = domain.NewItem(
		0,
		fakeItem.Name,
		fakeItem.Description,
		fakeItem.Date,
		fakeItem.Initials,
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
