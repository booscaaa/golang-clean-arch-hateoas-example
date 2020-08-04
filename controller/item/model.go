package item

type Item struct {
	ID        int64  `json:"id"`
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
	Data      string `json:"data"`
	Sigla     string `json:"sigla"`
	Links     []Link `json:"links"`
}

type Link struct {
	Href   string `json:"href"`
	Method string `json:"method"`
}
