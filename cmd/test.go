package cmd

import (
	"data-center/internal/repository"
	"log"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test command",
	Run: func(cmd *cobra.Command, args []string) {
		obj, err := repository.NewFrpToken()
		if err != nil {
			log.Fatal(err)
		}
		result, err := obj.Create()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Result: %v", result)
	},
}
