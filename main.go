package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"

	"golang-restfull-hateoas-example/provider"
)

func main() {
	// env := os.Getenv("GO_ENV")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Configure as variáveis de ambiente no arquivo .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
		log.Printf("Porta de acesso: "+GetLocalIP()+":%s", port)
	}

	r := provider.Routes()
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
