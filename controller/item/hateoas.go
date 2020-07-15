package item

import (
	"fmt"
	"net/http"
)

func (item Item) GetHateoas(response http.ResponseWriter, request *http.Request) Item {
	item.Links = []Link{
		Link{
			Href:   fmt.Sprintf("%s://%s/item/%d", request.URL.Scheme, request.Host, item.ID),
			Method: "GET",
		},
		Link{
			Href:   fmt.Sprintf("%s://%s/item/%d", request.URL.Scheme, request.Host, item.ID),
			Method: "PUT",
		},
		Link{
			Href:   fmt.Sprintf("%s://%s/item/%d", request.URL.Scheme, request.Host, item.ID),
			Method: "DELETE",
		},
	}

	return item
}
