/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/booscaaa/golang-clean-arch-hateoas-example/adapter/postgres"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/core/domain"
	"github.com/booscaaa/golang-clean-arch-hateoas-example/di"
	"github.com/spf13/cobra"
)

// itemCmd represents the item command
var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		conn := postgres.GetConnection(ctx)
		defer conn.Close()

		itemUseCase := di.ItemInjection(conn)

		create, _ := cmd.Flags().GetBool("create")

		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		initials, _ := cmd.Flags().GetString("initials")
		date, _ := cmd.Flags().GetString("date")
		dateTime, _ := time.Parse("2006-01-02T15:04:00Z", date)

		itemRequest, err := domain.NewItem(-1, name, description, dateTime, initials)
		if err != nil {
			log.Fatal(err)
		}

		if create {
			newItem, err := itemUseCase.Create(*itemRequest)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Item created", newItem)

			return
		}
	},
}

func init() {
	rootCmd.AddCommand(itemCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// itemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	itemCmd.Flags().BoolP("create", "c", false, "Help message for toggle")
	itemCmd.Flags().BoolP("delete", "d", false, "Help message for toggle")
	itemCmd.Flags().BoolP("update", "u", false, "Help message for toggle")
	itemCmd.Flags().BoolP("find", "f", false, "Help message for toggle")

	itemCmd.Flags().String("name", "n", "Help message for toggle")
	itemCmd.Flags().String("description", "", "Help message for toggle")
	itemCmd.Flags().String("date", "d", "Help message for toggle")
	itemCmd.Flags().String("initials", "i", "Help message for toggle")
}
