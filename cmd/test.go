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
		mailDomain, err := repository.NewMailDomain()
		if err != nil {
			log.Fatal(err)
		}
		_, _ = mailDomain.Drop()
		_, _ = mailDomain.Create()
	},
}
