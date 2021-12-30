package domain

import (
	"fmt"

	"github.com/spf13/viper"
)

type Link struct {
	Href   string `json:"href" example:"http(s)://<DOMAIN_OR_IP>/item/{id}"`
	Method string `json:"method" example:"GET"`
	Doc    string `json:"doc" example:"http(s)://<DOMAIN_OR_IP>/doc/index.html"`
}

func GenerateHateoasLinks(resource string, id int) []Link {
	baseUrl := viper.GetString(`hateoas.base`)

	return []Link{
		{
			Href:   fmt.Sprintf("%s/%s/%d", baseUrl, resource, id),
			Method: "GET",
			Doc:    fmt.Sprintf("%s/doc/index.html#/%s/get_%s", baseUrl, resource, resource),
		},
		{
			Href:   fmt.Sprintf("%s/item/%d", baseUrl, id),
			Method: "PUT",
			Doc:    fmt.Sprintf("%s/doc/index.html#/%s/put_%s__id_", baseUrl, resource, resource),
		},
		{
			Href:   fmt.Sprintf("%s/%s/%d", baseUrl, resource, id),
			Method: "DELETE",
			Doc:    fmt.Sprintf("%s/doc/index.html#/%s/delete_%s__id_", baseUrl, resource, resource),
		},
		{
			Href:   fmt.Sprintf("%s/%s", baseUrl, resource),
			Method: "POST",
			Doc:    fmt.Sprintf("%s/doc/index.html#/%s/post_%s", baseUrl, resource, resource),
		},
	}
}
