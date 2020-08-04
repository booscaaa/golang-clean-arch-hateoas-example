package provider

import (
	"api/controller/item"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//auth is a local function to control the session in middleware
func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE, PUT")
		response.Header().Set("Content-Type", "application/json")
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if request.Method == "OPTIONS" {
			response.WriteHeader(http.StatusOK)
		} else {
			// bearToken := request.Header.Get("Authorization") // bear token must be 2 params -- Bearer <token>
			// if isAuth, access := jwtauth.VerifyToken(bearToken); isAuth {
			// 	fmt.Println(access.Login)
			// 	request = jwtauth.SetContextData(request, &access)
			request.URL.Scheme = "https"

			next.ServeHTTP(response, request)
			// } else {
			// 	response.WriteHeader(http.StatusUnauthorized)
			// 	response.Write(jwtauth.ReturnMessage("Acesso negado"))
			// }
		}
	})
}

func Routes() *mux.Router {
	r := mux.NewRouter()

	// r.Handle("/item", auth(http.HandlerFunc(item.Index))).Methods("GET", "OPTIONS").Name("/item")
	r.Handle("/item", auth(http.HandlerFunc(item.Create))).Methods("POST", "OPTIONS").Name("/item")
	r.Handle("/item/{id}", auth(http.HandlerFunc(item.Update))).Methods("PUT", "OPTIONS").Name("/item")
	r.Handle("/item/{id}", auth(http.HandlerFunc(item.Get))).Methods("GET", "OPTIONS").Name("/item")
	r.Handle("/item/{id}", auth(http.HandlerFunc(item.Delete))).Methods("DELETE", "OPTIONS").Name("/item")

	r.Handle("/item", auth(http.HandlerFunc(item.Index))).Queries(
		"sigla", "{sigla}",
	).Methods("GET", "OPTIONS")

	r.Handle("/", auth(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		var links []Link
		var nomes []Nome

		r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			t, err := route.GetPathTemplate()
			if err != nil {
				return err
			}

			n, e := route.GetMethods()
			if e != nil {
				return err
			}
			fmt.Println(n[0], t, route.GetName())

			link := Link{
				Href:   fmt.Sprintf("%s://%s%s", request.URL.Scheme, request.Host, t),
				Method: n[0],
				Path:   route.GetName(),
			}

			links = append(links, link)

			return nil
		})

		for _, link := range links {
			teste := false
			for i, nome := range nomes {
				fmt.Println(link.Path, nome.Path)
				if link.Path == nome.Path {
					nome.Links = append(nome.Links, link)
					teste = true
					fmt.Println("aqui")
				}
				nomes[i] = nome
			}
			if !teste {
				nomes = append(nomes, Nome{Path: link.Path, Links: []Link{link}})
			}
		}

		fmt.Println(nomes)

		payload, _ := json.Marshal(nomes)
		response.Write(payload)

	}))).Methods("GET", "OPTIONS").Name("/")

	return r
}

type Nome struct {
	Path  string `json:"path"`
	Links []Link `json:"links"`
}

type Link struct {
	Path   string `json:"path"`
	Href   string `json:"href"`
	Method string `json:"method"`
}
