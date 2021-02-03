package domain

type Link struct {
	Href   string `json:"href" example:"http(s)://<DOMAIN_OR_IP>/item/{id}"`
	Method string `json:"method" example:"GET"`
}
