package item

import (
	"fmt"
	"net/http"
	"strings"
)

func (item Item) GetHateoas(response http.ResponseWriter, request *http.Request) Item {
	item.Links = []Link{
		Link{
			Href:   fmt.Sprintf("%s://%s/item/%d", strings.ToLower(strings.Split(request.Proto, "/")[0]), request.Host, item.ID),
			Method: "GET",
		},
		Link{
			Href:   fmt.Sprintf("%s://%s/item/%d", strings.ToLower(strings.Split(request.Proto, "/")[0]), request.Host, item.ID),
			Method: "PUT",
		},
		Link{
			Href:   fmt.Sprintf("%s://%s/item/%d", strings.ToLower(strings.Split(request.Proto, "/")[0]), request.Host, item.ID),
			Method: "DELETE",
		},
	}

	return item
}
