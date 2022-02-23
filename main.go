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

func main() {
	http_service.Run()
}
