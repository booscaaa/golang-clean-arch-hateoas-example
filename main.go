package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"

	"golang-clean-arch-hateoas-example/core/factory"
	"golang-clean-arch-hateoas-example/provider"
)

// @title Clean archtecture and Level 3 of REST
// @version 2021.2.1.0
// @description An application of studies on the implementation of clean architecture with golang with a plus of REST level 3 implementations
// @termsOfService todo-list-hateoas.herokuapp.com
// @contact.name Vinícius Boscardin
// @contact.email boscardinvinicius@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host todo-list-hateoas.herokuapp.com
// @BasePath /
func main() {
	env := os.Getenv("GO_ENV")
	err := godotenv.Load(".env." + env)
	// err := godotenv.Load()
	if err != nil {
		log.Fatal("Configure as variáveis de ambiente no arquivo .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
		log.Printf("Porta de acesso: "+GetLocalIP()+":%s", port)
	}

	db := factory.GetConnection()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	r := provider.Routes(db)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CompressHandler(r)))
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
