/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/cmd/cmd"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`../../config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	cmd.Execute()
}
