package main

import (
	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/http_service"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`./config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Clean architecture and Level 3 of REST
// @version 2022.2.4.0
// @description An application of studies on the implementation of clean architecture with golang with a plus of REST level 3 implementations
// @termsOfService todo-list-hateoas.herokuapp.com
// @contact.name Vinícius Boscardin
// @contact.email boscardinvinicius@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host todo-list-hateoas.herokuapp.com
// @BasePath /
func main() {
	http_service.Run()
}
