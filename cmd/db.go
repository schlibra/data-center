package cmd

import (
	"data-center/internal/repository"
	"log"

	"github.com/spf13/cobra"
)

var tableList = map[string]repository.Repository{
	"user":        new(repository.User),
	"permission":  new(repository.Permission),
	"group":       new(repository.Group),
	"frp-rule":    new(repository.FrpRule),
	"frp-token":   new(repository.FrpToken),
	"mail-domain": new(repository.MailDomain),
	"settings":    new(repository.Settings),
}

func createTable(name string) {
	repo := tableList[name]
	if repo == nil {
		log.Fatal("Repository not found for table: " + name)
	}
	if err := repo.New(); err != nil {
		log.Fatal("Failed to load database: " + err.Error())
	}
	if _, err := repo.Create(); err != nil {
		log.Fatal("Failed to create table: " + err.Error())
	}
	log.Println("Table" + name + " created successfully")
}
func dropTable(name string) {
	repo := tableList[name]
	if repo == nil {
		log.Fatal("Repository not found for table: " + name)
	}
	if err := repo.New(); err != nil {
		log.Fatal("Failed to load database: " + err.Error())
	}
	if _, err := repo.Drop(); err != nil {
		log.Fatal("Failed to drop table: " + err.Error())
	}
	log.Println("Table" + name + " dropped successfully")
}
func addDbCmd(rootCmd *cobra.Command) {
	dbCmd := &cobra.Command{
		Use:   "db",
		Short: "Database management commands",
	}
	dbCreateCmd := &cobra.Command{
		Use:   "create [name]",
		Short: "Create a new database",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			if name == "all" {
				for name := range tableList {
					createTable(name)
				}
			} else {
				createTable(name)
			}
		},
	}
	dbDropCmd := &cobra.Command{
		Use:   "drop [name]",
		Short: "Drop an existing database",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			if name == "all" {
				for name := range tableList {
					dropTable(name)
				}
			} else {
				dropTable(name)
			}
		},
	}
	dbCmd.AddCommand(dbCreateCmd)
	dbCmd.AddCommand(dbDropCmd)
	rootCmd.AddCommand(dbCmd)
}
